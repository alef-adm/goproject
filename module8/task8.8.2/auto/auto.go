package auto

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type dodge struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (m *dodge) Brand() string {
	return "Dodge"
}
func (m *dodge) Model() string {
	return m.model
}
func (m *dodge) Dimensions() Dimensions {
	return m.dimensions
}
func (m *dodge) MaxSpeed() int {
	return m.maxSpeed
}
func (m *dodge) EnginePower() int {
	return m.enginePower
}

//func NewAutoDodge(m string, d Dimensions, max int, pow int) *dodge {
//	return &dodge{
//		model:       m,
//		dimensions:  d,
//		maxSpeed:    max,
//		enginePower: pow,
//	}
//}

//type bmw struct {
//	model       string
//	dimensions  Dimensions
//	maxSpeed    int
//	enginePower int
//}
//
//func (b *bmw) Brand() string {
//	return "BMW"
//}
//func (b *bmw) Model() string {
//	return b.model
//}
//func (b *bmw) Dimensions() Dimensions {
//	return NewDimensionCM(
//		b.dimensions.Length().Get(CM),
//		//b.dimensions.Width().Get(CM),
//		//b.dimensions.Height().Get(CM),
//	)
//}
//func (b *bmw) MaxSpeed() int {
//	return b.maxSpeed
//}
//func (b *bmw) EnginePower() int {
//	return b.enginePower
//}
//func NewAutoBMW(model string, d Dimensions, maxSpeed int, enginePower int) *bmw {
//	return &bmw{
//		model:       model,
//		dimensions:  d,
//		maxSpeed:    maxSpeed,
//		enginePower: enginePower,
//	}
//}
