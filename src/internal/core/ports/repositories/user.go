package repositories

import (
	"context"

	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/models"
)

type (
	FindUserParams struct {
		ID    *uint64
		Name  *string
		Email *string
	}

	CreateUserParams struct {
		Name         string
		PasswordHash string
		Email        string
	}

	UpdateUserParams struct {
		CreateUserParams

		ID uint64
	}

	FindUsersParams struct {
		Limit  *uint64
		Offset *uint64
	}

	UserRepository interface {
		GetUser(ctx context.Context, params *FindUserParams) (*models.User, error)
		GetUsers(ctx context.Context, params *FindUsersParams) ([]models.User, error)
		CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
		UpdateUser(ctx context.Context, params *UpdateUserParams) (*models.User, error)
		DeleteUser(ctx context.Context, params *FindUserParams) error
	}
)
