package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funcoes := 5

	go manda(100, canal1)
	go outra(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

// convergencia
func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

// divergencia
func outra(funcoes int, canal1 chan int, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n
}
