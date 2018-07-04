package main

import (
	"syscall/js"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	inputArea := js.Global().Get("document").Call("getElementById", "input")
	inputText := inputArea.Get("value").String()
	result := format(inputText)

	outputDiv := js.Global().Get("document").Call("getElementById", "output")
	outputDiv.Set("innerHTML", result)
}

func format(input string) string {
	outputBytes := blackfriday.Run([]byte(input))
	output := string(outputBytes)
	return output
}
