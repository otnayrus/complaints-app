package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) GetComplaints(ctx context.Context) ([]model.Complaint, error) {
	rows, err := r.Db.QueryContext(ctx, getAllComplaintsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var complaints []model.Complaint
	for rows.Next() {
		var complaint model.Complaint
		err := rows.Scan(
			&complaint.ID,
			&complaint.UserID,
			&complaint.CategoryID,
			&complaint.Description,
			&complaint.Status,
			&complaint.Remarks,
			&complaint.ExtraFields,
		)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}

	return complaints, nil
}
