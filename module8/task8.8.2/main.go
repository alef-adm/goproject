package main

import (
	"fmt"
	"task8.8.2/auto"
)

func main() {
	//	newInch := auto.NewDimensionInch(550, 165, 148)
	charger := auto.NewDodgeAuto("Charger", auto.NewDimensionCM(550, 150, 148), 200, 300)
	fmt.Println(charger)
}
