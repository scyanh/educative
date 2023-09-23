package main

import (
	"fmt"
	"sync"
)

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

func (fanInFanOut) FanIn(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for el := range ch1 {
			ch <- el
		}
	}()

	go func() {
		defer wg.Done()
		for el := range ch2 {
			ch <- el
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func (f fanInFanOut) SimpleFanInFanOut() {
	fmt.Println("start SimpleFanInFanOut ...")

	c := f.FanIn(f.GenChan(), f.GenChan())

	for el := range c {
		fmt.Println(el)
	}

	fmt.Println("... end SimpleFanInFanOut")
}
