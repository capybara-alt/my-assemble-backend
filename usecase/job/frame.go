package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type FrameJob struct {
	convertor     convert.IConvertor[model.Frame]
	db_repo       repository.Frame
	external_repo []repository.ExternalFrame
	logger        *slog.Logger
}

func NewFrameJob(
	db_repo repository.Frame,
	external_repo []repository.ExternalFrame,
	convertor convert.IConvertor[model.Frame],
	logger *slog.Logger) ICrawlJobUsecase {
	return &FrameJob{
		db_repo:       db_repo,
		external_repo: external_repo,
		convertor:     convertor,
		logger:        logger,
	}
}

func (c *FrameJob) Execute(ctx context.Context) {
	models := []model.Frame{}

	for _, repo := range c.external_repo {
		results, err := repo.Fetch()
		if err != nil {
			c.logger.Error("Crawl failed", "detail", err)
		}
		c.logger.Info("Crawl successful")
		frame_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "detail", err)
		} else {
			models = append(models, frame_list...)
			c.logger.Info("Convert successful")
		}
	}

	if err := c.db_repo.UpsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
