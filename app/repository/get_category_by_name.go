package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	var (
		category model.Category
		err      error
	)

	var extraFieldsSchema string
	err = r.Db.QueryRowContext(ctx, getCategoryByNameQuery, name).Scan(&category.ID, &category.Name, &extraFieldsSchema)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorwrapper.WrapErr(errorwrapper.ErrResourceNotFound, "category not found")
		}
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	err = json.Unmarshal([]byte(extraFieldsSchema), &category.ExtraFieldsSchema)
	if err != nil {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return &category, nil
}
