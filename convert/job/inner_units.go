package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type InnerUnitsList struct {
	validationSchema model.ValidationUnitSchemaList
}

func NewInnerUnitsList(s model.ValidationUnitSchemaList) IConvertor[model.InnerUnit] {
	return &InnerUnitsList{
		validationSchema: s,
	}
}

func (f *InnerUnitsList) Convert(results model.CrawlResultJSON) ([]model.InnerUnit, error) {
	innerUnitList := []model.InnerUnit{}
	unitType := results.GetUnitType()
	categories := results[unitType]
	for cname, category := range categories {
		wc := newInnerUnits(cname, unitType, f.validationSchema)
		for name, info := range category {
			innerUnit, err := wc.convert(name, info)
			if err != nil {
				return nil, err
			}
			innerUnitList = append(innerUnitList, *innerUnit)
		}
	}

	return innerUnitList, nil
}

type innerUnits struct {
	validationSchema model.ValidationUnitSchemaList
	unitType         string
	category         string
}

func newInnerUnits(category, unitType string, s model.ValidationUnitSchemaList) *innerUnits {
	return &innerUnits{
		category:         category,
		unitType:         unitType,
		validationSchema: s,
	}
}

func (i *innerUnits) convert(name string, info model.UnitInfoJSON) (*model.InnerUnit, error) {
	converted, err := i.validationSchema.ConvertValues(info)
	if err != nil {
		return nil, err
	}

	return &model.InnerUnit{
		Name:                  name,
		Category:              i.category,
		Unit:                  i.unitType,
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
