package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()
}

func func1() {

	for i := 0; i < 10; i++ {
		fmt.Println("Func1() - ", i)
		time.Sleep(200)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2() - ", i)
		time.Sleep(200)

	}
	wg.Done()
}
