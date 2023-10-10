package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type CoreFrameCrawler struct {
	ICrawler
}

func NewCoreFrameCrawler() ICrawler {
	return &CoreFrameCrawler{
		ICrawler: NewAbstractCrawler(config.CORE_PARTS_PAGE, model.CORE_UNIT_TYPE),
	}
}

func (c *CoreFrameCrawler) JobName() string {
	return "CoreFrameCrawlerJob"
}
