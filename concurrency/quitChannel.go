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

	rand.Seed(time.Now().UnixNano())
	mini, maxi := 5, 20
	r := rand.Intn(maxi - mini)
	r += mini

	q.m[idx] = r
	time.Sleep(time.Duration(r) * time.Millisecond)

	quit <- idx
}

func (q quitChannel) Racing() {
	fmt.Println("start Racing ...")
	defer fmt.Println("... end Racing")

	racer := make(chan int)
	quit := make(chan int)
	q.m = make(map[int]int)

	for i := 0; i < 3; i++ {
		go q.DoRace(i, racer, quit)
	}

	for {
		select {
		case r := <-racer:
			fmt.Println(r)
		case finish := <-quit:
			fmt.Printf("quit by %d took %d milliseconds", finish, q.m[finish])
			return
		}
	}

}
