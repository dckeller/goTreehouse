// package main 

// import "fmt"

// type Monitor struct {
// 	Resolution string // This is field 1
// 	Connector string // This is field 2
// 	Value float64 // This is field 3
// }

// func showFields(m Monitor) {
// 	fmt.Println(m.Resolution, m.Connector, m.Value)
// }

// func main() {
// 	monitor := Monitor{"1080p", "HDMI", 234.48}
// 	showFields(monitor)
// }

package main

import "fmt"

type Clock struct {
	Hours int
	Minutes int
}

func Noon() (Clock) {
	time := Clock{}
	time.Hours = 12
	time.Minutes = 0
	return time
}

func main() {
	fmt.Println(Noon())
}
