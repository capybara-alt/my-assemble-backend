package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type HeadFrameCrawler struct {
	ICrawler
}

func NewHeadFrameCrawler() ICrawler {
	return &HeadFrameCrawler{
		ICrawler: NewAbstractCrawler(config.HEAD_PARTS_PAGE, model.HEAD_UNIT_TYPE),
	}
}

func (c *HeadFrameCrawler) JobName() string {
	return "HeadFrameCrawlerJob"
}
