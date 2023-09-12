package main

import (
	"fmt"
	"syscall/js"
)

// import js syscall lib

func main() {
	fmt.Println("Hello world from Go :)")

	// 1- find document in page
	document := js.Global().Get("document")
	// 2 - create p
	p := document.Call("createElement", "p")
	// 3- add some text to p
	p.Set("innerHTML", "Hello WASM from Go!")
	// 4- do some style
	p.Set("style", "font-size:30px;border-top:1px solid #ddd;padding-top:5px")
	//5- finally : append  P to the body
	document.Get("body").Call("appendChild", p)
}

//export add
func add(x, y int) int {
	return x + y
}
