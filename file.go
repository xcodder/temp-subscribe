package main

import (
	"fmt"
	"sync"
	"time"
)

type Thing struct {
	Ch chan bool
	l  int
	*sync.Mutex
}

func NewThing() *Thing {
	return &Thing{Ch: make(chan bool)}
}
func (t *Thing) Wait() {
	t.l++
	<-t.Ch
}
func (t *Thing) Broadcast() {
	for i := 0; i < t.l; i++ {
		t.Ch <- false
	}
	t.l = 0
}
func DoA(a int, c *Thing) {
	fmt.Println("Waitin on ", a)
	c.Wait()
	fmt.Println("Ok, doing ", a)
	time.Sleep(time.Second)
	fmt.Println("Ok, finished ", a)
}

func Runner(c *Thing) {
	c.Broadcast()
}
func main() {
	c := NewThing()
	go DoA(1, c)
	go DoA(2, c)
	go DoA(3, c)

	time.Sleep(time.Second)

	go Runner(c)

	time.Sleep(time.Second * 5)
	fmt.Println("Here")
}
