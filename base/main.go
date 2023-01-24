package main

import (
	"fmt"
	"unsafe"
)

// определить размер будет у структуры
type MyStruct struct {
	num1  byte
	isNum bool
	num2  int
}

func main() {
	s := MyStruct{}
	fmt.Println(unsafe.Sizeof(s))
}
