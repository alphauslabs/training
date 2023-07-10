package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const num = 5000
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("count:", counter)
}
