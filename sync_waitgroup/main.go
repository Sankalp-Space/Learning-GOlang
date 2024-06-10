package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n ", i)
	//some taskis happening here
	fmt.Printf("worker %d end\n ", i)

}

func main() {
	fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	//start 3 worker goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1) //increment the Waitgroup count
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("worker task completed")

}
