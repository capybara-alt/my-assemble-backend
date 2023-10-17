package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type ExpansionList struct {
	validationSchema model.ValidationUnitSchemaList
}

func NewExpansionList(s model.ValidationUnitSchemaList) IConvertor[model.Expansion] {
	return &ExpansionList{
		validationSchema: s,
	}
}

func (e *ExpansionList) Convert(results model.CrawlResultJSON) ([]model.Expansion, error) {
	expansionList := []model.Expansion{}
	unitType := results.GetUnitType()
	categories := results[unitType]
	for cname, category := range categories {
		wc := newExpansion(cname, unitType, e.validationSchema)
		for name, info := range category {
			expansion, err := wc.convert(name, info)
			if err != nil {
				return nil, err
			}
			expansionList = append(expansionList, *expansion)
		}
	}

	return expansionList, nil
}

type expansion struct {
	validationSchema model.ValidationUnitSchemaList
	unitType         string
	category         string
}

func newExpansion(category, unitType string, s model.ValidationUnitSchemaList) *expansion {
	return &expansion{
		category:         category,
		unitType:         unitType,
		validationSchema: s,
	}
}

func (e *expansion) convert(name string, info model.UnitInfoJSON) (*model.Expansion, error) {
	converted, err := e.validationSchema.ConvertValues(info)
	if err != nil {
		return nil, err
	}

	return &model.Expansion{
		Name:                name,
		Category:            e.category,
		Unit:                e.unitType,
		EffectRange:         converted["effect_range"].(int64),
		Duration:            converted["duration"].(int64),
		AttackPower:         converted["attack_power"].(int64),
		BlastRadius:         converted["blast_radius"].(int64),
		DirectHitAdjustment: converted["direct_hit_adjustment"].(int64),
		Resilience:          converted["resilience"].(int64),
		Impact:              converted["impact"].(int64),
		AccumulativeImpact:  converted["accumulative_impact"].(int64),
		CreatedAt:           core.Now(),
	}, nil
}
