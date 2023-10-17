package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type launcherWeapon struct{}

func NewLauncherWeapon() repository.ExternalWeapon {
	return &launcherWeapon{}
}

func (*launcherWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.LAUNCHER_CANNON_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE)
}
