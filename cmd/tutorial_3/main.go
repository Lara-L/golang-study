package main

import "fmt"

// ARRAYS,SLICES, MAPS E LOOPS

func main() {
	var intArr [3]int32
	intArr[0] = 1
	fmt.Println(intArr[1:3])

	exArr := [3]int32{1, 2, 3}
	fmt.Println(exArr)

	intSlice := []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{4, 5, 6}
	intSlice2 = append(intSlice2, 7, 8, 9)
	fmt.Printf("O tamanho é %v e a capacidade é %v\n", len(intSlice2), cap(intSlice2))
	intSlice2 = append(intSlice2, 9)
	fmt.Printf("O tamanho é %v e a capacidade é %v\n", len(intSlice2), cap(intSlice2))

	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Printf("%v", intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Sarah"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("A idade é %v", age)
	} else {
		fmt.Println("Nome inválido")
	}
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Value: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
