package main

import (
	"fmt"
	"log"
	"math"
)

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("Can't take square of negative number.")
	}
	return math.Sqrt(x), nil
}

func main() {
	squareRoot, err := squareRoot(9) // Here we set to variables
	if err != nil {                  // If err is not nil, run this
		log.Fatal(err)
	}
	fmt.Println(squareRoot)
}
