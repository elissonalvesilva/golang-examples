package main

import "fmt"

func main() {
	// simpleGoComaComCanais()
	funcao3()
}

func funcao3() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)
	receive(par, impar, quit)
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
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra: ", v)
			} else {
				fmt.Println("Encerrando. Recebemos: ", v)
			}
			return
		}
	}
}

func simpleGoComaComCanais() {
	channel := make(chan int)

	go func() {
		channel <- 42
		close(channel)
	}()

	v, ok := <-channel
	fmt.Println(v, ok)
}
