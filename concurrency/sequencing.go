package concurrency

import "fmt"

type sequencing struct {
}

func NewSequencing() sequencing {
	return sequencing{}
}

type CookInfo struct {
	food int
	done chan bool
}

func (sequencing) GeneratorFood() <-chan CookInfo {
	c := make(chan CookInfo)

	go func() {
		for i := 0; i < 3; i++ {
			chDone := make(chan bool)
			c <- CookInfo{i, chDone}
			<-chDone
		}
	}()

	return c
}

func (s sequencing) FanInFood(ch1, ch2 <-chan CookInfo) <-chan CookInfo {
	ch := make(chan CookInfo)

	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()

	return ch
}

func (s sequencing) SequenceFood() {

	c := s.FanInFood(s.GeneratorFood(), s.GeneratorFood())

	for i := 0; i < 3; i++ {
		ch1 := <-c
		ch2 := <-c

		fmt.Println(ch1.food)
		fmt.Println(ch2.food)

		ch1.done <- true
		ch2.done <- true

	}

	fmt.Println("done")
}
