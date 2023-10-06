package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
)

func (r *Repository) GetCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	var (
		id  int
		err error
	)

	err = r.Db.QueryRowContext(ctx, getCategoryByNameQuery, name).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID: id,
	}, nil
}
