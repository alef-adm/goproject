package main

import "fmt"

type struct1 struct {
	param1 string
	param2 string
}
type interface1 interface {
	Metod1() struct1
}
type interface2 interface {
	Metod2() interface1
}

type rez struct {
	b struct1
	c interface1
}

func (r *rez) Metod1() struct1 {
	return r.b
}
func (r *rez) Metod2() interface1 {
	return r.c
}

type rez2 struct {
	b struct1
}

func (r *rez2) Metod1() struct1 {
	return r.b
}

func main() {
	s2 := struct1{
		param1: "three",
		param2: "four",
	}

	r2 := rez2{
		b: s2,
	}

	r1 := rez{
		b: struct1{
			param1: "one",
			param2: "two",
		},
		c: &r2,
	}

	fmt.Println(r1.b.param1)
	fmt.Println(r1.b.param2)
	fmt.Println(r1.c.Metod1().param1)
	fmt.Println(r1.c.Metod1().param2)

	fmt.Println("====================")
	var i2 interface2
	i2 = &r1
	fmt.Println(i2.Metod2().Metod1().param1)
	fmt.Println(i2.Metod2().Metod1().param2)
}
