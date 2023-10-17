package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type FrameList struct {
	validationSchema model.ValidationUnitSchemaList
}

func NewFrameList(s model.ValidationUnitSchemaList) IConvertor[model.Frame] {
	return &FrameList{
		validationSchema: s,
	}
}

func (f *FrameList) Convert(results model.CrawlResultJSON) ([]model.Frame, error) {
	frameList := []model.Frame{}
	unitType := results.GetUnitType()
	categories := results[unitType]
	for cname, category := range categories {
		wc := newFrame(cname, unitType, f.validationSchema)
		for name, info := range category {
			frame, err := wc.convert(name, info)
			if err != nil {
				return nil, err
			}
			frameList = append(frameList, *frame)
		}
	}

	return frameList, nil
}

type frame struct {
	validationSchema model.ValidationUnitSchemaList
	unitType         string
	category         string
}

func newFrame(category, unitType string, s model.ValidationUnitSchemaList) *frame {
	return &frame{
		category:         category,
		unitType:         unitType,
		validationSchema: s,
	}
}

func (f *frame) convert(name string, info model.UnitInfoJSON) (*model.Frame, error) {
	converted, err := f.validationSchema.ConvertValues(info)
	if err != nil {
		return nil, err
	}

	return &model.Frame{
		Name:                  name,
		Unit:                  f.unitType,
		Category:              f.category,
		ABThrust:              converted["ab_thrust"].(int64),
		AbEnConsumption:       converted["ab_en_consumption"].(int64),
		AP:                    converted["ap"].(int64),
		ENLoad:                converted["en_load"].(int64),
		QBReloadIdealWeight:   converted["qb_reload_ideal_weight"].(int64),
		QBReloadTime:          converted["qb_reload_time"].(float64),
		QBJetDuration:         converted["qb_jet_duration"].(float64),
		QBThrust:              converted["qb_thrust"].(int64),
		QbEnConsumption:       converted["qb_en_consumption"].(int64),
		SystemRecovery:        converted["system_recovery"].(int64),
		GeneratorSupplyAdj:    converted["generator_supply_adj"].(int64),
		GeneratorOutputAdj:    converted["generator_output_adj"].(int64),
		ScanStandbyTime:       converted["scan_standby_time"].(float64),
		ScanEffectDuration:    converted["scan_effect_duration"].(float64),
		ScanDistance:          converted["scan_distance"].(int64),
		BoosterEfficiencyAdj:  converted["booster_efficiency_adj"].(int64),
		Maker:                 converted["maker"].(string),
		UpwardThrust:          converted["upward_thrust"].(int64),
		UpwardENConsumption:   converted["upward_en_consumption"].(int64),
		Price:                 converted["price"].(int64),
		RecoilControl:         converted["recoil_control"].(int64),
		JumpHeight:            converted["jump_height"].(int64),
		AttitudeStability:     converted["attitude_stability"].(int64),
		FirearmSpecialization: converted["firearm_specialization"].(int64),
		Thrust:                converted["thrust"].(int64),
		JumpDistance:          converted["jump_distance"].(int64),
		LoadLimit:             converted["load_limit"].(int64),
		AntiEnergyDefense:     converted["anti_energy_defense"].(int64),
		AntiKineticDefense:    converted["anti_kinetic_defense"].(int64),
		AntiExplosiveDefense:  converted["anti_explosive_defense"].(int64),
		ArmsLoadLimit:         converted["arms_load_limit"].(int64),
		TravelSpeed:           converted["travel_speed"].(int64),
		MeleeSpecialization:   converted["melee_specialization"].(int64),
		Weight:                converted["weight"].(int64),
		HighSpeedPref:         converted["high_speed_perf"].(int64),
		CreatedAt:             core.Now(),
	}, nil
}
