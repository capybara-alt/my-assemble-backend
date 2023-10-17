package crawler

import (
	"github.com/capybara-alt/my-assemble/config"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
)

type shieldWeapon struct{}

func NewShieldWeapon() repository.ExternalWeapon {
	return &shieldWeapon{}
}

func (*shieldWeapon) Fetch() (model.CrawlResultJSON, error) {
	return repository.Crawl(config.SHIELD_WEAPON_PAGE, model.SHIELD_WEAPON_UNIT_TYPE)
}
