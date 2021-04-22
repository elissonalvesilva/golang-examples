package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64
	contador = 0
	totalGoRoutine := 1000

	var wg sync.WaitGroup
	wg.Add(totalGoRoutine)

	for i := 0; i < totalGoRoutine; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("CONTADOR \t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final: ", contador)
}
