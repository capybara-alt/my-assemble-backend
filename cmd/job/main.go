package main

import (
	"context"
	"log/slog"
	"os"

	convert "github.com/capybara-alt/my-assemble/convert/job"
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/crawler"
	repo "github.com/capybara-alt/my-assemble/infrastructure/db"
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
	var weapon_crawlers = []crawler.ICrawler{
		crawler.NewCrossWeaponCrawler(),
		crawler.NewAmmoWeaponCrawler(),
		crawler.NewEnCoralWeaponCrawler(),
		crawler.NewExplodeWeaponCrawler(),
		crawler.NewExtraWeaponCrawler(),
		crawler.NewLauncherWeaponCrawler(),
		crawler.NewMissileWeaponCrawler(),
		crawler.NewOrbitTaletDroneWeaponCrawler(),
		crawler.NewShieldWeaponCrawler(),
	}
	weapon_repo := repo.NewWeapon()
	job.NewWeaponJob(weapon_crawlers, weapon_repo, weapon_convert, logger).Execute(ctx)

	frame_convert := convert.NewFrameList(validation_schema.GetFrameSchema())
	var frame_crawlers = []crawler.ICrawler{
		crawler.NewArmsFrameCrawler(),
		crawler.NewCoreFrameCrawler(),
		crawler.NewHeadFrameCrawler(),
		crawler.NewLegsFrameCrawler(),
		crawler.NewOtherLegsFrameCrawler(),
	}
	frame_repo := repo.NewFrame()
	job.NewFrameJob(frame_crawlers, frame_repo, frame_convert, logger).Execute(ctx)

	inner_unit_convert := convert.NewInnerUnitsList(validation_schema.GetInnerUnitsSchema())
	var inner_unit_crawlers = []crawler.ICrawler{
		crawler.NewBoosterInnerUnitCrawler(),
		crawler.NewFcsInnerUnitCrawler(),
		crawler.NewGeneratorInnerUnitCrawler(),
	}
	inner_unit_repo := repo.NewInnerUnit()
	job.NewInnerUnitJob(inner_unit_crawlers, inner_unit_repo, inner_unit_convert, logger).Execute(ctx)

	expansion_convert := convert.NewExpansionList(validation_schema.GetExpansionSchema())
	var expansion_crawlers = []crawler.ICrawler{
		crawler.NewExpansionCrawler(),
	}
	expansion_repo := repo.NewExpansion()
	job.NewExpansionJob(expansion_crawlers, expansion_repo, expansion_convert, logger).Execute(ctx)
}
