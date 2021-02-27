package main

import "fmt"

type hotdog int

var z hotdog

func main() {
	y := 42
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)
	z = 43
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	y = int(z)
	fmt.Println(y)
}
