package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) GetCategories(ctx context.Context) ([]model.Category, error) {
	rows, err := r.Db.QueryContext(ctx, getAllCategoriesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.ExtraFields,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
