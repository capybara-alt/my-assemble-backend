package main

import (
	"os"

	"log/slog"

	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  core.DB_DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: gorm_logger.New(slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo), gorm_logger.Config{
			LogLevel: gorm_logger.Info,
			Colorful: true,
		}),
	})
	if err != nil {
		logger.Error("Failed to connect to database")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Migrator().AutoMigrate(&model.Expansion{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.Expansion{}, "expansion_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.Expansion{}, "expansion_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.Frame{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.Frame{}, "frame_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.Frame{}, "frame_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.InnerUnit{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.InnerUnit{}, "inner_unit_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.InnerUnit{}, "inner_unit_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.Weapon{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.Weapon{}, "weapon_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.Weapon{}, "weapon_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.ValidationUnitSchema{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.ValidationUnitSchema{}, "validation_unit_schema_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.ValidationUnitSchema{}, "validation_unit_schema_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.User{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.User{}, "user_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.User{}, "user_idx"); err != nil {
				return err
			}
		}

		if err := tx.Migrator().AutoMigrate(&model.Assembly{}); err != nil {
			return err
		}
		if hasIndex := tx.Migrator().HasIndex(&model.Assembly{}, "assembly_idx"); !hasIndex {
			if err := tx.Migrator().CreateIndex(&model.Assembly{}, "assembly_idx"); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		logger.Error("Failed to create table", "detail", err)
	}
}
