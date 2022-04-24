package usecase

import (
	"context"
	"time"

	"github.com/fyk7/code-snippets-app/app/domain/model"
	"github.com/fyk7/code-snippets-app/app/domain/repository"
)

type UserService interface {
	List(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, userID uint64) (model.User, error)
	GetByKeyWord(ctx context.Context, userName string) ([]model.User, error)
	Create(ctx context.Context, user model.User) error
	Update(ctx context.Context, user model.User) error
}

type userService struct {
	repo           repository.UserRepository
	contextTimeout time.Duration
}

func NewUserService(repo repository.UserRepository, timeout time.Duration) UserService {
	return &userService{
		repo:           repo,
		contextTimeout: timeout,
	}
}

func (ss *userService) List(ctx context.Context) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippets, err := ss.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ss *userService) GetByKeyWord(ctx context.Context, userName string) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippets, err := ss.repo.FindByName(ctx, userName)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func (ss *userService) GetByID(ctx context.Context, userID uint64) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	snippet, err := ss.repo.GetByID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	return snippet, nil
}

func (ss *userService) Create(ctx context.Context, snippet model.User) error {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	if err := ss.repo.Create(ctx, snippet); err != nil {
		return err
	}

	return nil
}

func (ss *userService) Update(ctx context.Context, snippet model.User) error {
	ctx, cancel := context.WithTimeout(ctx, ss.contextTimeout)
	defer cancel()

	if err := ss.repo.Create(ctx, snippet); err != nil {
		return err
	}

	return nil
}
