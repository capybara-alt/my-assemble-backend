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

type innerUnit struct{}

func NewInnerUnit() repository.InnerUnit {
	return &innerUnit{}
}

func (r *innerUnit) UpsertBatch(ctx context.Context, innerUnitList []model.InnerUnit) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, v := range innerUnitList {
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

func (r *innerUnit) Get(ctx context.Context, key string) (*model.InnerUnit, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	innerUnit := &model.InnerUnit{}
	if err := db.First(&innerUnit, key).Error; err != nil {
		return nil, err
	}
	return innerUnit, nil
}

func (r *innerUnit) GetAll(ctx context.Context) ([]model.InnerUnit, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	innerUnitList := []model.InnerUnit{}
	if err := db.Find(&innerUnitList).Error; err != nil {
		return nil, err
	}

	return innerUnitList, nil
}
