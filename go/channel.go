package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := make(chan string)
	s2 := make(chan string)

	go func(value string) {
		for{
			s1 <- value
			time.Sleep(time.Second * 2)
		}
	}("#Sleep 1#")

	go func() {
		for{
			s2 <- "Sleep 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for{
			select {
			case msg1 := <- s1:
				fmt.Println(msg1)

			case msg2 := <- s2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

