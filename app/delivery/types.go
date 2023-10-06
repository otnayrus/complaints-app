package delivery

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/otnayrus/sb-rest/app/model"
)

type handler struct {
	validator *validator.Validate
	userRepo  userRepository
}

func New(userRepo userRepository) *handler {
	return &handler{
		validator: validator.New(),
		userRepo:  userRepo,
	}
}

type userRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
}
