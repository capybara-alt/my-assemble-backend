package model

import (
	"time"

	"gorm.io/gorm"
)

type Assembly struct {
	gorm.Model
	ID                uint `gorm:"primaryKey;autoImcrement;index:assembly_idx"`
	Name              string
	UserUid           string
	LeftArmUnitName   *string
	RightArmUnitName  *string
	LeftBackUnitName  *string
	RightBackUnitName *string
	HeadName          string
	CoreName          string
	LegsName          string
	ArmsName          string
	BoosterName       string
	FCSName           string
	GeneratorName     string
	ExpansionName     *string
	User              User       `gorm:"foreignKey:UserUid"`
	LeftArmUnit       *Weapon    `gorm:"foreignKey:LeftArmUnitName"`
	RightArmUnit      *Weapon    `gorm:"foreignKey:RightArmUnitName"`
	LeftBackUnit      *Weapon    `gorm:"foreignKey:LeftBackUnitName"`
	RightBackUnit     *Weapon    `gorm:"foreignKey:RightBackUnitName"`
	Head              Frame      `gorm:"foreignKey:HeadName"`
	Core              Frame      `gorm:"foreignKey:CoreName"`
	Legs              Frame      `gorm:"foreignKey:LegsName"`
	Arms              Frame      `gorm:"foreignKey:ArmsName"`
	Booster           InnerUnit  `gorm:"foreignKey:BoosterName"`
	FCS               InnerUnit  `gorm:"foreignKey:FCSName"`
	Generator         InnerUnit  `gorm:"foreignKey:GeneratorName"`
	Expansion         *Expansion `gorm:"foreignKey:ExpansionName"`
	CreatedAt         time.Time
}
