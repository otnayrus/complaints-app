package delivery

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/otnayrus/sb-rest/app/model"
)

type handler struct {
	validator *validator.Validate
	repo      repo
}

func New(repo repo) *handler {
	return &handler{
		validator: validator.New(),
		repo:      repo,
	}
}

type repo interface {
	userRepository
	categoryRepository
	complaintRepository
}

type userRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
	IsUserHaveRole(ctx context.Context, userID int, role string) (bool, error)
	GetUserRoles(ctx context.Context, userID int) (map[string]bool, error)
}

type categoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) (int, error)
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, id int) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type complaintRepository interface {
	CreateComplaint(ctx context.Context, complaint *model.Complaint) (int, error)
	GetComplaints(ctx context.Context) ([]model.Complaint, error)
	GetComplaintByID(ctx context.Context, id int) (*model.Complaint, error)
	UpdateComplaint(ctx context.Context, complaint *model.Complaint) error
	GetComplaintsByUser(ctx context.Context, userID int) ([]model.Complaint, error)
}
