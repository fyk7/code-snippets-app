package usecase

import (
	"context"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
)

type TagService interface {
	List(ctx context.Context) ([]model.Tag, error)
	GetByID(ctx context.Context, id uint64) (model.Tag, error)
	GetByKeyWord(ctx context.Context, keyword string) ([]model.Tag, error)
	Create(ctx context.Context, tag model.Tag, UserID uint64) error
}

type tagService struct {
	repo           repository.TagRepository
	contextTimeout time.Duration
}

func NewTagService(repo repository.TagRepository, timeout time.Duration) TagService {
	return &tagService{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (ts *tagService) List(ctx context.Context) ([]model.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	snippets, err := ts.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ts *tagService) GetByID(ctx context.Context, id uint64) (model.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	snippet, err := ts.repo.GetByID(ctx, id)
	if err != nil {
		return model.Tag{}, err
	}

	return snippet, nil
}

func (ts *tagService) GetByKeyWord(ctx context.Context, keyword string) ([]model.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	snippets, err := ts.repo.FindByKeyWord(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ts *tagService) Create(ctx context.Context, snippet model.Tag, UserID uint64) error {
	ctx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()

	if err := ts.repo.Create(ctx, snippet, UserID); err != nil {
		return err
	}

	return nil
}
