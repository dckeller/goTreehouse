package main 

import "fmt"

func main() {

	ages := map[string] float64{}
	ages["Alice"] = 12
	ages["Bob"] = 32
	fmt.Println(ages)
	fmt.Println(ages["Alice"], ages["Bob"])

	for name, age := range ages{
		fmt.Println(name, age)
	}
}

func HalfPriceSale(prices map[string]float64) map[string]float64 {
  
  halfPriced := make(map[string]float64) //Create an empty map
  
  for product, price := range prices { // Iterate through prices
  	halfPriced[product] = price / 2 // Assign the updated price (value), to the products (keys) all to the new map. 
  }
  return halfPriced
}