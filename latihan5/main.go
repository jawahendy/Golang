package main

import "fmt"

type Car struct {
	wheelCount int
}

func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Ferrari struct {
	Car
}

func (f Ferrari) sayHiToHendy() {
	fmt.Println("Hi Hendy this Your wheels!")
}

type AstonMartin struct {
	Car
}

func (a AstonMartin) sayHiToHend() {
	fmt.Println("Hi this many wheels")
}

func main() {
	f := Ferrari{Car{4}}
	fmt.Println("A Ferrari has this many wheels: ", f.numberOfWheels())
	f.sayHiToHendy()

	a := AstonMartin{Car{4}}
	fmt.Println("An Hendy has this many wheels: ", a.numberOfWheels())
	a.sayHiToHend()
}
