package main

import (
	"fmt"
	"time"
	"singleton/designPettern"
)

func test(ch chan<- interface{}) {
	_, message := designPettern.GetInstance()
	ch <- message
}

func main() {
	ch := make(chan interface{})
	go test(ch)
	go test(ch)
	go test(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	time.Sleep(3 * time.Second)
}
