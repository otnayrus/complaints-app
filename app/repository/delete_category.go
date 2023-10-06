package repository

import "context"

func (r *Repository) DeleteCategory(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, deleteCategoryQuery, id)
	if err != nil {
		return err
	}
	return nil
}
