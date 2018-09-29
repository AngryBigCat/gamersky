package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for {
				result := <-c
				fmt.Println(result)
			}
		}()
	}

	for i := 1; i <= 1000; i++ {
		c <- i
	}

}