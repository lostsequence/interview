package main

import "fmt"

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
}
