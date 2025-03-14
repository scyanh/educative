package main

import (
	"fmt"
	"time"
)

type timeoutSelect struct {
}

func NewTimeoutSelect() timeoutSelect {
	return timeoutSelect{}
}

func (timeoutSelect) Dynamite() {
	fmt.Println("start Dynamite ...")
	defer fmt.Println("... end Dynamite")

	c := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		select {
		case val, ok := <-c:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(val)
		case <-time.After(120 * time.Microsecond):
			fmt.Println("time out")
			return
		}
	}

}

func (timeoutSelect) Dynamite2() {
	fmt.Println("start")
	defer fmt.Println("end")

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		select {
		case val, ok := <-c:
			if ok {
				fmt.Println(val)
			} else {
				fmt.Println("c closed")
				return
			}

		case <-time.After(1000 * time.Microsecond):
			return
		}
	}

}
