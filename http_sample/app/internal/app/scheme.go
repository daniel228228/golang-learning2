package app

import (
	"context"

	"http_sample/internal/config"
	"http_sample/internal/logger"

	http_handler "http_sample/internal/handlers/http"
	"http_sample/internal/repo"
	"http_sample/internal/service"
)

type ComponentControl interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type ServerControl interface {
	Serve(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type application struct {
	ctx    context.Context
	config *config.Config
	logger logger.Logger

	Repo struct {
		ComponentControl
		Impl repo.Repo
	}
	Service struct {
		ComponentControl
		Impl service.Service
	}
	HttpHandler struct {
		ServerControl
		Impl http_handler.HttpHandler
	}
}
