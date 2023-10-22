package model

import (
	"time"

	myassemblyv1 "buf.build/gen/go/capybara/my-assemble/protocolbuffers/go/myassembly/v1"
)

type Frame struct {
	Name                  string `gorm:"primaryKey;index:frame_idx"`
	Category              string
	Unit                  string
	ABThrust              int64
	AbEnConsumption       int64
	AP                    int64
	ENLoad                int64
	QBReloadIdealWeight   int64
	QBReloadTime          float64
	QBJetDuration         float64
	QBThrust              int64
	QbEnConsumption       int64
	SystemRecovery        int64
	GeneratorSupplyAdj    int64
	GeneratorOutputAdj    int64
	ScanStandbyTime       float64
	ScanEffectDuration    float64
	ScanDistance          int64
	BoosterEfficiencyAdj  int64
	Maker                 string
	UpwardThrust          int64
	UpwardENConsumption   int64
	Price                 int64
	RecoilControl         int64
	JumpHeight            int64
	AttitudeStability     int64
	FirearmSpecialization int64
	Thrust                int64
	JumpDistance          int64
	LoadLimit             int64
	AntiEnergyDefense     int64
	AntiKineticDefense    int64
	AntiExplosiveDefense  int64
	ArmsLoadLimit         int64
	TravelSpeed           int64
	MeleeSpecialization   int64
	Weight                int64
	HighSpeedPref         int64
	CreatedAt             time.Time
}

func (f *Frame) ToPB() *myassemblyv1.Frame {
	return &myassemblyv1.Frame{
		Name:                  f.Name,
		Category:              f.Category,
		Unit:                  f.Unit,
		AbThrust:              int32(f.ABThrust),
		AbEnConsumption:       int32(f.AbEnConsumption),
		Ap:                    int32(f.AP),
		EnLoad:                int32(f.ENLoad),
		QbReloadIdealWeight:   int32(f.QBReloadIdealWeight),
		QbReloadTime:          float32(f.QBReloadTime),
		QbJetDuration:         float32(f.QBJetDuration),
		QbThrust:              int32(f.QBThrust),
		QbEnConsumption:       int32(f.QbEnConsumption),
		SystemRecovery:        int32(f.SystemRecovery),
		GeneratorSupplyAdj:    int32(f.GeneratorSupplyAdj),
		GeneratorOutputAdj:    int32(f.GeneratorOutputAdj),
		ScanStandbyTime:       float32(f.ScanStandbyTime),
		ScanEffectDuration:    float32(f.ScanEffectDuration),
		ScanDistance:          int32(f.ScanDistance),
		BoosterEfficiencyAdj:  int32(f.BoosterEfficiencyAdj),
		Maker:                 f.Maker,
		UpwardThrust:          int32(f.UpwardThrust),
		UpwardEnConsumption:   int32(f.UpwardENConsumption),
		Price:                 int32(f.Price),
		RecoilControl:         int32(f.RecoilControl),
		JumpHeight:            int32(f.JumpHeight),
		AttitudeStability:     int32(f.AttitudeStability),
		FirearmSpecialization: int32(f.FirearmSpecialization),
		Thrust:                int32(f.Thrust),
		JumpDistance:          int32(f.JumpDistance),
		LoadLimit:             int32(f.LoadLimit),
		AntiEnergyDefense:     int32(f.AntiEnergyDefense),
		AntiKineticDefense:    int32(f.AntiKineticDefense),
		AntiExplosiveDefense:  int32(f.AntiExplosiveDefense),
		ArmsLoadLimit:         int32(f.ArmsLoadLimit),
		TravelSpeed:           int32(f.TravelSpeed),
		MeleeSpecialization:   int32(f.MeleeSpecialization),
		Weight:                int32(f.Weight),
		HighSpeedPref:         int32(f.HighSpeedPref),
		CreatedAt:             f.CreatedAt.UnixMilli(),
	}
}

func (f *Frame) FromPB(frame *myassemblyv1.Frame) *Frame {
	return &Frame{
		Name:                  frame.Name,
		Category:              frame.Category,
		Unit:                  frame.Unit,
		ABThrust:              int64(frame.AbThrust),
		AbEnConsumption:       int64(frame.AbEnConsumption),
		AP:                    int64(frame.Ap),
		ENLoad:                int64(frame.EnLoad),
		QBReloadIdealWeight:   int64(frame.QbReloadIdealWeight),
		QBReloadTime:          float64(frame.QbReloadTime),
		QBJetDuration:         float64(frame.QbJetDuration),
		QBThrust:              int64(frame.QbThrust),
		QbEnConsumption:       int64(frame.QbEnConsumption),
		SystemRecovery:        int64(frame.SystemRecovery),
		GeneratorSupplyAdj:    int64(frame.GeneratorSupplyAdj),
		GeneratorOutputAdj:    int64(frame.GeneratorOutputAdj),
		ScanStandbyTime:       float64(frame.ScanStandbyTime),
		ScanEffectDuration:    float64(frame.ScanEffectDuration),
		ScanDistance:          int64(frame.ScanDistance),
		BoosterEfficiencyAdj:  int64(frame.BoosterEfficiencyAdj),
		Maker:                 frame.Maker,
		UpwardThrust:          int64(frame.UpwardThrust),
		UpwardENConsumption:   int64(frame.UpwardEnConsumption),
		Price:                 int64(frame.Price),
		RecoilControl:         int64(frame.RecoilControl),
		JumpHeight:            int64(frame.JumpHeight),
		AttitudeStability:     int64(frame.AttitudeStability),
		FirearmSpecialization: int64(frame.FirearmSpecialization),
		Thrust:                int64(frame.Thrust),
		JumpDistance:          int64(frame.JumpDistance),
		LoadLimit:             int64(frame.LoadLimit),
		AntiEnergyDefense:     int64(frame.AntiEnergyDefense),
		AntiKineticDefense:    int64(frame.AntiKineticDefense),
		AntiExplosiveDefense:  int64(frame.AntiExplosiveDefense),
		ArmsLoadLimit:         int64(frame.ArmsLoadLimit),
		TravelSpeed:           int64(frame.TravelSpeed),
		MeleeSpecialization:   int64(frame.MeleeSpecialization),
		Weight:                int64(frame.Weight),
		HighSpeedPref:         int64(frame.HighSpeedPref),
		CreatedAt:             time.UnixMilli(frame.CreatedAt),
	}
}
