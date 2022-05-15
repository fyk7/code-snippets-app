package repository

import (
	"context"

	"github.com/fyk7/code-snippets-app/app/domain/model"
)

type SnippetRepository interface {
	GetAll(ctx context.Context) ([]model.Snippet, error)
	GetByID(ctx context.Context, id uint64) (model.Snippet, error)
	FindByKeyWord(ctx context.Context, keyword string) ([]model.Snippet, error)
	FindByTag(ctx context.Context, tagID uint64) ([]model.Snippet, error)
	AssociateWithTag(ctx context.Context, snippetID, tagID, userID int64) error
	Create(ctx context.Context, s model.Snippet, UserID uint64) error
	Update(ctx context.Context, s model.Snippet, UserID uint64) error
	Delete(ctx context.Context, id uint64) error
}
