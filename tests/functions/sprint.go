package main

import "fmt"

type CoffeeMachine interface {
	Type() string
	Brand() string
	Model() string
}
type baseCoffeeMachine struct {
	brand string
	model string
}

func (m *baseCoffeeMachine) Brand() string {
	return m.brand
}
func (m *baseCoffeeMachine) Model() string {
	return m.model
}

// капельная кофемашина
type DripCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *DripCoffeeMachine) Type() string {

	return "drip"
}

// капсульная кофемашина
type CapsuleCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *CapsuleCoffeeMachine) Type() string {
	m.brand = "apple"
	return "capsule"
}

// рожковая кофемашина
type CarobCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *CarobCoffeeMachine) Type() string { return "carob" }

func main() {
	var coffeeMachine CoffeeMachine = &CapsuleCoffeeMachine{}

	fmt.Println(coffeeMachine.Type())
}
