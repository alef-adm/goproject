package main

func main() {
	//str := "⌘"
	//
	//fmt.Println("представление байт в разных системах счисления", len(str))
	//fmt.Println(fmt.Sprintf("десятичное: %d", str))
	//fmt.Println(fmt.Sprintf("двоичноеЖ %b", str))
	//fmt.Println(fmt.Sprintf("шестнадцатиричное: %x", str))
	//
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%x", str[i])
	//}
	///////////////////////////////////////////////////////////////////
	//str := "Игорь"
	//fmt.Println(str[2:6])
	/////////////////////////////////////////////////////////
	//str := "Игорь"
	//for i, item := range str {
	//	fmt.Println(i, fmt.Sprintf("%#U", item))
	//}

}

type UnitType string

type Unit struct {
	Value float64
	T     UnitType
}

type Dimensions interface {
	Length() Unit
}

type Auto interface {
	Brand() string
}
type automobile struct {
	brand    string
	model    string
	len      Unit
	width    Unit
	height   Unit
	maxspeed int
	engine   int
}

func (a *automobile) Brand() string {
	return "Jeep"
}
func (a *automobile) Model() string {
	return a.model
}
func (a *automobile) Length(num int, typ UnitType) Unit {
	return a.len
}
func (a *automobile) Width(num int, typ UnitType) Unit {
	return a.width
}
func (a *automobile) Height(num int, typ UnitType) Unit {
	return a.height
}
func NewAuto(brand, model string) *automobile {
	return &automobile{brand: brand, model: model}
}
