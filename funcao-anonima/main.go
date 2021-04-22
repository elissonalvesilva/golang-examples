package main

import "fmt"

func main() {
	x := 45

	func(x int) {
		fmt.Println(x)
	}(x)
}
