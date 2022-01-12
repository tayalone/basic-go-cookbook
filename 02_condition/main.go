package main

import (
	"fmt"
	"strconv"
)

func ifLikeAProfessionalGoLang() {
	// // ปกติต้อง
	// n, err := strconv.Atoi("5s");
	// if err!= { }
	// //
	if n, err := strconv.Atoi("5s"); err != nil {
		fmt.Println("NaN:", err)
	} else {
		fmt.Println(n)
	}
}

func isOdd(n int) bool {
	if n == 1 || n == 2 {
		return true
	} else {
		if n%2 == 1 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	ifLikeAProfessionalGoLang()
	fmt.Println("1 isOdd: ", isOdd(1))
	fmt.Println("2 isOdd: ", isOdd(2))
	fmt.Println("3 isOdd: ", isOdd(3))
	fmt.Println("3 isOdd: ", isOdd(4))
}
