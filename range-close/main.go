package main

import "fmt"

func main() {
	c := make(chan int)
	go loop(10, c)
	prints(c)
}

func loop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal: ", v)
	}
}
