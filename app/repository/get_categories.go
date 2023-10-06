package repository

import "github.com/otnayrus/sb-rest/app/model"

func (r *Repository) GetCategories() ([]model.Category, error) {
	rows, err := r.Db.Query(getAllCategoriesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.ExtraFields,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
