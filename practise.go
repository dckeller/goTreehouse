package main 

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"time"
	"reflect"
	"net"
)

var myNumber = 1.23

func main () {
	// newNumber := 2.23
	float := 2.0
	num := int(float)

	fmt.Println(stringutil.Reverse("hello"))
	fmt.Println(reflect.TypeOf(time.Now()))
	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(num)
}