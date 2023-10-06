package repository

import (
	"context"
	"strings"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) UpdateUser(ctx context.Context, input *model.User) error {
	_, err := r.Db.ExecContext(ctx, updateUserQuery, input.Name, input.Email, input.Password, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errorwrapper.WrapErr(errorwrapper.ErrResourceAlreadyExists, "user with this email already exists")
		}
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	return nil
}
