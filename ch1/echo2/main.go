package main

import (
	"fmt"
	"os"
)

func main() {
	var s string = ""
	var sep string = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Print(my_name)

}

var my_name string = "Benjamin"
