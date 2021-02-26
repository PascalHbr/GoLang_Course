package main

import "fmt"

type effect int
var x effect

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Println(x)
	y := int(42)
	fmt.Printf("%v\t%T\n", y, y)
}
