package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	example01()
}

func example01() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

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
func outra(canal1 chan int, canal2 chan int) {
	var wg sync.WaitGroup
	for v := range canal1 {
		wg.Add(1)
		go func(v int) {
			canal2 <- trabalho(v)
			wg.Done()

		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return n
}
