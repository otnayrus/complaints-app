package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/model"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) CreateUser(ctx context.Context, input *model.User) (int, error) {
	existingUser, err := r.GetUserByEmail(ctx, input.Email)
	if err != nil && !errorwrapper.IsErrorContainingCode(err, errorwrapper.ErrResourceNotFound) {
		return 0, err
	}
	if existingUser != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrResourceAlreadyExists, "user with this email already exists")
	}

	var id int
	err = r.Db.QueryRowContext(ctx, createUserQuery).Scan(&id)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	return id, nil
}
