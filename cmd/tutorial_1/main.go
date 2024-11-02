package main

// VARIÁVEIS E TIPOS BÁSICOS DE DADOS

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 20500
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 1988889.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	var floatNum1 float32 = float32(intNum1) / float32(intNum2)
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)
	fmt.Println(floatNum1)

	var myString string = "Olá" + " " + "Mundo!"
	fmt.Println(myString)

	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune) //entender melhor o tipo rune

	var myBoolean bool //por padrão: false
	fmt.Println(myBoolean)

	var intNum3 int // por padrão: 0
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)

	var3, var4 := 3, 4
	fmt.Println(var3, var4)

	const myConst string = "valor constante, imutável"
	fmt.Println(myConst)

}
