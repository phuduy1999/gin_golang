package main

import (
	// "fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (c *circle) area() float64 {
	return c.radius * math.Pi * c.radius
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func sum(sh []shape) float64 {
	sum := float64(0)
	for _, v := range sh {
		sum += v.area()
	}
	return sum
}

// func main() {
// 	c1 := circle{4.5}
// 	r1 := rect{5, 7}
// 	shapes := []shape{&c1,&r1}
// 	fmt.Println(sum(shapes));

// }
