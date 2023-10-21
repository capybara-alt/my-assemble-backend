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

type frame struct{}

func NewFrame() repository.Frame {
	return &frame{}
}

func (r *frame) UpsertBatch(ctx context.Context, frameList []model.Frame) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, v := range frameList {
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

func (r *frame) Get(ctx context.Context, key string) (*model.Frame, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	frame := &model.Frame{}
	if err := db.First(&frame, key).Error; err != nil {
		return nil, err
	}
	return frame, nil
}

func (r *frame) GetAll(ctx context.Context) ([]model.Frame, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	frameList := []model.Frame{}
	if err := db.Find(&frameList).Error; err != nil {
		return nil, err
	}

	return frameList, nil
}
