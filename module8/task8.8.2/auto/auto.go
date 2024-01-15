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
		len: Unit{Value: lengh, T: CM},
		wid: Unit{Value: width, T: CM},
		hei: Unit{Value: height, T: CM},
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
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}
type auto struct {
	brand string
	mod   string
	dim   Dimensions
	max   int
	eng   int
}

func (m auto) Brand() string {
	return m.brand
}
func (m auto) Model() string {
	return m.mod
}
func (m auto) Dimensions() Dimensions {

	return m.dim
}
func (m auto) MaxSpeed() int {
	return m.max
}
func (m auto) EnginePower() int {
	return m.eng
}
func NewAuto(brand string, mod string, dim Dimensions, max int, eng int) *auto {
	return &auto{
		brand: brand,
		mod:   mod,
		dim:   dim,
		max:   max,
		eng:   eng,
	}
}
