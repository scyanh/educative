package concurrency

import (
	"fmt"
	"time"
)

type timeoutSelect struct {

}

func NewTimeoutSelect() timeoutSelect{
	return timeoutSelect{}
}

func (timeoutSelect) Dynamite(){
	defer fmt.Println("end")
	c:=make(chan int)

	go func(){
		for i:=0; i<1000; i++{
			c<-i
		}
	}()

	for{
		select {
		case el:=<-c:
			fmt.Println(el)
		case <-time.After(time.Microsecond*10):
			return
		}
	}

}
