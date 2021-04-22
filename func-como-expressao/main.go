package main

import "fmt"

func main() {
	x := 456

	y := func(x int) {
		fmt.Println(x)
	}

	y(x)
}
