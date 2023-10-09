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

	// begin transaction
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	defer tx.Rollback()

	// insert to users
	var id int
	err = r.Db.QueryRowContext(ctx, createUserQuery, input.Name, input.Email, input.Password).Scan(&id)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	// get roles where name is user (default)
	var roleID int
	err = r.Db.QueryRowContext(ctx, getRoleIdByNameQuery, model.RoleUser).Scan(&roleID)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	// insert to users_roles
	_, err = r.Db.ExecContext(ctx, createUserRoleQuery, id, roleID)
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return 0, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return id, nil
}
