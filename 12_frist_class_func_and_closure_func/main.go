package main

import "fmt"

var add = func(a, b int) int {
	return a + b
}

type Math struct {
	add func(int, int) int
}

func coujureFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("1 + 2 is: ", add(1, 2))
	m := Math{add: add}
	fmt.Println("2 + 3 is: ", m.add(2, 3))
	fn := coujureFunc()
	fmt.Println("coujureFunc 1", fn())
	fmt.Println("coujureFunc 2", fn())
	fmt.Println("coujureFunc 3", fn())
}
