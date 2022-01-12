package main

import "fmt"

// // contant variable

const defaultValue = 1

// iota
const (
	sunday  = iota // 0
	monday         // 1
	tuesday        // 2
)

// // return void
func greeting(name string, age int) {
	fmt.Printf("My name is %s, I'am %d years old", name, age)
}

// // return 1 result
func add(a, b int) int {
	return a + b
}

// // return multi results eg. 2 result of int
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Printf("print constant variable: %d \n", defaultValue)
	fmt.Printf("sunday: %d, monday: %d, tuesday: %d \n", sunday, monday, tuesday)

	// // variable
	var i_v int = 14
	var s_v string = "hello"
	var ok_v bool = true
	fmt.Printf("%d %s %t \n", i_v, s_v, ok_v)
	// // inference **** popular
	i_i := 14
	s_i := "hello"
	ok_i := false
	fmt.Printf("%d %s %t \n", i_i, s_i, ok_i)

	greeting("tay", 30)

	fmt.Println(add(1, 2))

	first := 1
	second := 2

	s_first, s_second := swap(first, second)

	fmt.Printf("(%d, %d) -> (%d, %d)\n", first, second, s_first, s_second)
}
