package repository

import (
	"context"

	"github.com/fyk7/code-snippets-app/app/domain/model"
)

type TagRepository interface {
	GetAll(ctx context.Context) ([]model.Tag, error)
	GetByID(ctx context.Context, tagID uint64) (model.Tag, error)
	FindByKeyWord(ctx context.Context, keyword string) ([]model.Tag, error)
	Create(ctx context.Context, tag model.Tag, UserID uint64) error
}
