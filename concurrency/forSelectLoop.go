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

	go func() {
		arr := []string{"a", "b", "c", "d"}
		for _, el := range arr {
			c <- el
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
