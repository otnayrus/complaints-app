package repository

import (
	"context"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) GetCategories(ctx context.Context) ([]model.Category, error) {
	rows, err := r.Db.QueryContext(ctx, getAllCategoriesQuery)
	if err != nil {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var (
			category          model.Category
			extraFieldsSchema string
		)
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&extraFieldsSchema,
		)
		if err != nil {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
		}

		err = json.Unmarshal([]byte(extraFieldsSchema), &category.ExtraFieldsSchema)
		if err != nil {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
		}

		categories = append(categories, category)
	}

	return categories, nil
}
