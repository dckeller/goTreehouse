package main

import (
	"fmt"
	"./parts"
)

type Part interface {
	Specs() string // params if any, and return value (string)
	Price() string 
}

func Summary(p Part) string {
	return p.Specs() + "\n" + p.Price()
}


func main() {
	catalog := []Part{}
	catalog = append(catalog, parts.Monitor{"HDMI", "1080p", 249.99})
	catalog = append(catalog, parts.HardDrive{"SSD", "SATA", 149.99})
	catalog = append(catalog, parts.Keyboard{108, "Mechanical", 99.99})

	// Ignoe the index
	for _, part := range catalog {
		fmt.Println(Summary(part))
		fmt.Println("---------------")
	}
}