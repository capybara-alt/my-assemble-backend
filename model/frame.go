package model

import "time"

type Frame struct {
	Name                  string `gorm:"primaryKey"`
	Category              string
	Unit                  string
	ABThrust              int64
	AbEnConsumption       int64
	AP                    int64
	ENLoad                int64
	QBReloadIdealWeight   int64
	QBReloadTime          float64
	QBJetDuration         float64
	QBThrust              int64
	QbEnConsumption       int64
	SystemRecovery        int64
	GeneratorSupplyAdj    int64
	GeneratorOutputAdj    int64
	ScanStandbyTime       float64
	ScanEffectDuration    float64
	ScanDistance          int64
	BoosterEfficiencyAdj  int64
	Maker                 string
	UpwardThrust          int64
	UpwardENConsumption   int64
	Price                 int64
	RecoilControl         int64
	JumpHeight            int64
	AttitudeStability     int64
	FirearmSpecialization int64
	Thrust                int64
	JumpDistance          int64
	LoadLimit             int64
	AntiEnergyDefense     int64
	AntiKineticDefense    int64
	AntiExplosiveDefense  int64
	ArmsLoadLimit         int64
	TravelSpeed           int64
	MeleeSpecialization   int64
	Weight                int64
	HighSpeedPref         int64
	CreatedAt             time.Time
}
