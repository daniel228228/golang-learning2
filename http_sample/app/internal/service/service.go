package service

import (
	"context"
	"errors"

	"http_sample/internal/config"
	"http_sample/internal/logger"
	"http_sample/internal/models/dto"
	"http_sample/internal/repo"
)

type Service interface {
	Get(int) (string, error)
	Write(int, string) error
	CalculateMatrix(*dto.Matrixes) (*dto.Matrix, error)
}

var (
	ErrBigMatrix            = errors.New("matrixes is too big")
	ErrIncompatibleMatrixes = errors.New("matrixes is incompatible")
	ErrInternalError        = errors.New("internal error")
)

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

func (s *service) CalculateMatrix(m *dto.Matrixes) (*dto.Matrix, error) {
	if m.A.N > 1000 || m.A.M > 1000 || m.B.N > 1000 || m.B.M > 1000 {
		return nil, ErrBigMatrix
	} else if m.A.N != m.B.M {
		return nil, ErrIncompatibleMatrixes
	}

	result := &dto.Matrix{
		N:    m.A.N,
		M:    m.B.M,
		Nums: make([][]int, m.A.N),
	}

	for i := 0; i < m.A.N; i++ {
		result.Nums[i] = make([]int, m.B.M)

		for j := 0; j < m.B.M; j++ {
			for k := 0; k < m.A.M; k++ {
				result.Nums[i][j] += m.A.Nums[i][k] * m.B.Nums[k][j]
			}
		}
	}

	return result, nil
}
