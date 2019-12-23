package main

import "fmt"

func main() {
	var aRect = Rectangle{
		NamedObj{"Segitiga"},
		Shape{NamedObj{"Triangle"}, 0, true},
		Point{0, 0},
		20, 2.5,
	}

	fmt.Println(aRect.Name)
	fmt.Println(aRect.Shape.Name)
}

type NamedObj struct {
	Name string
}

type Shape struct {
	NamedObj
	color     int
	isRegular bool
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	NamedObj
	Shape
	center        Point
	Width, Height float64
}
