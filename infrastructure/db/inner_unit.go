package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"gorm.io/gorm"
)

type inner_unit struct{}

func NewInnerUnit() repository.InnerUnit {
	return &inner_unit{}
}

func (r *inner_unit) InsertBatch(ctx context.Context, inner_unit_list []model.InnerUnit) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM INNER_UNITS").Error; err != nil {
			return err
		}

		if err := tx.CreateInBatches(inner_unit_list, len(inner_unit_list)).Error; err != nil {
			return err
		}

		return nil
	})
}
