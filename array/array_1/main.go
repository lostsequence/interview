package main

import (
	"fmt"
	"reflect"
)

// определить типы массивов
func main() {
	var a [4]int
	fmt.Println(reflect.TypeOf(a))

	var b [5]int
	fmt.Println(reflect.TypeOf(b))
}
