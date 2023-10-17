package main

import (
	"context"
	"log/slog"
	"os"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/infrastructure/crawler"
	repo "github.com/capybara-alt/my-assemble/infrastructure/db"
	"github.com/capybara-alt/my-assemble/repository"
	"github.com/capybara-alt/my-assemble/usecase/common"
	"github.com/capybara-alt/my-assemble/usecase/job"
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
	ctx := context.Background()
	ctx = core.SetTx(ctx, db)

	validation_schema_repo := repo.NewValidationUnitSchema()
	validation_schema := common.NewValidationSchema(validation_schema_repo, logger)
	validation_schema.Execute(ctx)

	weapon_convert := convert.NewWeaponList(validation_schema.GetWeaponSchema())
	var external_weapon_repo = []repository.ExternalWeapon{
		crawler.NewCrossWeapon(),
		crawler.NewAmmoWeapon(),
		crawler.NewEnCoralWeapon(),
		crawler.NewExplodeWeapon(),
		crawler.NewExtraWeapon(),
		crawler.NewLauncherWeapon(),
		crawler.NewMissileWeapon(),
		crawler.NewOrbitTaletDroneWeapon(),
		crawler.NewShieldWeapon(),
	}
	weapon_repo := repo.NewWeapon()
	job.NewWeaponJob(weapon_repo, external_weapon_repo, weapon_convert, logger).Execute(ctx)

	frame_convert := convert.NewFrameList(validation_schema.GetFrameSchema())
	var external_frame_repo = []repository.ExternalFrame{
		crawler.NewArmsFrame(),
		crawler.NewCoreFrame(),
		crawler.NewHeadFrame(),
		crawler.NewLegsFrame(),
		crawler.NewOtherLegsFrame(),
	}
	frame_repo := repo.NewFrame()
	job.NewFrameJob(frame_repo, external_frame_repo, frame_convert, logger).Execute(ctx)

	inner_unit_convert := convert.NewInnerUnitsList(validation_schema.GetInnerUnitsSchema())
	var external_inner_units_repo = []repository.ExternalInnerUnit{
		crawler.NewBoosterInnerUnit(),
		crawler.NewFcsInnerUnit(),
		crawler.NewGeneratorInnerUnit(),
	}
	inner_unit_repo := repo.NewInnerUnit()
	job.NewInnerUnitJob(inner_unit_repo, external_inner_units_repo, inner_unit_convert, logger).Execute(ctx)

	expansion_convert := convert.NewExpansionList(validation_schema.GetExpansionSchema())
	var external_expansion_repo = []repository.ExternalExpansion{
		crawler.NewExpansion(),
	}
	expansion_repo := repo.NewExpansion()
	job.NewExpansionJob(expansion_repo, external_expansion_repo, expansion_convert, logger).Execute(ctx)
}
