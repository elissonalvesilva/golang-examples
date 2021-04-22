package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// exampleTodd()
	examplePiker()
}

func exampleTodd() {
	par := make(chan int)
	inpar := make(chan int)
	converge := make(chan int)

	go envia(par, inpar)
	go receive(par, inpar, converge)

	for v := range converge {
		fmt.Println("VAlor Recebido: ", v)
	}
}

func examplePiker() {
	canal := converge(trabalho("maça"), trabalho("pêra"))

	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Funcao %v diz %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}(s, canal)

	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)

	go func() {
		for {
			novo <- <-x
		}
	}()

	go func() {
		for {
			novo <- <-y
		}
	}()

	return novo
}

func envia(p, i chan int) {
	x := 100

	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}

	close(p)
	close(i)
}

func receive(p, i, c chan int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
