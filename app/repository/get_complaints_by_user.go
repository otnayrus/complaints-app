package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) GetComplaintsByUser(ctx context.Context, userID int) ([]model.Complaint, error) {
	rows, err := r.Db.QueryContext(ctx, getComplaintsByUserIDQuery, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrResourceNotFound, err.Error())
		}
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	defer rows.Close()

	var complaints []model.Complaint
	for rows.Next() {
		var (
			complaint   model.Complaint
			extraFields string
		)
		err := rows.Scan(
			&complaint.ID,
			&complaint.CategoryID,
			&complaint.Description,
			&complaint.Status,
			&complaint.Remarks,
			&extraFields,
			&complaint.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(extraFields), &complaint.ExtraFields)
		if err != nil {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
		}
		complaints = append(complaints, complaint)
	}

	return complaints, nil
}
