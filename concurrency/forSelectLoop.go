package main

import "fmt"

type forSelectLoop struct {
}

func NewForSelectLoop() forSelectLoop {
	return forSelectLoop{}
}

func (forSelectLoop) Fruits() {
	fmt.Println("start Fruits ...")
	defer fmt.Println("... end Fruits")

	c := make(chan string)

	ar := []string{"a", "b", "c", "d"}

	go func() {
		for _, el := range ar {
			c <- el
		}
		close(c)
	}()

	for {
		select {
		case val := <-c:
			fmt.Println(val)
			
			if val == "d" {
				return
			}
		}
	}

}
