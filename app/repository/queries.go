package repository

const (
	createUserQuery = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	updateUserQuery = `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	deleteUserQuery = `DELETE FROM users WHERE id = $1`

	getUserByEmailQuery = `SELECT id, name, email, password FROM users WHERE email = $1`
	getUserByIDQuery    = `SELECT id, name, email, password FROM users WHERE id = $1`
)
