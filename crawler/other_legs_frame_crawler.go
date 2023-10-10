package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type OtherLegsFrameCrawler struct {
	ICrawler
}

func NewOtherLegsFrameCrawler() ICrawler {
	return &OtherLegsFrameCrawler{
		ICrawler: NewAbstractCrawler(config.LEGS_OTHER_PARTS_PAGE, model.LEGS_UNIT_TYPE),
	}
}

func (c *OtherLegsFrameCrawler) JobName() string {
	return "OthreLegsFrameCrawlerJob"
}
