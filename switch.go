package main 

import "fmt"

func main() {
	fmt.Print("You win: ")
	doorNumber := 1
	switch doorNumber {
	case 1:
		fmt.Println("Prize 1")
		fallthrough  // Allows you to drop to next case down
	case 2: 
		fmt.Println("Prize 2")
	default:
		fmt.Println("Prize 3")	
	}
}

