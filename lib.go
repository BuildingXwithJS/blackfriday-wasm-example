package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("blackfridayFormat", js.NewCallback(format))

	<-c
	fmt.Println("Bye Wasm !")
}

func format(input []js.Value) {
	fmt.Println("Format hello")
}
