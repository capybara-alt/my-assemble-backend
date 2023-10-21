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

type expansion struct{}

func NewExpansion() repository.Expansion {
	return &expansion{}
}

func (r *expansion) UpsertBatch(ctx context.Context, expansionList []model.Expansion) error {
	db := core.GetTx(ctx)
	if db == nil {
		return errors.New("DB not connected")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, v := range expansionList {
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

func (r *expansion) Get(ctx context.Context, key string) (*model.Expansion, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	expansion := &model.Expansion{}
	if err := db.First(&expansion, key).Error; err != nil {
		return nil, err
	}
	return expansion, nil
}

func (r *expansion) GetAll(ctx context.Context) ([]model.Expansion, error) {
	db := core.GetTx(ctx)
	if db == nil {
		return nil, errors.New("DB not connected")
	}

	expansionList := []model.Expansion{}
	if err := db.Find(&expansionList).Error; err != nil {
		return nil, err
	}

	return expansionList, nil
}
