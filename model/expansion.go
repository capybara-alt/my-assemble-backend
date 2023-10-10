package model

import "time"

type Expansion struct {
	Name                string `gorm:"primaryKey"`
	Category            string
	Unit                string
	EffectRange         int64
	Duration            int64
	AttackPower         int64
	BlastRadius         int64
	DirectHitAdjustment int64
	Resilience          int64
	Impact              int64
	AccumulativeImpact  int64
	CreatedAt           time.Time
}
