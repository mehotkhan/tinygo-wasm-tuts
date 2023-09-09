package main

import "fmt"

func main() {
	fmt.Println("Hello world from Go :)")
}

//export add
func add(x, y int) int {
	return x + y
}
