package app

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"http_sample/internal/config"
	"http_sample/internal/err_chan"

	"http_sample/internal/logger"
)

var App *application

func Run(ctx context.Context, logger logger.Logger, config *config.Config) {
	App = &application{
		ctx:    ctx,
		config: config,
		logger: logger,
	}

	// repo
	repoImpl := repo.NewRepo(App.config, App.logger)
	App.Repo.ComponentControl = repoImpl
	App.Repo.Impl = repoImpl

	if err := App.Repo.Start(App.ctx); err != nil {
		err_chan.New(err)
	}
	defer shutdown(App.Repo.Shutdown)

	// service
	serviceImpl := service.NewService(App.config, App.logger, App.Repo.Impl)
	App.Service.ComponentControl = serviceImpl
	App.Service.Impl = serviceImpl

	if err := App.Service.Start(App.ctx); err != nil {
		err_chan.New(err)
	}
	defer shutdown(App.Service.Shutdown)

	// http_handler
	httpHandlerImpl := http_handler.NewHttpHandler(App.config, App.logger, App.Service.Impl)
	App.HttpHandler.ServerControl = httpHandlerImpl
	App.HttpHandler.Impl = httpHandlerImpl

	serve(App.HttpHandler.Serve)
	defer shutdown(App.HttpHandler.Shutdown)

	<-ctx.Done()
}

func shutdown(shutdown func(ctx context.Context) error) {
	if err := shutdown(context.Background()); err != nil &&
		!errors.Is(err, context.Canceled) &&
		!errors.Is(err, context.DeadlineExceeded) {
		err_chan.New(fmt.Errorf("finished %T with error: %w", shutdown, err))
	}
}

func serve(server func(ctx context.Context) error) {
	var cancel context.CancelFunc
	App.ctx, cancel = context.WithCancel(App.ctx)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				App.logger.Fatalf("panic: %v\n\n%s", r, string(debug.Stack()))
				cancel()
			}
		}()

		if err := server(context.Background()); err != nil {
			err_chan.New(fmt.Errorf("error serving %T: %w", server, err))
		}

		cancel()
	}()
}
