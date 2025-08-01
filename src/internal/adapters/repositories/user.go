package repositories

import (
	"context"

	gen "github.com/RenZorRUS/todo-backend/src/internal/adapters/databases/postgres/generated"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/mappers"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/models"
	repos "github.com/RenZorRUS/todo-backend/src/internal/core/ports/repositories"
)

type UserRepository struct {
	q *gen.Queries
}

func NewUserRepository(queries *gen.Queries) *UserRepository {
	return &UserRepository{q: queries}
}

func (r *UserRepository) GetUser(
	ctx context.Context,
	params *repos.FindUserParams,
) (*models.User, error) {
	getUserParams, err := mappers.ToGetUserParams(params)
	if err != nil {
		return nil, err
	}

	entity, err := r.q.GetUser(ctx, getUserParams)
	if err != nil {
		return nil, err
	}

	return mappers.ToUser(entity), nil
}

func (r *UserRepository) GetUsers(
	ctx context.Context,
	params *repos.FindUsersParams,
) ([]models.User, error) {
	args, err := mappers.ToGetUsersParams(params)
	if err != nil {
		return nil, err
	}

	entities, err := r.q.GetUsers(ctx, args)
	if err != nil {
		return nil, err
	}

	return mappers.ToUsers(entities), nil
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	params *repos.CreateUserParams,
) (*models.User, error) {
	args := mappers.ToCreateUserParams(params)

	entity, err := r.q.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}

	return mappers.ToUser(entity), nil
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	params *repos.UpdateUserParams,
) (*models.User, error) {
	args, err := mappers.ToUpdateUserParams(params)
	if err != nil {
		return nil, err
	}

	entity, err := r.q.UpdateUser(ctx, args)
	if err != nil {
		return nil, err
	}

	return mappers.ToUser(entity), nil
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	params *repos.FindUserParams,
) error {
	args, err := mappers.ToDeleteUserParams(params)
	if err != nil {
		return err
	}

	return r.q.DeleteUser(ctx, args)
}
