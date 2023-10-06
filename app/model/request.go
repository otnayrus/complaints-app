package model

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *CreateUserRequest) MakeModel() *User {
	return &User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *CreateUserRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (r *UpdateUserRequest) MakeModel(existing User) *User {
	if r.Name != nil {
		existing.Name = *r.Name
	}
	if r.Email != nil {
		existing.Email = *r.Email
	}
	if r.Password != nil {
		existing.Password = *r.Password
	}
	return &existing
}

func (r *UpdateUserRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type DeleteUserRequest struct {
	ID int64 `json:"id" validate:"required"`
}
