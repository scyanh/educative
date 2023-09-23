package main

import "fmt"

type quitChannel struct {
}

func NewQuitChannel() quitChannel {
	return quitChannel{}
}

func (quitChannel) DoRace(idx int, racer, quit chan int) {
	racer <- idx
	// do something
	quit <- idx
}

func (q quitChannel) Racing() {
	fmt.Println("start Racing ...")
	defer fmt.Println("... end Racing")
	c := make(chan int)
	quit := make(chan int)

	for i := 0; i < 3; i++ {
		go q.DoRace(i, c, quit)
	}

	for {
		select {
		case racer := <-c:
			fmt.Println(racer)
		case finished := <-quit:
			fmt.Println("quit by ", finished)
			return
		}
	}
}
