package repository

import (
	"context"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) UpdateComplaint(ctx context.Context, input *model.Complaint) error {
	extraFieldsJson, err := json.Marshal(input.ExtraFields)
	if err != nil {
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	_, err = r.Db.ExecContext(
		ctx,
		updateComplaintQuery,
		input.CreatedBy,
		input.CategoryID,
		input.Description,
		input.Status,
		input.Remarks,
		extraFieldsJson,
		input.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
