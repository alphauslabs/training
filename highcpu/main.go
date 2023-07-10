package main

import (
	"runtime"
	"time"
)

func main() {
	done := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		// for i := 0; i < 2; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
	close(done)
}
