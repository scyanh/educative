package main

import (
	"fmt"
	"math/rand"
	"time"
)

type quitChannel struct {
	m map[int]int
}

func NewQuitChannel() quitChannel {
	return quitChannel{}
}

func (q quitChannel) DoRace(idx int, racer, quit chan int) {
	racer <- idx
	// do something
	rand.Seed(time.Now().UnixNano())
	mini, maxi := 5, 20
	r := rand.Intn(maxi - mini)
	r += mini
	time.Sleep(time.Duration(r) * time.Millisecond)

	q.m[idx] = r

	quit <- idx
}

func (q quitChannel) Racing() {
	fmt.Println("start Racing ...")
	defer fmt.Println("... end Racing")

	c := make(chan int)
	quit := make(chan int)
	q.m = make(map[int]int)

	for i := 0; i < 3; i++ {
		go q.DoRace(i, c, quit)
	}

	for {
		select {
		case racer := <-c:
			fmt.Println(racer)
		case finished := <-quit:
			fmt.Printf("quit by %d in %d milliseconds", finished, q.m[finished])
			return
		}
	}
}
