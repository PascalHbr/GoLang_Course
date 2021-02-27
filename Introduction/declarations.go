package main

import "fmt"

var y = 43
var z = `James said: "Shaken, not stirred."`

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	foo()
}

func foo() {
	fmt.Println("again", y)
}
