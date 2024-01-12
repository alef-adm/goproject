package main

import (
	"fmt"
	"interfaces/second"
)

type greeter interface {
	greet(string) string
	hi(string) string
}

type russian struct{}
type american struct{}

func (r *russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}
func SayHello(g greeter, name string) {
	fmt.Println(g.greet(name))

}
func SayyHello(t sec.Greeterr, namee, sur string) {
	fmt.Println(t.Greett(namee))
}
func main() {
	SayyHello(&sec.Russiann{}, "Петяя", "Иванов")
}
