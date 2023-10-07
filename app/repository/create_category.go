package repository

import (
	"context"
	"encoding/json"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) CreateCategory(ctx context.Context, input *model.Category) (int, error) {
	var (
		id  int
		err error
	)

	existing, err := r.GetCategoryByName(ctx, input.Name)
	if err != nil && !errorwrapper.IsErrorContainingCode(err, errorwrapper.ErrResourceNotFound) {
		return 0, err
	}
	if existing != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrResourceAlreadyExists, "category already exists")
	}

	extraFieldsSchema, err := json.Marshal(input.ExtraFieldsSchema)
	if err != nil {
		return 0, err
	}

	err = r.Db.QueryRowContext(ctx, createCategoryQuery, input.Name, extraFieldsSchema).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
