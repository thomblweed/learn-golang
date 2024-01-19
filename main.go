package main

import (
	"fmt"

	car "github.com/thomblweed/learn-golang/src"
)

func doStuff(message string, ch chan struct{}) {
	<-ch

	fmt.Println(message)
}

func main() {
	car.MyCar()

	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	doStuff("hello", ch)

	chTwo := make(chan struct{})
	go func() {
		chTwo <- struct{}{}
	}()
	doStuff("hello again", chTwo)
}
