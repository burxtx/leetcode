package concurrency

import (
	"testing"
)

func TestFoo(t *testing.T) {
	fb := NewFooBar(5)
	go fb.Foo()
	go fb.Bar()
	// start off here
	fb.fooChan <- 1
}
