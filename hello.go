package main 

import "fmt"

// The first letter that is capitalized will be exported by Go" //
// func foo() {
// 	fmt.Println("The sqaure root of 4 is" ,math.Sqrt(4))
// }

// Main will be the only function called //
// func main() {
// 	foo()
// }

// func main() {
// 	fmt.Println("A number between 0 - 99 is",rand.Intn(100))
// }


// need to state what kind of data you'll be using (x, y), and then need to state what you would like return outside the parenthesis //
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a,b
}

// use the : as long as the datatype isn't going to change//
// func main() {
// 	num1, num2 := 5.6, 9.5
	
// 	w1, w2 := "Hey", "there!"

// 	fmt.Println(multiple(w1, w2))
// 	fmt.Println(add(num1, num2))
// }

func main() {

	// this will convert a to a float64
	var a int = 62
	var b float64 = float64(a)

	// x will be trype int
	x := a
}