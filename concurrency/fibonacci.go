package main

import "fmt"

type concurrency struct {
}

func NewConcurrency() concurrency {
	return concurrency{}
}

func (concurrency) Fibonnaci(number int) chan int {
	chanFibo := make(chan int)

	k := 0
	go func() {
		for i, j := 0, 1; k < number; k++ {
			chanFibo <- i
			i, j = i+j, i
		}
		close(chanFibo)
	}()

	return chanFibo
}

func (concurrency) Fibonnaci2(number int) chan int {
	c := make(chan int)

	k := 0
	go func() {
		for i, j := 0, 1; k < number; k++ {
			c <- i
			i, j = i+j, i
		}
		close(c)
	}()

	return c
}

func (c concurrency) GetFibonacci(k int) {
	fmt.Println("start GetFibonacci ...")
	for el := range c.Fibonnaci(k) {
		fmt.Println(el)
	}
	fmt.Println("... end GetFibonacci")
}
