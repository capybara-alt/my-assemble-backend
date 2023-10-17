package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type missileWeapon struct{}

func NewMissileWeapon() repository.ExternalWeapon {
	return &missileWeapon{}
}

func (*missileWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.MISSILE_WEAPON_PAGE, model.DEFAULT_WEAPON_BACK_UNIT_TYPE)
}
