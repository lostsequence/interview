package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for true {
			fmt.Println("Infinite loop")
		}
	}()

	time.Sleep(time.Millisecond * 100)

	fmt.Println("Should I get here ?")
}
