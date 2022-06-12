package main

import (
	"time"
	"singleton/designPettern"
)

func test(ch chan<- interface{}) {
	ch <- designPettern.GetInstance()
}

func main() {
	ch := make(chan interface{})
	go test(ch)
	go test(ch)
	go test(ch)
	<-ch
	<-ch
	<-ch

	time.Sleep(3 * time.Second)
}
