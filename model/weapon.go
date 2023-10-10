package model

import (
	"time"
)

type Weapon struct {
	Name                 string `gorm:"primaryKey"`
	Category             string
	Maker                string
	ENLoad               int64
	Price                int64
	Weight               int64
	Unit                 string
	PAInterference       int64
	Cooling              int64
	Recoil               int64
	TKHeatBuildup        int64
	DeploymentRange      int64
	AmmunitionCost       string
	MagazineRounds       int64
	IdealRange           int64
	AttackPower          string
	ATKHeatBuildup       int64
	DamageMitigation     int64
	ImpactDampening      int64
	IdleDamageMitigation int64
	IdleTime             float64
	IdleImpactMitigation int64
	ReloadTime           float64
	MaxLockCount         int64
	EffectiveRange       int64
	BlastRadius          int64
	DirectHitAdjustment  int64
	TotalRounds          int64
	Impact               string
	AccumulativeImpact   string
	Guidance             int64
	HomingLockTime       float64
	RapidFire            float64
	ConsecutiveHits      int64
	IGDuration           float64
	IGDamageMitigation   int64
	IGImpactDampening    int64
	ChgEnLoad            int64
	ChgTime              float64
	ChgAttackPower       string
	ChgHeatBuildup       int64
	ChgAmmoConsumption   int64
	ChgBlastRadius       int64
	ChgImpact            string
	ChgAccumImpact       string
	FullChgAttackPower   string
	FullChgHeatBuildup   int64
	FullChgAmmoConsump   int64
	FullChgBlastRadius   int64
	FullChgImpact        string
	FullChgAccumImpact   string
	FullChgTime          float64
	CreatedAt            time.Time
}
