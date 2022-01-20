package main

import "fmt"

func whichType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("this is a string %s\n", v)
	case int:
		fmt.Printf("this is a int %d\n", v)
	default:
		fmt.Printf("Idiot \n")

	}
}

func main() {
	whichType("sadasdasd")
	whichType(12)
	whichType(12.2)
}
