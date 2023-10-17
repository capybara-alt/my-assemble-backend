package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type headFrame struct{}

func NewHeadFrame() repository.ExternalFrame {
	return &headFrame{}
}

func (*headFrame) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.HEAD_PARTS_PAGE, model.HEAD_UNIT_TYPE)
}
