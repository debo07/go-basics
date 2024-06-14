package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sidelength float64
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape, shapeType string) {
	area := s.getArea()
	fmt.Printf("Area of the %s is %.2f\n", shapeType, area)
}

func main() {
	t := triangle{height: 10.0, base: 5.5}
	printArea(t, "triangle")
	s := square{sidelength: 5.5}
	printArea(s, "square")
}

