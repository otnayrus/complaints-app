package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) CreateComplaint(ctx context.Context, input *model.Complaint) error {
	_, err := r.Db.ExecContext(
		ctx,
		createComplaintQuery,
		input.UserID,
		input.CategoryID,
		input.Description,
		input.Status,
		input.Remarks,
		input.ExtraFields,
	)
	if err != nil {
		return err
	}

	return nil
}
