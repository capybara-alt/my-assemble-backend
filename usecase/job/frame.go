package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type FrameJob struct {
	crawlers  []crawler.ICrawler
	convertor convert.IConvertor[model.Frame]
	repo      repository.Frame
	logger    *slog.Logger
}

func NewFrameJob(
	crawlers []crawler.ICrawler,
	repo repository.Frame,
	convertor convert.IConvertor[model.Frame],
	logger *slog.Logger) ICrawlJobUsecase {
	return &FrameJob{
		crawlers:  crawlers,
		repo:      repo,
		convertor: convertor,
		logger:    logger,
	}
}

func (c *FrameJob) Execute(ctx context.Context) {
	models := []model.Frame{}

	for _, crawler := range c.crawlers {
		results, err := crawler.Crawl()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", crawler.JobName(), "detail", err)
		}
		c.logger.Info("Crawl successful", "JobName", crawler.JobName())
		frame_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "JobName", crawler.JobName(), "detail", err)
		} else {
			models = append(models, frame_list...)
			c.logger.Info("Convert successful", "JobName", crawler.JobName())
		}
	}

	if err := c.repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
