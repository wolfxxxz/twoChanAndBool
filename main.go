package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i < 11; i++ {
			fmt.Println("write ", i)
			ch <- i
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			fmt.Println("read ", num)
			time.Sleep(time.Second * 1)
		}
		done <- true

	}()
	<-done
}
