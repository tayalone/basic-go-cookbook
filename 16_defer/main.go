package main

import "fmt"

func main() {
	defer fmt.Println("defer#1")
	defer fmt.Println("defer#2")

	// // ใช้มากๆกับ พวก io/database
	fmt.Println("before defer")
}
