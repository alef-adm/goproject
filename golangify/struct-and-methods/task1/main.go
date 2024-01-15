package main

func main() {
	type location struct {
		lat, long float64
	}

}

type location struct {
	lat, long float64
}
type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
// Кратер Брэдбери: 4°35'22.2" S, 137°26'30.1" E
var lat, long coordinate
lat{4, 35, 22.2, 'S'}
long = coordinate{137, 26, 30.12, 'E'}
var qq coordinate

fmt.Println(lat.decimal(), long.decimal()) // Выводит: -4.5895 137.4417