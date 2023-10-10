package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type GeneratorInnerUnitCrawler struct {
	ICrawler
}

func NewGeneratorInnerUnitCrawler() ICrawler {
	return &GeneratorInnerUnitCrawler{
		ICrawler: NewAbstractCrawler(config.GENERATOR_PAGE, model.GENERATOR_INNER_UNIT_TYPE),
	}
}

func (c *GeneratorInnerUnitCrawler) JobName() string {
	return "GeneratorInnerUnitCrawlerJob"
}
