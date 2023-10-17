package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type FrameJob struct {
	convertor    convert.IConvertor[model.Frame]
	dbRepo       repository.Frame
	externalRepo []repository.ExternalFrame
	logger       *slog.Logger
}

func NewFrameJob(
	dbRepo repository.Frame,
	externalRepo []repository.ExternalFrame,
	convertor convert.IConvertor[model.Frame],
	logger *slog.Logger) ICrawlJobUsecase {
	return &FrameJob{
		dbRepo:       dbRepo,
		externalRepo: externalRepo,
		convertor:    convertor,
		logger:       logger,
	}
}

func (c *FrameJob) Execute(ctx context.Context) {
	models := []model.Frame{}

	for _, repo := range c.externalRepo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "detail", err)
		}
		c.logger.Info("Crawl successful")
		frameList, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, frameList...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.dbRepo.UpsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
