package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
)

type ValidationUnitSchemaList []ValidationUnitSchema

func NewValidationUnitSchemaList() ValidationUnitSchemaList {
	return []ValidationUnitSchema{}
}

func (v ValidationUnitSchemaList) ConvertValues(info UnitInfoJSON) (map[string]interface{}, error) {
	converted := make(map[string]interface{})
	errs := []error{}
	for _, schema := range v {
		value, ok := info[schema.NameJa]
		if !ok {
			value = ""
		}
		v, err := schema.convertValue(value)
		if err != nil {
			msg := fmt.Sprintf("(%s>%s) Invalid value type expect: %s, value: %s", schema.UnitType, schema.PropName, schema.ValueType, value)
			errs = append(errs, errors.New(msg))
		} else {
			converted[schema.PropName] = v
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	} else {
		return converted, nil
	}
}

func (v ValidationUnitSchemaList) ToPB() []*myassemblyv1.ValidationUnitSchema {
	vus_list := make([]*myassemblyv1.ValidationUnitSchema, len(v))
	for i, vus := range v {
		vus_list[i] = vus.ToPB()
	}

	return vus_list
}

type ValidationUnitSchema struct {
	PropName  string `gorm:"primaryKey;index:validation_unit_schema_idx"`
	NameJa    string
	NameEn    string
	ValueType UnitInfoValueTypes
	UnitType  PrimaryUnitType `gorm:"primaryKey"`
}

func NewValidationUnitSchema() *ValidationUnitSchema {
	return &ValidationUnitSchema{}
}

func (v ValidationUnitSchema) ToPB() *myassemblyv1.ValidationUnitSchema {
	return &myassemblyv1.ValidationUnitSchema{
		PropName:  v.PropName,
		NameJa:    v.NameJa,
		NameEn:    v.NameEn,
		ValueType: string(v.ValueType),
		UnitType:  string(v.UnitType),
	}
}

func (v ValidationUnitSchema) FromPB(validation_unit_schema ValidationUnitSchema) *ValidationUnitSchema {
	return &ValidationUnitSchema{
		PropName:  validation_unit_schema.PropName,
		NameJa:    validation_unit_schema.NameJa,
		NameEn:    validation_unit_schema.NameEn,
		ValueType: validation_unit_schema.ValueType,
		UnitType:  validation_unit_schema.UnitType,
	}
}

func (v ValidationUnitSchema) convertValue(value string) (interface{}, error) {
	switch v.ValueType {
	case "INT":
		return v.convertToInt(value)
	case "FLOAT":
		return v.convertToFloat(value)
	}

	return v.convertToStr(value)
}

func (v ValidationUnitSchema) convertToStr(value string) (string, error) {
	value = strings.Trim(value, "\n")
	value = strings.Split(value, " (")[0]

	return value, nil
}

func (v ValidationUnitSchema) convertToInt(value string) (int64, error) {
	var ivalue int64 = 0
	if value == "" || value == "-" || value == "." {
		return ivalue, nil
	}

	value = strings.Trim(value, "\n")
	value = strings.ReplaceAll(value, ",", "")
	value = strings.Split(value, "(")[0]
	value = strings.Split(value, " ")[0]

	return strconv.ParseInt(value, 10, 64)
}

func (v ValidationUnitSchema) convertToFloat(value string) (float64, error) {
	var ivalue float64 = 0.0
	if value == "" || value == "-" || value == "." {
		return ivalue, nil
	}

	value = strings.Trim(value, "\n")
	value = strings.ReplaceAll(value, ",", "")
	value = strings.Split(value, "(")[0]
	value = strings.Split(value, " ")[0]

	return strconv.ParseFloat(value, 64)
}
