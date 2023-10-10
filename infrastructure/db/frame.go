package db

import (
	"context"
	"errors"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"github.com/capybara-alt/my-assemble/repository"
	"gorm.io/gorm"
)

type frame struct{}

func NewFrame() repository.Frame {
	return &frame{}
}

func (r *frame) InsertBatch(ctx context.Context, frame_list []model.Frame) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM FRAMES").Error; err != nil {
			return err
		}

		if err := tx.CreateInBatches(frame_list, len(frame_list)).Error; err != nil {
			return err
		}

		return nil
	})
}
