package main

import "fmt"

func main() {
	rw := make(chan string)

	go greeting(rw)
	rw <- "John" // input way

	fmt.Println(<-rw) // output way

	ch := make(chan int)
	qCh := make(chan struct{}) // signal chanel
	go fibo(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println("inner for", i)
		fmt.Println(<-ch)
	}
	qCh <- struct{}{}
}

func greeting(rw chan string) {
	s := <-rw
	rw <- "Hello, " + s
}

func fibo(ch chan int, qCh chan struct{}) {
	a, b := 0, 1
	fmt.Println("fibo func", a, b)

	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			fmt.Println("signal มาแล้ว ปิดกิจการ")
			return
		}
	}
}
