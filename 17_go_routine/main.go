package main

import (
	"fmt"
	"time"
)

func doSomeThing() {
	fmt.Println("call from do Some Thing")
}

func main() {
	fmt.Println("Start")

	go doSomeThing()

	time.Sleep(time.Second)
	fmt.Println("End")
}
