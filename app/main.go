package main

import (
	"agedito/udemy/modules/toolkit"
	"fmt"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println("Random string: ", s)
}
