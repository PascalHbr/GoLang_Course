package main

import "fmt"

func main() {
	// first method
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// second method
	k := 0
	for k <= 5 {
		fmt.Println(k)
		k++
	}

	// third method
	x := 0
	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}
}
