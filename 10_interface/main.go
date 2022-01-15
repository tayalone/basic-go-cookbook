package main

import (
	"fmt"

	"github.com/tayalone/basic-go-cookbook/10_interface/demo"
)

//  interface เป็น type ประเภทนึงของ golang
// // interface มีหน้าที่ กำหนดพฤติกรรมของ type ด้วย function

// // นี้คือ empty inter face
var EmptyInterface interface{}

type reactagle struct {
	width, heigth float64
}

func (r *reactagle) Area() float64 {
	return r.width * r.heigth
}

type triangle struct {
	base, heigth float64
}

func (t *triangle) Area() float64 {
	return (t.base * t.heigth) / 2
}

// // type Stape เป็น Interface มีต้องมี function Area()
// // reactangle == Shape
// // triangle == Shape

type Shape interface {
	Area() float64
}

func myArea(s Shape) {
	fmt.Println("my area is ", s.Area())
}

func main() {
	rec := reactagle{width: 10, heigth: 30}
	myArea(&rec)
	tri := triangle{base: 10, heigth: 30}
	myArea(&tri)
	// // Demo
	demo.Execute()
}
