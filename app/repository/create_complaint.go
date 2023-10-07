package repository

import (
	"context"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) CreateComplaint(ctx context.Context, input *model.Complaint) error {
	extraFieldsJson, err := json.Marshal(input.ExtraFields)
	if err != nil {
		return err
	}

	_, err = r.Db.ExecContext(
		ctx,
		createComplaintQuery,
		input.UserID,
		input.CategoryID,
		input.Description,
		input.Status,
		input.Remarks,
		extraFieldsJson,
	)
	if err != nil {
		return err
	}

	return nil
}
