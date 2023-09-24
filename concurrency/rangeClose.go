package main

import "fmt"

type rangeClose struct {
}

type Money struct {
	idx    int
	amount int
}

func NewRangeClose() rangeClose {
	return rangeClose{}
}

func (rangeClose) GenMoney(c chan Money) {
	for i := 0; i < 3; i++ {
		c <- Money{
			idx:    i,
			amount: 500,
		}
	}
	close(c)
}

func (r rangeClose) RangeCloseMoney() {
	fmt.Println("start RangeCloseMoney ...")
	c := make(chan Money)
	go r.GenMoney(c)

	for el := range c {
		fmt.Printf("%d generate %d\n", el.idx, el.amount)
	}

	fmt.Println("... end RangeCloseMoney")
}
