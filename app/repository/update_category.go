package repository

import (
	"context"
	"strings"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) UpdateCategory(ctx context.Context, input *model.Category) error {
	_, err := r.Db.ExecContext(ctx, updateCategoryQuery, input.Name, input.ExtraFieldsSchema, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errorwrapper.WrapErr(errorwrapper.ErrResourceAlreadyExists, "category with this name already exists")
		}
		return err
	}

	return nil
}
