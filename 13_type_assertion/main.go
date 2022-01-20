package main

import "fmt"

type Obj interface {
	Area() float64
}

type reactagle struct {
	w, l float64
}

func (r reactagle) Area() float64 {
	return r.l * r.w
}

type triangle struct {
	b, h float64
}

func (t triangle) Area() float64 {
	return (t.b * t.h) / 2
}

func main() {
	var i interface{}
	i = "text"
	var s string

	if v, ok := i.(string); ok { // // <- dis is type  assertion
		s = v
		fmt.Println("s is string", s)
	}
	var r, t Obj
	r = reactagle{w: 4, l: 6}
	t = triangle{b: 4, h: 6}
	if _, ok := r.(reactagle); ok {
		fmt.Println(r.Area())
	}
	if _, ok := t.(triangle); ok {
		fmt.Println(t.Area())
	}
}
