package repository

import (
	"context"

	"github.com/fyk7/code-snippets-app/app/domain/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, userID uint64) (model.User, error)
	FindByName(ctx context.Context, keyword string) ([]model.User, error)
	Create(ctx context.Context, user model.User) error
	Update(ctx context.Context, user model.User) error
}
