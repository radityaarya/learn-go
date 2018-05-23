package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflect.TypeOf
	str := " h e l l o"
	in := 12345
	bool := true
	sl := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(in))
	fmt.Println(reflect.TypeOf(bool))
	fmt.Println(reflect.TypeOf(sl))

}
