package concurrency

import "fmt"

type FooBar struct {
	cnt     int
	fooChan chan int
	barChan chan int
}

func (this *FooBar) Foo() {
	for i := 0; i < this.cnt; i++ {
		<-this.fooChan
		fmt.Printf("foo")
		this.barChan <- 1
	}
	// <-this.fooChan
}

func (this *FooBar) Bar() {
	for i := 0; i < this.cnt; i++ {
		<-this.barChan
		fmt.Printf("bar")
		this.fooChan <- 1
	}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{cnt: n, fooChan: make(chan int, 0), barChan: make(chan int, 0)}
}
