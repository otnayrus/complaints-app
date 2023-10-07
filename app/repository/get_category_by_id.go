package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	var (
		category model.Category
		err      error
	)

	var extraFieldsSchema string
	err = r.Db.QueryRowContext(ctx, getCategoryByIDQuery, id).Scan(&category.ID, &category.Name, &extraFieldsSchema)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrResourceNotFound, "category not found")
		}
		return nil, err
	}

	err = json.Unmarshal([]byte(extraFieldsSchema), &category.ExtraFieldsSchema)
	if err != nil {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return &category, nil
}
