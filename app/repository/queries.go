package repository

const (
	createUserQuery = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	updateUserQuery = `UPDATE users SET name = $1, email = $2, password = $3, updated_at = NOW() WHERE id = $4`
	deleteUserQuery = `DELETE FROM users WHERE id = $1`

	getUserByEmailQuery = `SELECT id, name, email, password FROM users WHERE email = $1`
	getUserByIDQuery    = `SELECT id, name, email, password FROM users WHERE id = $1`
)

const (
	createCategoryQuery = `INSERT INTO categories (name, extra_fields_schema) VALUES (LOWER($1), $2) RETURNING id`
	updateCategoryQuery = `UPDATE categories SET name = LOWER($1), extra_fields_schema = $2, updated_at = NOW() WHERE id = $3`
	deleteCategoryQuery = `DELETE FROM categories WHERE id = $1`

	getAllCategoriesQuery  = `SELECT id, name, extra_fields_schema FROM categories ORDER BY name`
	getCategoryByIDQuery   = `SELECT id, name, extra_fields_schema FROM categories WHERE id = $1`
	getCategoryByNameQuery = `SELECT id, name, extra_fields_schema FROM categories WHERE name = LOWER($1)`
)

const (
	createComplaintQuery = `INSERT INTO complaints (user_id, category_id, description, status, remarks, extra_fields) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	updateComplaintQuery = `UPDATE complaints SET user_id = $1, category_id = $2, description = $3, status = $4, remarks = $5, extra_fields = $6, updated_at = NOW() WHERE id = $7`
	deleteComplaintQuery = `DELETE FROM complaints WHERE id = $1`

	getAllComplaintsQuery = `SELECT id, user_id, category_id, description, status, remarks, extra_fields FROM complaints ORDER BY id`
	getComplaintByIDQuery = `SELECT id, user_id, category_id, description, status, remarks, extra_fields FROM complaints WHERE id = $1`
)
