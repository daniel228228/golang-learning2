package service

import (
	"context"

	"http_sample/internal/config"
	"http_sample/internal/logger"
	"http_sample/internal/repo"
)

type Service interface {
	Get(int) (string, error)
	Write(int, string) error
}

type service struct {
	config *config.Config
	log    logger.Logger
	repo   repo.Repo
}

func NewService(config *config.Config, log logger.Logger, repo repo.Repo) *service {
	return &service{
		config: config,
		log:    log,
		repo:   repo,
	}
}

func (s *service) Start(ctx context.Context) error {
	s.log.Print("starting Service")
	defer s.log.Print("started Service")

	// noop

	return nil
}

func (s *service) Shutdown(ctx context.Context) error {
	s.log.Print("stopping Service")
	defer s.log.Print("stopped Service")

	// noop

	return nil
}

func (s *service) Get(id int) (string, error) {
	return s.repo.Get(id)
}

func (s *service) Write(id int, value string) error {
	return s.repo.Write(id, value)
}
