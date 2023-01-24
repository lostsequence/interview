package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("Starting listener")
	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
		ch <- 1
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
		ch <- 2
		wg.Done()
	}()

	fmt.Println(<-ch)
	wg.Wait()
}
