package model

type CategoryGroupedJSON map[string]NameGroupedJSON
type NameGroupedJSON map[string]UnitInfoJSON
type UnitInfoJSON map[string]string
type UnitInfoValueTypes string
type CrawlResultJSON map[string]CategoryGroupedJSON

func (c CrawlResultJSON) GetUnitType() string {
	for key := range c {
		return key
	}

	return ""
}

const (
	INT    UnitInfoValueTypes = "int"
	FLOAT  UnitInfoValueTypes = "float"
	STRING UnitInfoValueTypes = "string"
)

type SecondaryUnitType string

const (
	CROSS_WEAPON_UNIT_TYPE        SecondaryUnitType = "LEFT_ARM_UNIT"
	DEFAULT_WEAPON_ARM_UNIT_TYPE  SecondaryUnitType = "ARM_UNIT"
	SHIELD_WEAPON_UNIT_TYPE       SecondaryUnitType = "LEFT_BACK_UNIT"
	DEFAULT_WEAPON_BACK_UNIT_TYPE SecondaryUnitType = "BACK_UNIT"
	HEAD_UNIT_TYPE                SecondaryUnitType = "HEAD_UNIT"
	CORE_UNIT_TYPE                SecondaryUnitType = "CORE_UNIT"
	ARMS_UNIT_TYPE                SecondaryUnitType = "ARMS_UNIT"
	LEGS_UNIT_TYPE                SecondaryUnitType = "LEGS_UNIT"
	BOOSTER_INNER_UNIT_TYPE       SecondaryUnitType = "BOOSTER_UNIT"
	FCS_INNER_UNIT_TYPE           SecondaryUnitType = "FCS_UNIT"
	GENERATOR_INNER_UNIT_TYPE     SecondaryUnitType = "GENERATOR_UNIT"
	EXPANSION_TYPE                SecondaryUnitType = "EXPANSION"
)

type PrimaryUnitType string

const (
	WEAPON      PrimaryUnitType = "WEAPON"
	FRAME       PrimaryUnitType = "FRAME"
	INNER_UNITS PrimaryUnitType = "INNER_UNITS"
	EXPANSION   PrimaryUnitType = "EXPANSION"
)
