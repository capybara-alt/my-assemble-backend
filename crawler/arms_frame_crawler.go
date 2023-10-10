package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type ArmsFrameCrawler struct {
	ICrawler
}

func NewArmsFrameCrawler() ICrawler {
	return &ArmsFrameCrawler{
		ICrawler: NewAbstractCrawler(config.ARMS_PARTS_PAGE, model.ARMS_UNIT_TYPE),
	}
}

func (c *ArmsFrameCrawler) JobName() string {
	return "ArmsFrameCrawlerJob"
}
