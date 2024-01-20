package main

func main() {

}
func init()  {
	rand.Seed(time.Now().Unix())
}

// RandomMatrix generate square matrix of size n and fills cells with random values
func RandomMatrix(n int) [][]int {
	randMatrix := make([][]int, n)
	for i := range randMatrix {
		randMatrix[i] = make([]int, n)
	}
	rand
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			randMatrix[i][j] = rand.Int()
		}
	}

	return randMatrix
}
//type UnitType string
//
//type Unit struct {
//	Value float64
//	T     UnitType
//}
//
//type Dimensions interface {
//	Length() Unit
//}
//
//type Auto interface {
//	Brand() string
//}
//type automobile struct {
//	brand    string
//	model    string
//	len      Unit
//	width    Unit
//	height   Unit
//	maxspeed int
//	engine   int
//}
//
//func (a *automobile) Brand() string {
//	return "Jeep"
//}
//func (a *automobile) Model() string {
//	return a.model
//}
//func (a *automobile) Length(num int, typ UnitType) Unit {
//	return a.len
//}
//func (a *automobile) Width(num int, typ UnitType) Unit {
//	return a.width
//}
//func (a *automobile) Height(num int, typ UnitType) Unit {
//	return a.height
//}
//func NewAuto(brand, model string) *automobile {
//	return &automobile{brand: brand, model: model}
//}
