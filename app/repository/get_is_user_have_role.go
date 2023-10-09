package repository

import (
	"context"
	"database/sql"

	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) IsUserHaveRole(ctx context.Context, userID int, role string) (bool, error) {
	var (
		err    error
		roleID int
	)

	// get role id
	err = r.Db.QueryRowContext(
		ctx,
		getRoleIdByNameQuery,
		role,
	).Scan(&roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errorwrapper.WrapErr(errorwrapper.ErrResourceNotFound, err.Error())
		}
		return false, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	// check if user have role
	var exists bool
	err = r.Db.QueryRowContext(
		ctx,
		getUserRoleQuery,
		userID,
		roleID,
	).Scan(&exists)
	if err != nil {
		return false, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return exists, nil
}
