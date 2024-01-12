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
type dimensionCM struct {
	len Unit
	wid Unit
	hei Unit
}

func (d dimensionCM) Length() Unit {
	return d.len
}
func (d dimensionCM) Width() Unit {
	return d.wid
}
func (d dimensionCM) Height() Unit {
	return d.hei
}
func NewDimensionCM(lengh, width, height float64) *dimensionCM {
	return &dimensionCM{
		Unit{lengh, CM},
		Unit{width, CM},
		Unit{height, CM},
	}
}

type dimensionInch struct {
	len Unit
	wid Unit
	hei Unit
}

func (d dimensionInch) Length() Unit {
	return d.len
}
func (d dimensionInch) Width() Unit {
	return d.wid
}
func (d dimensionInch) Height() Unit {
	return d.hei
}
func NewDimensionInch(len, wid, hei float64) *dimensionInch {
	return &dimensionInch{
		len: Unit{Value: len, T: Inch},
		wid: Unit{Value: wid, T: Inch},
		hei: Unit{Value: hei, T: Inch},
	}
}

type Auto interface {
	//	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type dodgeAuto struct {
	mod string
	dim Dimensions
	max int
	eng int
}

func (m dodgeAuto) Model() string {
	return m.mod
}
func (m dodgeAuto) Dimensions() Dimensions {

	return m.dim
}
func (m dodgeAuto) MaxSpeed() int {
	return m.max
}
func (m dodgeAuto) EnginePower() int {
	return m.eng
}
func NewDodgeAuto(mod string, dim Dimensions, max int, eng int) *dodgeAuto {
	return &dodgeAuto{
		mod: mod,
		dim: dim,
		max: max,
		eng: eng,
	}
}
