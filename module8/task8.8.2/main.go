package main

import (
	"fmt"
	"task8.8.2/auto"
)

func main() {
	charger := auto.NewAuto("Dodge", "Charger", auto.NewDimensionCM(550, 150, 148), 200, 300)
	x5 := auto.NewAuto("BMW", "X5", auto.NewDimensionInch(492, 174, 221), 231, 230)
	e200 := auto.NewAuto("Mercedes", "E200", auto.NewDimensionCM(492, 185, 146), 233, 184)
	fmt.Println(charger)
	fmt.Println(x5)
	fmt.Println(e200)
}
