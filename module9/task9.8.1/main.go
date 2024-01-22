package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Vasya": {"Vasya", "Vasiliev", 30, "male", 0},
		"Petya": {"Petya", "Petrov", 31, "male", 1},
		"Vitya": {"Vitya", "Vitin", 32, "male", 2},
		"Serg":  {"Serg", "Sergeev", 33, "male", 3},
	}
	suspects := []string{"Vasya", "Petya", "Moriarti"}

	var crim Man // most crime
	var k bool
	for _, name := range suspects {
		p, ok := people[name]
		if !ok {
			continue
		}
		if p.Crimes > crim.Crimes {
			crim = p
			k = true
		}
	}
	if k {
		fmt.Println("Самый опасный:", crim.LastName)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}

//fmt.Println(k, c)
