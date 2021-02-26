package main

import "fmt"

type hotdog int

var b hotdog

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
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	y = int(b)
	fmt.Println(y)
}
