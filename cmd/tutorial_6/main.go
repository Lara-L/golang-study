package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature string = "tubarão"
	var pointer *string = &creature

	fmt.Println("creature = ", creature)
	fmt.Println("pointer = ", pointer)

	//pra mostrar o valor e não o endereço
	fmt.Println("*pointer = ", *pointer)

	var creature2 Creature = Creature{Species: "cachorro"}

	fmt.Println("1)", creature2)
	changeCreature2(creature2)
	fmt.Println("3)", creature2)

	var creature3 Creature = Creature{Species: "borboleta"}

	fmt.Println("1)", creature3)
	changeCreature3(&creature3)
	fmt.Println("3)", creature3)
}

func changeCreature3(creature3 *Creature) {
	creature3.Species = "pantera"
	fmt.Println("2)", creature3)
}

func changeCreature2(creature2 Creature) {
	creature2.Species = "gato"
	fmt.Println("2)", creature2)
}
