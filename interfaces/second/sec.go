package sec

import "fmt"

type Greeterr interface {
	Greett(string) string
}
type Russiann struct{}

func (r Russiann) Greett(namee, sur string) string {
	return fmt.Sprintf("Привет, %s%s", namee, sur)
}
