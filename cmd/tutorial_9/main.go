package main

import "fmt"

// Generics
type gasEngine struct {
	gallons float32
	mpg     float32
}

type eletricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | eletricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	var eletricCar = car[eletricEngine]{
		carMake:  "Tesla",
		carModel: "Civic",
		engine: eletricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}
	fmt.Printf("%v\n", gasCar)
	fmt.Printf("%v\n", eletricCar)
}
