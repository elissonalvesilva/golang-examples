package main

import "fmt"

func main() {
	// funcao1()
	// funcao2()
	funcao3()
}

func funcao1() {
	x := 500
	a := make(chan int)
	b := make(chan int)
	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("CAnal A: ", v)
		case v := <-b:
			fmt.Println("CAnal B: ", v)

		}
	}
}

func funcao2() {
	c := make(chan int)
	quit := make(chan int)
	go recebeQuit(c, quit)
	enviaParaCanal(c, quit)
}

// correcao com comma ok
func funcao3() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)
	receive(par, impar, quit)
}

func recebeQuit(c chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido: ", <-c)
	}
	quit <- 0
}

func enviaParaCanal(c chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case c <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}

func mandaNumeros(par chan int, impar chan int, quit chan bool) {
	total := 100

	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receive(par chan int, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O numero V:", v, "é par")
		case v := <-impar:
			fmt.Println("O numero V:", v, "é Impar")
		case <-quit:
			return
		}
	}
}
