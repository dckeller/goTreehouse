package main 

import "fmt"

func main() {
	x := 15
	a := &x // memory address

	fmt.Println(a) // Points to memory address
	fmt.Println(*a) // Will point to the actual line

	*a = 5 // This will change the value of a --> x
	fmt.Println(x)

	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
}