package main

import "fmt"

func main() {
	var s string
	var p *string

	p = &s
	*p = "hello word"
	fmt.Println("value in pointer is ", *p)
	fmt.Println("value in string is", s)

	v := 2
	fmt.Println("value before call func 'double' ", v)
	double(&v)
	fmt.Println("value after call func 'double' ", v)
}

func double(v *int) {
	*v = *v * 2
}
