package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) GetComplaintByID(ctx context.Context, id int) (*model.Complaint, error) {
	var (
		complaint model.Complaint
		err       error
	)

	var extraFields string
	err = r.Db.QueryRowContext(ctx, getComplaintByIDQuery, id).Scan(
		&complaint.ID,
		&complaint.CategoryID,
		&complaint.Description,
		&complaint.Status,
		&complaint.Remarks,
		&extraFields,
		&complaint.CreatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrResourceNotFound, "complaint not found")
		}

		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	err = json.Unmarshal([]byte(extraFields), &complaint.ExtraFields)
	if err != nil {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return &complaint, nil
}
