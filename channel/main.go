package main

func main() {
	canal := make(chan int, 1)

	// usando goroutine
	// go func() {
	// 	canal <- 42
	// }()
	canal <- 42
	x := <-canal
	println(x)
}
