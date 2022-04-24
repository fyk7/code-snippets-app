package di

import "github.com/fyk7/code-snippets-app/app/usecase"

type ServiceContainer struct {
	SnippetService usecase.SnippetService
	TagService     usecase.TagService
	UserService    usecase.UserService
}

func NewServiceContainer(
	snippetService usecase.SnippetService,
	tagService usecase.TagService,
	userService usecase.UserService,
) *ServiceContainer {
	return &ServiceContainer{
		SnippetService: snippetService,
		TagService:     tagService,
		UserService:    userService,
	}
}
