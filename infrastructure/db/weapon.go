package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"gorm.io/gorm"
)

type weapon struct{}

func NewWeapon() repository.Weapon {
	return &weapon{}
}

func (r *weapon) InsertBatch(ctx context.Context, weapon_list []model.Weapon) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM WEAPONS").Error; err != nil {
			return err
		}

		if err := tx.CreateInBatches(weapon_list, len(weapon_list)).Error; err != nil {
			return err
		}

		return nil
	})
}
