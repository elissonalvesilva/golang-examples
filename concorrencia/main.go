package main

import (
	"fmt"
	"runtime"
	"sync"
)

// funções acessando a mesma variavel solução possivel com mutex
var wg sync.WaitGroup

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	contador := 0
	totalGoRoutine := 1000

	wg.Add(totalGoRoutine)

	for i := 0; i < totalGoRoutine; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	fmt.Println("Valor final: ", contador)
}
