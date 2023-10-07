package model

import (
	"github.com/go-playground/validator/v10"
)

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
	ID int `json:"id" validate:"required"`
}

// Category ==========================================================================================================

type CreateCategoryRequest struct {
	Name              string  `json:"name" validate:"required"`
	ExtraFieldsSchema []Field `json:"extra_fields_schema" validate:"required"`
}

func (r *CreateCategoryRequest) MakeModel() *Category {
	return &Category{
		Name:              r.Name,
		ExtraFieldsSchema: r.ExtraFieldsSchema,
	}
}

func (r *CreateCategoryRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type UpdateCategoryRequest struct {
	ID                int     `json:"id" validate:"required"`
	Name              *string `json:"name" validate:"required"`
	ExtraFieldsSchema []Field `json:"extra_fields_schema" validate:"required"`
}

func (r *UpdateCategoryRequest) MakeModel(existing Category) *Category {
	if r.Name != nil {
		existing.Name = *r.Name
	}
	if r.ExtraFieldsSchema != nil {
		existing.ExtraFieldsSchema = r.ExtraFieldsSchema
	}
	return &existing
}

func (r *UpdateCategoryRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type DeleteCategoryRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetCategoryByIDRequest struct {
	ID int `uri:"id"`
}

// Complaint ==============================================

type CreateComplaintRequest struct {
	CategoryID  int     `json:"category_id" validate:"required"`
	Description string  `json:"description" validate:"required"`
	ExtraFields []Field `json:"extra_fields" validate:"required"`
}

func (r *CreateComplaintRequest) MakeModel(userID int) *Complaint {
	return &Complaint{
		CategoryID:  r.CategoryID,
		Description: r.Description,
		ExtraFields: r.ExtraFields,
		Status:      ComplaintStatusPending,
		CreatedBy:   userID,
	}
}

func (r *CreateComplaintRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type UpdateComplaintRequest struct {
	ID          int     `json:"id" validate:"required"`
	CategoryID  *int    `json:"category_id"`
	Description *string `json:"description"`
	ExtraFields []Field `json:"extra_fields"`
	Status      *int    `json:"status"`
	Remarks     *string `json:"remarks"`
}

func (r *UpdateComplaintRequest) MakeModel(existing Complaint, userID int) *Complaint {
	if r.CategoryID != nil {
		existing.CategoryID = *r.CategoryID
	}
	if r.Description != nil {
		existing.Description = *r.Description
	}
	if r.ExtraFields != nil {
		existing.ExtraFields = r.ExtraFields
	}
	if r.Status != nil {
		existing.Status = *r.Status
	}
	if r.Remarks != nil {
		existing.Remarks = *r.Remarks
	}
	return &existing
}

func (r *UpdateComplaintRequest) Validate(v *validator.Validate) error {
	return v.Struct(r)
}

type GetComplaintByIDRequest struct {
	ID int `uri:"id"`
}
