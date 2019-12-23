package main

import (
	"fmt"
	"math"
)

type hitung interface {
	Luas() float64
	keliling() float64
}

type circle struct {
	diameter float64
}

func (c circle) jariJari() float64 {
	return c.diameter / 2
}

func (c circle) Luas() float64 {
	return math.Pi * math.Pow(c.jariJari(), 2)
}

func (c circle) keliling() float64 {
	return math.Pi * c.diameter
}

type square struct {
	sisi float64
}

func (s square) Luas() float64 {
	return math.Pow(s.sisi, 2)
}

func (s square) keliling() float64 {
	return s.sisi * 4
}

func main() {
	var luasnya hitung

	luasnya = square{12.0}
	fmt.Println("luasnya  :", luasnya.Luas())
	fmt.Println("kelilingnya :", luasnya.keliling())

	luasnya = circle{28.0}
	fmt.Println("luasnya   :", luasnya.Luas())
	fmt.Println("Keliling  :", luasnya.keliling())
	fmt.Println("jari Jari  :", luasnya.(circle).jariJari())
}
