package main 

import ("fmt"
				"strings")

type Name string 
type Minutes int

// name in main is passed in as t
func (t Name) FixCase() string {
	return strings.Title(string(t)) // Convert the Title object to a string
}


func main() {
	name := Name("the matrix")
	fixed := name.FixCase()
	fmt.Println(fixed)
}	


// Change to a pointer to use the actual object value and not the passed argument value
// func (m *Minutes) Incremenent() {
// 	*m = (*m + 1) % 60
// }

// func main() {
// 	minutes := Minutes(58)

// 	for i := 1; i <= 3; i++{
// 		(&minutes).Incremenent()
// 		fmt.Println(minutes)
// 	}	
// }