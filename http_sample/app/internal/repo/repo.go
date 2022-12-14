package repo

import (
	"context"

	"http_sample/internal/config"
	"http_sample/internal/logger"
)

type Repo interface {
	Get(int) (string, error)
	Write(int, string) error
}

type repo struct {
	config *config.Config
	log    logger.Logger
}

func NewRepo(config *config.Config, logger logger.Logger) *repo {
	return &repo{
		config: config,
		log:    logger,
	}
}

func (r *repo) Start(ctx context.Context) error {
	r.log.Print("starting Repo")
	defer r.log.Print("started Repo")

	// noop

	return nil
}

func (r *repo) Shutdown(ctx context.Context) error {
	r.log.Print("stopping Repo")
	defer r.log.Print("stopped Repo")

	// noop

	return nil
}

func (r *repo) Get(id int) (string, error) {
	return "", nil
}

func (r *repo) Write(id int, value string) error {
	return nil
}
