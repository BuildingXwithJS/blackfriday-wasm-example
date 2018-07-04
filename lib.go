package main

import (
	"syscall/js"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("blackfridayFormat", js.NewCallback(format))
	<-c
}

func format(input []js.Value) {
	message := input[0].String()
	outputBytes := blackfriday.Run([]byte(message))
	output := string(outputBytes)
	input[1].Invoke(output)
}
