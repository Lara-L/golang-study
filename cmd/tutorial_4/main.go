package main

import (
	"fmt"
	"strings"
)

// STRINGS, RUNES E BYTES

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\nO tamanho de 'myString' é %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRUne = %v", myRune)

	//ineficiente
	var strSlice = []string{"t", "e", "s", "t"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	var strSlice2 = []string{"s", "e", "c", "o", "n", "d", " ", "t", "e", "s", "t"}
	var strBuilder strings.Builder
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i])
	}
	var carStr2 = strBuilder.String()
	fmt.Printf("\n%v", carStr2)
}
