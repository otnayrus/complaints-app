package repository

import "context"

func (r *Repository) GetUserRoles(ctx context.Context, userID int) (map[string]bool, error) {
	var roles = make(map[string]bool)
	rows, err := r.Db.QueryContext(ctx, getUserRolesQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		err = rows.Scan(&role)
		if err != nil {
			return nil, err
		}

		roles[role] = true
	}

	return roles, nil
}
