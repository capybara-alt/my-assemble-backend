package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type InnerUnitsList struct {
	validation_schema model.ValidationUnitSchemaList
}

func NewInnerUnitsList(s model.ValidationUnitSchemaList) IConvertor[model.InnerUnit] {
	return &InnerUnitsList{
		validation_schema: s,
	}
}

func (f *InnerUnitsList) Convert(results model.CrawlResultJSON) ([]model.InnerUnit, error) {
	inner_units_list := []model.InnerUnit{}
	unit_type := results.GetUnitType()
	categories := results[unit_type]
	for category_name, category := range categories {
		wc := newInnerUnits(category_name, unit_type, f.validation_schema)
		for name, inner_unit_info := range category {
			inner_units, err := wc.convert(name, inner_unit_info)
			if err != nil {
				return nil, err
			}
			inner_units_list = append(inner_units_list, *inner_units)
		}
	}

	return inner_units_list, nil
}

type innerUnits struct {
	validation_schema model.ValidationUnitSchemaList
	unit_type         string
	category          string
}

func newInnerUnits(category, unit_type string, s model.ValidationUnitSchemaList) *innerUnits {
	return &innerUnits{
		category:          category,
		unit_type:         unit_type,
		validation_schema: s,
	}
}

func (i *innerUnits) convert(name string, inner_units_info model.UnitInfoJSON) (*model.InnerUnit, error) {
	converted, err := i.validation_schema.ConvertValues(inner_units_info)
	if err != nil {
		return nil, err
	}

	return &model.InnerUnit{
		Name:                  name,
		Category:              i.category,
		Unit:                  i.unit_type,
		Maker:                 converted["maker"].(string),
		ABThrust:              converted["ab_thrust"].(int64),
		AbEnConsumption:       converted["ab_en_consumption"].(int64),
		ENOutput:              converted["en_output"].(int64),
		ENCapacity:            converted["en_capacity"].(int64),
		EnergyFirearmSpec:     converted["energy_firearm_spec"].(int64),
		ENRecharge:            converted["en_recharge"].(int64),
		ENLoad:                converted["en_load"].(int64),
		QBReloadIdealWeight:   converted["qb_reload_ideal_weight"].(int64),
		QBJetDuration:         converted["qb_jet_duration"].(float64),
		QBThrust:              converted["qb_thrust"].(int64),
		QBReloadTime:          converted["qb_reload_time"].(float64),
		QbEnConsumption:       converted["qb_en_consumption"].(int64),
		UpwardThrust:          converted["upward_thrust"].(int64),
		UpwardENConsumption:   converted["upward_en_consumption"].(int64),
		MediumRangeAssist:     converted["medium_range_assist"].(int64),
		SupplyRecovery:        converted["supply_recovery"].(int64),
		Price:                 converted["price"].(int64),
		PostRecoveryENSupply:  converted["post_recovery_en_supply"].(int64),
		Thrust:                converted["thrust"].(int64),
		MissileLockCorrection: converted["missile_lock_correction"].(int64),
		MultiLockCorrection:   converted["multi_lock_correction"].(int64),
		MeleeAttackThrust:     converted["melee_attack_thrust"].(int64),
		MeleeAtkENConsump:     converted["melee_atk_en_consump"].(int64),
		CloseRangeAssist:      converted["close_range_assist"].(int64),
		LogRangeAssist:        converted["log_range_assist"].(int64),
		Weight:                converted["weight"].(int64),
		CreatedAt:             core.Now(),
	}, nil
}
