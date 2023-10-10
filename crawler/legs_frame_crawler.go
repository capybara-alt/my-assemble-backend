package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type LegsFrameCrawler struct {
	ICrawler
}

func NewLegsFrameCrawler() ICrawler {
	return &LegsFrameCrawler{
		ICrawler: NewAbstractCrawler(config.LEGS_PARTS_PAGE, model.LEGS_UNIT_TYPE),
	}
}

func (c *LegsFrameCrawler) JobName() string {
	return "LegsFrameCrawlerJob"
}
