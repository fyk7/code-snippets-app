package usecase

import (
	"context"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
)

type SnippetService interface {
	List(ctx context.Context) ([]model.Snippet, error)
	GetByID(ctx context.Context, id uint64) (model.Snippet, error)
	GetByKeyWord(ctx context.Context, keyword string) ([]model.Snippet, error)
	GetByKeyTagID(ctx context.Context, tagID uint64) ([]model.Snippet, error)
	Create(ctx context.Context, snippet model.Snippet, UserID uint64) error
	Update(ctx context.Context, snippet model.Snippet, UserID uint64) error
	Delete(ctx context.Context, id uint64) error
}

type snippetService struct {
	repo           repository.SnippetRepository
	contextTimeout time.Duration
}

func NewSnippetService(repo repository.SnippetRepository, timeout time.Duration) SnippetService {
	return &snippetService{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (ss *snippetService) List(ctx context.Context) ([]model.Snippet, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippets, err := ss.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ss *snippetService) GetByID(ctx context.Context, id uint64) (model.Snippet, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippet, err := ss.repo.GetByID(ctx, id)
	if err != nil {
		return model.Snippet{}, err
	}

	return snippet, nil
}

func (ss *snippetService) GetByKeyWord(ctx context.Context, keyword string) ([]model.Snippet, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippets, err := ss.repo.FindByKeyWord(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ss *snippetService) GetByKeyTagID(ctx context.Context, tagID uint64) ([]model.Snippet, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippets, err := ss.repo.FindByTag(ctx, tagID)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ss *snippetService) Create(ctx context.Context, snippet model.Snippet, UserID uint64) error {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	if err := ss.repo.Create(ctx, snippet, UserID); err != nil {
		return err
	}

	return nil
}

func (ss *snippetService) Update(ctx context.Context, snippet model.Snippet, UserID uint64) error {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	if err := ss.repo.Update(ctx, snippet, UserID); err != nil {
		return err
	}

	return nil
}

func (ss *snippetService) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	if err := ss.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
