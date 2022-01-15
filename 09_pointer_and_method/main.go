package main

import "fmt"

// dis is strunc
type reactangle struct {
	w, l float64
}

// dis is method
// // รับไป pointer ไปเลยกันลืม
func (rec *reactangle) area() float64 {
	return rec.l * rec.w
}

func main() {
	rec := reactangle{
		w: 4,
		l: 5,
	}

	rec.w = 6
	fmt.Println("rec.w", rec.w)
	fmt.Println("rec.l", rec.l)

	fmt.Println("Area of reactagle: ", rec.area())
}
