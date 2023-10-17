package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type armsFrame struct{}

func NewArmsFrame() repository.ExternalFrame {
	return &armsFrame{}
}

func (*armsFrame) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.ARMS_PARTS_PAGE, model.ARMS_UNIT_TYPE)
}
