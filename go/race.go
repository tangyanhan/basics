package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			fmt.Println("count++=", count)
			time.Sleep(time.Millisecond)
			count--
			fmt.Println("count--", count)
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
