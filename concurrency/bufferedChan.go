package main

import "fmt"

type bufferedChan struct {
}

func NewBufferedChan() bufferedChan {
	return bufferedChan{}
}

func (bufferedChan) SimpleBuffered() {
	fmt.Println("start SimpleBuffered ...")
	c := make(chan int, 33)

	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
	fmt.Println("... end SimpleBuffered")

}
