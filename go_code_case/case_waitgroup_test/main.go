package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Job 1 is done.")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Job 2 is done.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all Done.")
}
