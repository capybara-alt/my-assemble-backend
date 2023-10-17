package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type weapon struct{}

func NewWeapon() repository.Weapon {
	return &weapon{}
}

func (r *weapon) UpsertBatch(ctx context.Context, weaponList []model.Weapon) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, v := range weaponList {
			err := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&v).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
