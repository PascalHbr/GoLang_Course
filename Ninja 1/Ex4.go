package main

import "fmt"

type effect int
var x effect

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Println(x)
}
