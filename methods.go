package main 

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // min 0 max 65535, always going forward
	break_pedal uint16
	steering_wheel int16 // -32k - +32k
	top_speed_km float64
}

// Value Receiver  OR Method on Value//
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_km / usixteenbitmax)
}

// Pointer Receiver OR Method on Pointer//
func (c *car) mph() float64 {
	c.top_speed_km = 500
	return float64(c.gas_pedal) * (c.top_speed_km / usixteenbitmax / kmh_multiple)
}

func (c *car) new_top_speed(new_speed float64) { // Don't forget pointer (*) since reading car type
	c.top_speed_km = new_speed
}

func main() {
	a_car := car{gas_pedal: 65000, 
							break_pedal: 0, 
							steering_wheel: 12561, 
							top_speed_km: 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())												
}