package repository

import (
	"context"

	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, deleteUserQuery, id)
	if err != nil {
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	return nil
}
