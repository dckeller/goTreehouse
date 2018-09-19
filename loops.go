rpackage main 

import "fmt"

// func main() {
// 	for i:= 1; i <= 3; i++ {
// 		fmt.Println(i)
// 	}
// }

func main() {
	months := [3]string{"May", "June", "July"}
	sales := [3]float64{122.22, 145.54, 17.62}
	
	// for i := 0; i < len(months); i++ {
	// 	fmt.Println(months[i], sales[i])
	// }

	for i, month := range months {
		fmt.Println(month, sales[i])
	}
}