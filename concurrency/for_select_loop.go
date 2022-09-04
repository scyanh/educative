package concurrency

import "fmt"

type forSelectLoop struct {
}

func NewForSelectLoop() forSelectLoop {
	return forSelectLoop{}
}

func (forSelectLoop) Fruits() {
	defer fmt.Println("done")
	c := make(chan string)

	go func(){
		arr:=[]string{"a", "b", "c", "d"}
		for _,el:=range arr{
			c<-el
		}
	}()

	for {
		select {
		case el := <-c:
			fmt.Println(el)
			if el == "d" {
				return
			}
		}
	}
}
