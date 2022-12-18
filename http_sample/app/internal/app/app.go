package app

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"http_sample/internal/config"
	"http_sample/internal/errset"
	"http_sample/internal/repo"

	"http_sample/internal/logger"
)

var App *application

func init() {
	App = &application{}
}

func Run(ctx context.Context, logger logger.Logger, config *config.Config) {
	App.ctx = ctx
	App.logger = logger
	App.config = config

	// repo
	select {
	case <-App.ctx.Done():
		return
	default:
	}
	repoImpl := repo.NewRepo(App.config, App.logger)
	App.Repo.ComponentControl = repoImpl
	App.Repo.Impl = repoImpl

	if err := App.Repo.Start(App.ctx); err != nil {
		errset.New(fmt.Errorf("%q error: %w", "repo", err))
	}
	defer shutdown(App.Repo.Shutdown, "repo")

	// service
	select {
	case <-App.ctx.Done():
		return
	default:
	}
	serviceImpl := service.NewService(App.config, App.logger, App.Repo.Impl)
	App.Service.ComponentControl = serviceImpl
	App.Service.Impl = serviceImpl

	if err := App.Service.Start(App.ctx); err != nil {
		errset.New(fmt.Errorf("%q error: %w", "service", err))
	}
	defer shutdown(App.Service.Shutdown, "service")

	// http_handler
	select {
	case <-App.ctx.Done():
		return
	default:
	}
	httpHandlerImpl := http_handler.NewHttpHandler(App.config, App.logger, App.Service.Impl)
	App.HttpHandler.ServerControl = httpHandlerImpl
	App.HttpHandler.Impl = httpHandlerImpl

	serve(App.HttpHandler.Serve, "http_handler")
	defer shutdown(App.HttpHandler.Shutdown, "http_handler")

	<-App.ctx.Done()
}

func shutdown(shutdown func(ctx context.Context) error, name string) {
	if err := shutdown(App.ctx); err != nil &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, context.DeadlineExceeded) {
		errset.New(fmt.Errorf("finished %q with error: %w", name, err))
	}
}

func serve(server func(ctx context.Context) error, name string) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				App.logger.Fatalf("panic: %v\n\n%s", r, string(debug.Stack()))
				errset.New(fmt.Errorf("panic in %q", name))
			}
		}()

		if err := server(App.ctx); err != nil {
			errset.New(fmt.Errorf("error serving %q: %w", name, err))
		}
	}()
}
