package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Category struct {
	ID          int              `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	ExtraFields map[string]Field `json:"extra_fields" db:"extra_fields"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

type Complaint struct {
	ID          int              `json:"id" db:"id"`
	UserID      int              `json:"user_id" db:"user_id"`
	CategoryID  int              `json:"category_id" db:"category_id"`
	Description string           `json:"description" db:"description"`
	Status      int              `json:"status" db:"status"`
	Remarks     string           `json:"remarks" db:"remarks"`
	ExtraFields map[string]Field `json:"extra_fields" db:"extra_fields"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}
