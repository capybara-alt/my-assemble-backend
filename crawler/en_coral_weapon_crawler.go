package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type EnCoralWeaponCrawler struct {
	ICrawler
}

func NewEnCoralWeaponCrawler() ICrawler {
	return &EnCoralWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.EN_CORAL_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE),
	}
}

func (c *EnCoralWeaponCrawler) JobName() string {
	return "EnCoralWeaponCrawlerJob"
}
