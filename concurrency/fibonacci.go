package concurrency

import "fmt"

func (concurrency) Fibonnaci(number int) chan int {
	chanFibo := make(chan int)

	k := 0
	go func(){
		for i, j := 0, 1; k < number; k++ {
			chanFibo <- i
			i, j = i+j, i
		}

		close(chanFibo)
	}()

	return chanFibo
}

func (c concurrency) GetFibonacci(k int) {
	for el := range c.Fibonnaci(k) {
		fmt.Println(el)
	}
}
