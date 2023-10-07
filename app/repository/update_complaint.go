package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) UpdateComplaint(ctx context.Context, input *model.Complaint) error {
	_, err := r.Db.ExecContext(
		ctx,
		updateComplaintQuery,
		input.CreatedBy,
		input.CategoryID,
		input.Description,
		input.Status,
		input.Remarks,
		input.ExtraFields,
		input.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
