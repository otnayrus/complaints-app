package repository

import (
	"context"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) CreateComplaint(ctx context.Context, input *model.Complaint) (int, error) {
	extraFieldsJson, err := json.Marshal(input.ExtraFields)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	var id int
	err = r.Db.QueryRowContext(ctx, createComplaintQuery,
		input.CategoryID,
		input.Description,
		input.Status,
		input.Remarks,
		extraFieldsJson,
		input.CreatedBy,
	).Scan(&id)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return id, nil
}
