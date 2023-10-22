package model

import (
	"time"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
)

type InnerUnit struct {
	Name                  string `gorm:"primaryKey;index:inner_unit_idx"`
	Category              string
	Unit                  string
	Maker                 string
	ABThrust              int64
	AbEnConsumption       int64
	ENOutput              int64
	ENCapacity            int64
	EnergyFirearmSpec     int64
	ENRecharge            int64
	ENLoad                int64
	QBReloadIdealWeight   int64
	QBJetDuration         float64
	QBThrust              int64
	QBReloadTime          float64
	QbEnConsumption       int64
	UpwardThrust          int64
	UpwardENConsumption   int64
	MediumRangeAssist     int64
	SupplyRecovery        int64
	Price                 int64
	PostRecoveryENSupply  int64
	Thrust                int64
	MissileLockCorrection int64
	MultiLockCorrection   int64
	MeleeAttackThrust     int64
	MeleeAtkENConsump     int64
	CloseRangeAssist      int64
	LogRangeAssist        int64
	Weight                int64
	CreatedAt             time.Time
}

func (i *InnerUnit) ToPB() *myassemblyv1.InnerUnit {
	return &myassemblyv1.InnerUnit{
		Name:                  i.Name,
		Category:              i.Category,
		Unit:                  i.Unit,
		Maker:                 i.Maker,
		AbThrust:              int32(i.ABThrust),
		AbEnConsumption:       int32(i.AbEnConsumption),
		EnOutput:              int32(i.ENOutput),
		EnCapacity:            int32(i.ENCapacity),
		EnergyFirearmSpec:     int32(i.EnergyFirearmSpec),
		EnRecharge:            int32(i.ENRecharge),
		EnLoad:                int32(i.ENLoad),
		QbReloadIdealWeight:   int32(i.QBReloadIdealWeight),
		QbJetDuration:         float32(i.QBJetDuration),
		QbThrust:              int32(i.QBThrust),
		QbReloadTime:          float32(i.QBReloadTime),
		QbEnConsumption:       int32(i.QbEnConsumption),
		UpwardThrust:          int32(i.UpwardThrust),
		UpwardEnConsumption:   int32(i.UpwardENConsumption),
		MediumRangeAssist:     int32(i.MediumRangeAssist),
		SupplyRecovery:        int32(i.SupplyRecovery),
		Price:                 int32(i.Price),
		PostRecoveryEnSupply:  int32(i.PostRecoveryENSupply),
		Thrust:                int32(i.Thrust),
		MissileLockCorrection: int32(i.MissileLockCorrection),
		MultiLockCorrection:   int32(i.MultiLockCorrection),
		MeleeAttackThrust:     int32(i.MeleeAttackThrust),
		MeleeAtkEnConsump:     int32(i.MeleeAtkENConsump),
		CloseRangeAssist:      int32(i.CloseRangeAssist),
		LogRangeAssist:        int32(i.LogRangeAssist),
		Weight:                int32(i.Weight),
		CreatedAt:             i.CreatedAt.UnixMilli(),
	}
}

func (i *InnerUnit) FromPB(innerUnit *myassemblyv1.InnerUnit) *InnerUnit {
	return &InnerUnit{
		Name:                  innerUnit.Name,
		Category:              innerUnit.Category,
		Unit:                  innerUnit.Unit,
		Maker:                 innerUnit.Maker,
		ABThrust:              int64(innerUnit.AbThrust),
		AbEnConsumption:       int64(innerUnit.AbEnConsumption),
		ENOutput:              int64(innerUnit.EnOutput),
		ENCapacity:            int64(innerUnit.EnCapacity),
		EnergyFirearmSpec:     int64(innerUnit.EnergyFirearmSpec),
		ENRecharge:            int64(innerUnit.EnRecharge),
		ENLoad:                int64(innerUnit.EnLoad),
		QBReloadIdealWeight:   int64(innerUnit.QbReloadIdealWeight),
		QBJetDuration:         float64(innerUnit.QbJetDuration),
		QBThrust:              int64(innerUnit.QbThrust),
		QBReloadTime:          float64(innerUnit.QbReloadTime),
		QbEnConsumption:       int64(innerUnit.QbEnConsumption),
		UpwardThrust:          int64(innerUnit.UpwardThrust),
		UpwardENConsumption:   int64(innerUnit.UpwardEnConsumption),
		MediumRangeAssist:     int64(innerUnit.MediumRangeAssist),
		SupplyRecovery:        int64(innerUnit.SupplyRecovery),
		Price:                 int64(innerUnit.Price),
		PostRecoveryENSupply:  int64(innerUnit.PostRecoveryEnSupply),
		Thrust:                int64(innerUnit.Thrust),
		MissileLockCorrection: int64(innerUnit.MissileLockCorrection),
		MultiLockCorrection:   int64(innerUnit.MultiLockCorrection),
		MeleeAttackThrust:     int64(innerUnit.MeleeAttackThrust),
		MeleeAtkENConsump:     int64(innerUnit.MeleeAtkEnConsump),
		CloseRangeAssist:      int64(innerUnit.CloseRangeAssist),
		LogRangeAssist:        int64(innerUnit.LogRangeAssist),
		Weight:                int64(innerUnit.Weight),
		CreatedAt:             time.UnixMilli(innerUnit.CreatedAt),
	}
}
