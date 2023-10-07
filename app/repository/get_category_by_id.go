package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	var (
		category model.Category
		err      error
	)

	err = r.Db.QueryRowContext(ctx, getCategoryByIDQuery, id).Scan(&category.ID, &category.Name, &category.ExtraFields)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
