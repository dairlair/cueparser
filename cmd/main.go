package main

import (
	"fmt"

	"github.com/dairlair/cueparser"
)

func main() {
	fmt.Println("Hi!")
	tokens := cueparser.Tokenize("qwe")
	fmt.Printf("Tokens found: %+v\n", tokens)
}
