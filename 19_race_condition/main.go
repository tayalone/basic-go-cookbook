package main

import (
	"fmt"
	"sync"
)

var i int

var mux sync.Mutex

type thing struct {
	i   int
	mux sync.Mutex
}

func main() {
	t := thing{
		i: 0,
		// mux: sync.Mutex{},
	}
	// fmt.Println(t.write)))
	t.write(1)
	fmt.Println(t.read())
	go func() {
		for {
			// fmt.Println(read())
			fmt.Println("thing", t.read())
		}
	}()

	// for i := 0; i < 10; i++ {
	// 	write(i)
	// 	// t.write(i)
	// }
}

func (t *thing) write(n int) {
	t.mux.Lock()
	t.i = n
	t.mux.Unlock()
}

func (t thing) read() int {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.i
}

func write(n int) {
	mux.Lock()
	i = n
	mux.Unlock()
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}
