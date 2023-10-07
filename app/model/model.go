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
	ID                int       `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	ExtraFieldsSchema []Field   `json:"extra_fields_schema" db:"extra_fields_schema"`
	CreatedAt         time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Complaint struct {
	ID          int       `json:"id" db:"id"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Description string    `json:"description" db:"description"`
	Status      int       `json:"status" db:"status"`
	Remarks     string    `json:"remarks" db:"remarks"`
	ExtraFields []Field   `json:"extra_fields" db:"extra_fields"`
	CreatedBy   int       `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Field struct {
	FieldType FieldType   `json:"field_type"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value,omitempty"`
}

type ElementKey string
