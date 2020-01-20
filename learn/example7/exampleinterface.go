package main

import (
	"fmt"
	"math"
)

type bangunan interface {
	keliling() float64
	luas() float64
}

type cube struct {
	panjang, lebar float64
}

type circle struct {
	radius float64
}

func hitungbangunan(b bangunan) {
	fmt.Println("kelilingnya ", b.keliling())
	fmt.Println("Luasnya ", b.luas())
}

func (c cube) keliling() float64 {
	return 2*c.panjang + 2*c.lebar
}

func (c cube) luas() float64 {
	return c.panjang * c.lebar
}

func (l circle) keliling() float64 {
	return 2 * math.Pi * l.radius
}

func (l circle) luas() float64 {
	return math.Pi * l.radius * l.radius
}

func main() {

	cube1 := cube{5, 10}
	hitungbangunan(cube1)

	lingkaran := circle{10}
	hitungbangunan(lingkaran)
}
