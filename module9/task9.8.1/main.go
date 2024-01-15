package main

import (
	"fmt"
)

func main() {

	type Man struct {
		Name     string
		LastName string
		Age      int
		Gender   string
		Crimes   int
	}
	people := map[string]Man{
		"Vasya": {"Vasya", "Vasiliev", 30, "male", 0},
		"Petya": {"Petya", "Petrov", 31, "male", 1},
		"Vitya": {"Vitya", "Vitin", 32, "male", 2},
		"Serg":  {"Serg", "Sergeev", 33, "male", 3},
	}
	for key, value := range people {
		c := 0

		if c < value.Crimes {
			c = value.Crimes
			k := key
			fmt.Println(k)

		}
		//	fmt.Println(key, value)
	}

}
