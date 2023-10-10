package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"gorm.io/gorm"
)

type expansion struct{}

func NewExpansion() repository.Expansion {
	return &expansion{}
}

func (r *expansion) InsertBatch(ctx context.Context, expansion_list []model.Expansion) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM EXPANSIONS").Error; err != nil {
			return err
		}

		if err := tx.CreateInBatches(expansion_list, len(expansion_list)).Error; err != nil {
			return err
		}

		return nil
	})
}
