package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
)

type LauncherWeaponCrawler struct {
	ICrawler
}

func NewLauncherWeaponCrawler() ICrawler {
	return &LauncherWeaponCrawler{
		ICrawler: NewAbstractCrawler(config.LAUNCHER_CANNON_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE),
	}
}

func (c *LauncherWeaponCrawler) JobName() string {
	return "LauncherCannonWeaponCrawlerJob"
}
