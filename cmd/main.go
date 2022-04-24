package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fyk7/code-snippets-app/app/config"
	"github.com/fyk7/code-snippets-app/app/di"
)

func main() {
	cfg := config.LoadConf()
	timeoutContext := time.Duration(cfg.AppTimeOut * time.Second)
	// Dependency Injection
	serviceContainer := di.Initialize(cfg, timeoutContext)

	snippet111, err := serviceContainer.SnippetService.GetByID(context.TODO(), 111)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(snippet111)

	// e := echo.New()
	// mw := _middleware.InitMiddleware()
	// e.Use(mw.CORS)
	// log.Fatal(e.Start(":8080"))
}
