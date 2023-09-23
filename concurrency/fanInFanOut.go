package main

import "fmt"

type fanInFanOut struct {
}

func NewFanInFanOut() fanInFanOut {
	return fanInFanOut{}
}

func (fanInFanOut) GenChan() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func (fanInFanOut) FanIn(c1, c2 <-chan int) <-chan int {
	cOut := make(chan int)

	go func() {
		for el := range c1 {
			cOut <- el
		}
	}()
	go func() {
		for el := range c2 {
			cOut <- el
		}
	}()

	return cOut
}

func (f fanInFanOut) SimpleFanInFanOut() {
	fmt.Println("start SimpleFanInFanOut ...")
	c := f.FanIn(f.GenChan(), f.GenChan())

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("... end SimpleFanInFanOut")
}
