package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines

var contador int
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	contador++
	fmt.Println("Contador: ", contador)
	time.Sleep(500 * time.Millisecond)
	mutex.Unlock()
}

//func imprimirNumeros() {
//	for i := 1; i < 10; i++ {
//		fmt.Println(i)
//		time.Sleep(500 * time.Millisecond)
//	}
//}

func main() {
	//go imprimirNumeros()
	//fmt.Println("Goroutine iniciada!")
	//
	//time.Sleep(3 * time.Second)
	//fmt.Println("Fim do Programa")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}
	wg.Wait()
	fmt.Println("Valor final do contador: ", contador)
}
