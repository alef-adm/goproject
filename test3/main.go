package main

import (
	"fmt"
	"strings"
)

type Formatter interface {
	Format(string) string
}

type SuperString struct {
	base     string
	formater Formatter
}

func (s SuperString) Print() {
	fmt.Println(s.formater.Format(s.base))
}

type F1 struct{}

func (f F1) Format(s string) string {
	return fmt.Sprintf("Формат 1 %v", strings.ToUpper(s))
}

type F2 struct{}

func (f F2) Format(s string) string {
	return fmt.Sprintf("Формат 2 %v", strings.ToLower(s))
}

func main() {
	s := SuperString{
		base:     "cТрОкa",
		formater: F1{},
	}
	s.Print()

	s.formater = F2{}

	s.Print()
}
