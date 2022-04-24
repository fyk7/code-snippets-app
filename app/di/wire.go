//go:build wireinject
// +build wireinject

package di

import (
	"time"

	"github.com/fyk7/code-snippets-app/app/config"
	"github.com/fyk7/code-snippets-app/app/infrastructure/database"
	_repository "github.com/fyk7/code-snippets-app/app/interface_adapter/repository"
	"github.com/fyk7/code-snippets-app/app/usecase"
	"github.com/google/wire"
)

func Initialize(cfg *config.Config, timeout time.Duration) *ServiceContainer {
	wire.Build(
		database.NewDB,
		_repository.NewSnippetRepository,
		_repository.NewTagRepository,
		_repository.NewUserRepository,
		usecase.NewSnippetService,
		usecase.NewTagService,
		usecase.NewUserService,
		NewServiceContainer,
	)
	return nil
}
