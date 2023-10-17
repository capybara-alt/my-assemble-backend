package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type ammoWeapon struct{}

func NewAmmoWeapon() repository.ExternalWeapon {
	return &ammoWeapon{}
}

func (*ammoWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.AMMO_WEAPON_PAGE, model.DEFAULT_WEAPON_ARM_UNIT_TYPE)
}
