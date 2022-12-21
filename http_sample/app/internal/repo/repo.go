package repo

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"http_sample/internal/config"
	"http_sample/internal/logger"
)

type Repo interface {
	Get(int) (string, error)
	Write(int, string) error
}

const (
	ConnRetries  = 5
	GormLogLevel = gormLogger.Info // TODO: Change log level from Info to Silent
)

type repo struct {
	config *config.Config
	log    logger.Logger
	db     *gorm.DB
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

	period := time.Second / 2

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for i := 0; i < ConnRetries; i++ {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}

		r.log.Printf("attempt %d: connecting to DB", i+1)

		if err := r.dbConnect(); err != nil {
			r.log.Printf("attempt %d failed: repository error: %s", i+1, err.Error())

			if i+1 < ConnRetries {
				period *= 2
				ticker.Reset(period)
			} else {
				return err
			}
		} else {
			r.log.Print("connected to DB")
			return nil
		}
	}

	return nil
}

func (r *repo) dbConnect() error {
	if db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s", r.config.DB.User, r.config.DB.Pass, r.config.DB.Host, r.config.DB.Port, r.config.DB.Name),
		),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(GormLogLevel),
		}); err != nil {
		return err
	} else {
		r.db = db
		return nil
	}
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
