package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type legsFrame struct{}

func NewLegsFrame() repository.ExternalFrame {
	return &legsFrame{}
}

func (*legsFrame) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.LEGS_PARTS_PAGE, model.LEGS_UNIT_TYPE)
}
