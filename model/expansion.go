package model

import (
	"time"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
)

type Expansion struct {
	Name                string `gorm:"primaryKey;index:expansion_idx"`
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

func (e *Expansion) ToPB() *myassemblyv1.Expansion {
	return &myassemblyv1.Expansion{
		Name:                e.Name,
		Category:            e.Category,
		Unit:                e.Unit,
		EffectRange:         int32(e.EffectRange),
		Duration:            int32(e.Duration),
		AttackPower:         int32(e.AttackPower),
		BlastRadius:         int32(e.BlastRadius),
		DirectHitAdjustment: int32(e.DirectHitAdjustment),
		Resilience:          int32(e.Resilience),
		Impact:              int32(e.Impact),
		AccumulativeImpact:  int32(e.AccumulativeImpact),
		CreatedAt:           e.CreatedAt.UnixMilli(),
	}
}

func (e *Expansion) FromPB(expansion *myassemblyv1.Expansion) *Expansion {
	return &Expansion{
		Name:                expansion.Name,
		Category:            expansion.Category,
		Unit:                expansion.Unit,
		EffectRange:         int64(expansion.EffectRange),
		Duration:            int64(expansion.Duration),
		AttackPower:         int64(expansion.AttackPower),
		BlastRadius:         int64(expansion.BlastRadius),
		DirectHitAdjustment: int64(expansion.DirectHitAdjustment),
		Resilience:          int64(expansion.Resilience),
		Impact:              int64(expansion.Impact),
		AccumulativeImpact:  int64(expansion.AccumulativeImpact),
		CreatedAt:           time.UnixMilli(expansion.CreatedAt),
	}
}
