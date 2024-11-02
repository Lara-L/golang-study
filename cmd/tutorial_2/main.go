package main

// FUNÇÕES E ESTRUTURAS DE CONTROLE

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Olá Mundo!"
	printMe(printValue)

	var numerador int = 13
	var denominador int = 2
	result, remainder, err := intDivision(numerador, denominador)
	if err != nil {
		fmt.Println(err)
	} else if denominador == 0 {
		fmt.Printf("O resultado da divisão inteira foi %v", result)
	} else {
		fmt.Printf("O resultado da divisão inteira é %v com restante %v", result, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerador int, denominador int) (int, int, error) {
	var err error
	if denominador == 0 {
		err = errors.New("Não pode dividir por zero")
		return 0, 0, err
	}

	var result int = numerador / denominador
	var remainder int = numerador % denominador
	return result, remainder, err
}
