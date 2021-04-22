package main

import "fmt"

func main() {
	x := 0
	recebeXPorParametro(x)
	fmt.Println("Na função X=", x)
	recebeXComoUmPonteiroParaX(&x)
	fmt.Println("Com Ponteiro X=", x)

}

func recebeXPorParametro(x int) {
	x++
}

func recebeXComoUmPonteiroParaX(x *int) {
	*x++
}
