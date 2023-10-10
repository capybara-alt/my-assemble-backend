package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type ExpansionList struct {
	validation_schema model.ValidationUnitSchemaList
}

func NewExpansionList(s model.ValidationUnitSchemaList) IConvertor[model.Expansion] {
	return &ExpansionList{
		validation_schema: s,
	}
}

func (e *ExpansionList) Convert(results model.CrawlResultJSON) ([]model.Expansion, error) {
	expansion_list := []model.Expansion{}
	unit_type := results.GetUnitType()
	categories := results[unit_type]
	for category_name, category := range categories {
		wc := newExpansion(category_name, unit_type, e.validation_schema)
		for name, expansion_info := range category {
			expansion, err := wc.convert(name, expansion_info)
			if err != nil {
				return nil, err
			}
			expansion_list = append(expansion_list, *expansion)
		}
	}

	return expansion_list, nil
}

type expansion struct {
	validation_schema model.ValidationUnitSchemaList
	unit_type         string
	category          string
}

func newExpansion(category, unit_type string, s model.ValidationUnitSchemaList) *expansion {
	return &expansion{
		category:          category,
		unit_type:         unit_type,
		validation_schema: s,
	}
}

func (e *expansion) convert(name string, expansion_info model.UnitInfoJSON) (*model.Expansion, error) {
	converted, err := e.validation_schema.ConvertValues(expansion_info)
	if err != nil {
		return nil, err
	}

	return &model.Expansion{
		Name:                name,
		Category:            e.category,
		Unit:                e.unit_type,
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
