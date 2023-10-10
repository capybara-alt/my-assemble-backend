package job

import (
	"context"
	"log/slog"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/crawler"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type WeaponJob struct {
	crawlers  []crawler.ICrawler
	repo      repository.Weapon
	convertor convert.IConvertor[model.Weapon]
	logger    *slog.Logger
}

func NewWeaponJob(
	crawlers []crawler.ICrawler,
	repo repository.Weapon,
	convertor convert.IConvertor[model.Weapon],
	logger *slog.Logger) ICrawlJobUsecase {
	return &WeaponJob{
		crawlers:  crawlers,
		repo:      repo,
		convertor: convertor,
		logger:    logger,
	}
}

func (c *WeaponJob) Execute(ctx context.Context) {
	models := []model.Weapon{}

	for _, crawler := range c.crawlers {
		results, err := crawler.Crawl()
		if err != nil {
			c.logger.Error("Crawl failed", "JobName", crawler.JobName(), "detail", err)
		}
		c.logger.Info("Crawl successful", "JobName", crawler.JobName())
		weapon_list, err := c.convertor.Convert(results)
		if err != nil {
			c.logger.Error("Validation error", "JobName", crawler.JobName(), "detail", err)
		} else {
			models = append(models, weapon_list...)
			c.logger.Info("Convert successful", "JobName", crawler.JobName())
		}
	}

	if err := c.repo.InsertBatch(ctx, models); err != nil {
		c.logger.Error("InsertBatch failed", "detail", err)
	}
}
