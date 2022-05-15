package main

import (
	"log"
	"time"

	"github.com/fyk7/code-snippets-app/app/config"
	"github.com/fyk7/code-snippets-app/app/di"
	_handler "github.com/fyk7/code-snippets-app/app/interface_adapter/handler"
	_middleware "github.com/fyk7/code-snippets-app/app/interface_adapter/handler/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConf()
	timeoutContext := time.Duration(cfg.AppTimeOut * time.Second)
	// Dependency Injection
	serviceContainer := di.Initialize(cfg, timeoutContext)

	e := echo.New()
	mw := _middleware.InitMiddleware()
	e.Use(mw.CORS)
	// Register handlers.
	_handler.NewSnippetHandler(e, serviceContainer.SnippetService)
	_handler.NewTagHandler(e, serviceContainer.TagService)
	log.Fatal(e.Start(":8080"))
}
