package mappers

import (
	"errors"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	gen "github.com/RenZorRUS/todo-backend/src/internal/adapters/databases/postgres/generated"
	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/models"
	repos "github.com/RenZorRUS/todo-backend/src/internal/core/ports/repositories"
	"github.com/RenZorRUS/todo-backend/src/internal/utils"
	"github.com/samber/lo"
)

func ToGetUserParams(params *repos.FindUserParams) (*gen.GetUserParams, error) {
	userID, err := utils.NilableUintToInt[uint64, int64](params.ID)
	if err != nil {
		return nil, err
	}

	return &gen.GetUserParams{
		ID:    userID,
		Name:  params.Name,
		Email: params.Email,
	}, nil
}

func ToDeleteUserParams(params *repos.FindUserParams) (*gen.DeleteUserParams, error) {
	userID, err := utils.NilableUintToInt[uint64, int64](params.ID)
	if err != nil {
		return nil, err
	}

	return &gen.DeleteUserParams{
		ID:    userID,
		Name:  params.Name,
		Email: params.Email,
	}, nil
}

func ToGetUsersParams(params *repos.FindUsersParams) (*gen.GetUsersParams, error) {
	var errs error

	limit, err := GetOrDefaultRowsConstraint(params.Limit, consts.LimitDefault)
	errs = errors.Join(errs, err)

	offset, err := GetOrDefaultRowsConstraint(params.Limit, consts.OffsetDefault)

	errs = errors.Join(errs, err)
	if errs != nil {
		return nil, errs
	}

	return &gen.GetUsersParams{
		RowsLimit:  limit,
		RowsOffset: offset,
	}, nil
}

func ToCreateUserParams(params *repos.CreateUserParams) *gen.CreateUserParams {
	return &gen.CreateUserParams{
		Name:         params.Name,
		PasswordHash: params.PasswordHash,
		Email:        params.Email,
	}
}

func ToUpdateUserParams(params *repos.UpdateUserParams) (*gen.UpdateUserParams, error) {
	userID, err := utils.UintToInt[uint64, int64](params.ID)
	if err != nil {
		return nil, err
	}

	return &gen.UpdateUserParams{
		ID:           userID,
		Name:         params.Name,
		PasswordHash: params.PasswordHash,
		Email:        params.Email,
	}, nil
}

func ToUser(user *gen.User) *models.User {
	return &models.User{
		ID:           user.ID,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func ToUsers(users []*gen.User) []models.User {
	return lo.Map(users, func(user *gen.User, _ int) models.User {
		return *ToUser(user)
	})
}
