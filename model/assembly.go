package model

import (
	"time"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
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

func (a *Assembly) ToPB() *myassemblyv1.Assembly {
	return &myassemblyv1.Assembly{
		Id:            int64(a.ID),
		Name:          a.Name,
		User:          a.User.ToPB(),
		LeftArmUnit:   a.LeftArmUnit.ToPB(),
		RightArmUnit:  a.RightArmUnit.ToPB(),
		LeftBackUnit:  a.LeftBackUnit.ToPB(),
		RightBackUnit: a.RightBackUnit.ToPB(),
		Head:          a.Head.ToPB(),
		Core:          a.Core.ToPB(),
		Legs:          a.Legs.ToPB(),
		Arms:          a.Arms.ToPB(),
		Booster:       a.Booster.ToPB(),
		Fcs:           a.FCS.ToPB(),
		Generator:     a.Generator.ToPB(),
		Expansion:     a.Expansion.ToPB(),
		CreatedAt:     a.CreatedAt.UnixMilli(),
	}
}

func (a *Assembly) FromPB(assembly *myassemblyv1.Assembly) *Assembly {
	user := &User{}
	leftArmUnit := &Weapon{}
	rightArmUnit := &Weapon{}
	leftBackUnit := &Weapon{}
	rightBackUnit := &Weapon{}
	head := &Frame{}
	core := &Frame{}
	legs := &Frame{}
	arms := &Frame{}
	booster := &InnerUnit{}
	fcs := &InnerUnit{}
	generator := &InnerUnit{}
	expansion := &Expansion{}

	return &Assembly{
		ID:            uint(assembly.Id),
		Name:          assembly.Name,
		User:          *user.FromPB(assembly.User),
		LeftArmUnit:   leftArmUnit.FromPB(assembly.LeftArmUnit),
		RightArmUnit:  rightArmUnit.FromPB(assembly.RightArmUnit),
		LeftBackUnit:  leftBackUnit.FromPB(assembly.LeftBackUnit),
		RightBackUnit: rightBackUnit.FromPB(assembly.RightBackUnit),
		Head:          *head.FromPB(assembly.Head),
		Core:          *core.FromPB(assembly.Core),
		Legs:          *legs.FromPB(assembly.Legs),
		Arms:          *arms.FromPB(assembly.Arms),
		Booster:       *booster.FromPB(assembly.Booster),
		FCS:           *fcs.FromPB(assembly.Fcs),
		Generator:     *generator.FromPB(assembly.Generator),
		Expansion:     expansion.FromPB(assembly.Expansion),
		CreatedAt:     time.UnixMilli(assembly.CreatedAt),
	}
}
