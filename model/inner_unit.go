package model

import "time"

type InnerUnit struct {
	Name                  string `gorm:"primaryKey"`
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
