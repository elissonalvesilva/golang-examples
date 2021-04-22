package main

import "fmt"

func main() {
	t := somaSomentePares(soma, []int{50, 51, 52, 53, 54, 56, 57, 58, 59, 60}...)
	fmt.Println(t)

	printSomenteElisson(print, []string{"Danilo", "Eduarda", "Elisson"}...)
}

func print(x string) {
	fmt.Println(x)
}

func printSomenteElisson(callback func(name string), names ...string) {
	for _, name := range names {
		if name == "Elisson" {
			callback(name)
		}
	}
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}

	return n
}

func somaSomentePares(callback func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := callback(slice...)
	return total
}
