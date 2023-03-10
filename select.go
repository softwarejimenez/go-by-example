package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("receive c1", msg)
		case msg := <-c2:
			fmt.Println("receive c2", msg)
		}
	}
}
