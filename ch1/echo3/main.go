package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
	fmt.Print(os.Args[0])

	// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
	for index, value := range os.Args {
		fmt.Println(index, value)
	}

	// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join
	func_a()
	func_b()
}

func func_a() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.8f", secs)
}

func func_b() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("%.8f", secs)
}
