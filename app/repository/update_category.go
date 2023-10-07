package repository

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) UpdateCategory(ctx context.Context, input *model.Category) error {
	extraFieldsSchema, err := json.Marshal(input.ExtraFieldsSchema)
	if err != nil {
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	_, err = r.Db.ExecContext(ctx, updateCategoryQuery, input.Name, extraFieldsSchema, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errorwrapper.WrapErr(errorwrapper.ErrResourceAlreadyExists, "category with this name already exists")
		}
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return nil
}
