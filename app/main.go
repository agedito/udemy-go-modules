package main

import (
	"fmt"
	"github.com/agedito/udemy-go-modules"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println("Random string: ", s)
}
