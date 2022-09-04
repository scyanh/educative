package concurrency

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

func (rangeClose) GetMoney(c chan<- Money) {
	for i := 0; i < 3; i++ {
		c <- Money{i, 500}
	}
	close(c)
}

func (r rangeClose) RangeCloseMoney() {
	c := make(chan Money)
	go r.GetMoney(c)

	for el := range c {
		fmt.Println(fmt.Sprintf("%d %d", el.idx, el.amount))
	}

	fmt.Println("end")
}
