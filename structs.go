package main 

import "fmt"

type car struct {
	gas_pedal uint16 // min 0 max 65535, always going forward
	break_pedal uint16
	steering_wheel int16 // -32k - +32k
	top_speed_km float64
}

func main() {
	a_car := car{gas_pedal: 22341, 
							break_pedal: 0, 
							steering_wheel: 12561, 
							top_speed_km: 225.1}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.top_speed_km)						
}