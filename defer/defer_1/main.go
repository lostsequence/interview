package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("hello world")
	log.Fatal("Oops")
}
