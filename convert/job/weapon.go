package convert

import (
	"github.com/capybara-alt/my-assemble/core"
	"github.com/capybara-alt/my-assemble/model"
)

type WeaponList struct {
	validation_schema model.ValidationUnitSchemaList
}

func NewWeaponList(s model.ValidationUnitSchemaList) IConvertor[model.Weapon] {
	return &WeaponList{
		validation_schema: s,
	}
}

func (w *WeaponList) Convert(results model.CrawlResultJSON) ([]model.Weapon, error) {
	weapon_list := []model.Weapon{}
	unit_type := results.GetUnitType()
	categories := results[unit_type]
	for category_name, category := range categories {
		wc := newWeapon(category_name, unit_type, w.validation_schema)
		for name, weapon_info := range category {
			weapon, err := wc.convert(name, weapon_info)
			if err != nil {
				return nil, err
			}
			weapon_list = append(weapon_list, *weapon)
		}
	}

	return weapon_list, nil
}

type weapon struct {
	category          string
	unit_type         string
	validation_schema model.ValidationUnitSchemaList
}

func newWeapon(category, unit_type string, s model.ValidationUnitSchemaList) *weapon {
	return &weapon{
		category:          category,
		unit_type:         unit_type,
		validation_schema: s,
	}
}

func (w *weapon) convert(name string, weapon_info model.UnitInfoJSON) (*model.Weapon, error) {
	converted, err := w.validation_schema.ConvertValues(weapon_info)
	if err != nil {
		return nil, err
	}

	return &model.Weapon{
		Name:                 name,
		Category:             w.category,
		Maker:                converted["maker"].(string),
		ENLoad:               converted["en_load"].(int64),
		Price:                converted["price"].(int64),
		Weight:               converted["weight"].(int64),
		Unit:                 w.unit_type,
		PAInterference:       converted["pa_interference"].(int64),
		Cooling:              converted["cooling"].(int64),
		Recoil:               converted["recoil"].(int64),
		TKHeatBuildup:        converted["tk_heat_buildup"].(int64),
		DeploymentRange:      converted["deployment_range"].(int64),
		AmmunitionCost:       converted["ammunition_cost"].(string),
		MagazineRounds:       converted["magazine_rounds"].(int64),
		IdealRange:           converted["ideal_range"].(int64),
		AttackPower:          converted["attack_power"].(string),
		ATKHeatBuildup:       converted["atk_heat_buildup"].(int64),
		DamageMitigation:     converted["damage_mitigation"].(int64),
		ImpactDampening:      converted["impact_dampening"].(int64),
		IdleDamageMitigation: converted["idle_damage_mitigation"].(int64),
		IdleTime:             converted["idle_time"].(float64),
		IdleImpactMitigation: converted["idle_impact_mitigation"].(int64),
		ReloadTime:           converted["reload_time"].(float64),
		MaxLockCount:         converted["max_lock_count"].(int64),
		EffectiveRange:       converted["effective_range"].(int64),
		BlastRadius:          converted["blast_radius"].(int64),
		DirectHitAdjustment:  converted["direct_hit_adjustment"].(int64),
		TotalRounds:          converted["total_rounds"].(int64),
		Impact:               converted["impact"].(string),
		AccumulativeImpact:   converted["accumulative_impact"].(string),
		Guidance:             converted["guidance"].(int64),
		HomingLockTime:       converted["homing_lock_time"].(float64),
		RapidFire:            converted["rapid_fire"].(float64),
		ConsecutiveHits:      converted["consecutive_hits"].(int64),
		IGDuration:           converted["ig_duration"].(float64),
		IGDamageMitigation:   converted["ig_damage_mitigation"].(int64),
		IGImpactDampening:    converted["ig_impact_dampening"].(int64),
		ChgEnLoad:            converted["chg_en_load"].(int64),
		ChgTime:              converted["chg_time"].(float64),
		ChgAttackPower:       converted["chg_attack_power"].(string),
		ChgHeatBuildup:       converted["chg_heat_buildup"].(int64),
		ChgAmmoConsumption:   converted["chg_ammo_consumption"].(int64),
		ChgBlastRadius:       converted["chg_blast_radius"].(int64),
		ChgImpact:            converted["chg_impact"].(string),
		ChgAccumImpact:       converted["chg_accum_impact"].(string),
		FullChgAttackPower:   converted["full_chg_attack_power"].(string),
		FullChgHeatBuildup:   converted["full_chg_heat_buildup"].(int64),
		FullChgAmmoConsump:   converted["full_chg_ammo_consump"].(int64),
		FullChgBlastRadius:   converted["full_chg_blast_radius"].(int64),
		FullChgImpact:        converted["full_chg_impact"].(string),
		FullChgAccumImpact:   converted["full_chg_accum_impact"].(string),
		FullChgTime:          converted["full_chg_time"].(float64),
		CreatedAt:            core.Now(),
	}, nil
}
