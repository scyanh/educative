package concurrency

import "fmt"

type bufferedChan struct {

}

func NewBufferedChan() bufferedChan{
	return bufferedChan{}
}

func (bufferedChan) SimpleBuffered(){
	c:=make(chan int, 3)

	c<-1
	c<-2
	c<-3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
	fmt.Println("done")

}
