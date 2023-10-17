package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type boosterInnerUnit struct{}

func NewBoosterInnerUnit() repository.ExternalInnerUnit {
	return &boosterInnerUnit{}
}

func (*boosterInnerUnit) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.BOOSTER_PAGE, model.BOOSTER_INNER_UNIT_TYPE)
}
