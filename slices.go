package main

import "fmt"

func main() {
	// a := [5]int{1, 2, 3, 4, 5}
	// s1 := a[0:3]
	// s2 := a[2:5]
	// a[2] = 88
	// fmt.Println(a, s1, s2)
	// fmt.Println("This array has: ", len(s1), "items, and has a cap of: ", cap(s1))
	// s1 = append(s1, 5)
	// fmt.Println("This array has: ", len(s1), "items, and has a cap of: ", cap(s1))
	// fmt.Println(a, s1, s2)

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	for i, v := range s {
		fmt.Println("Element: ", i, "Value: ", v)
	}
}
