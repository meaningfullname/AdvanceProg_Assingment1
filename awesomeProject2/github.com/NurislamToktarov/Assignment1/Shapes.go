package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) area() float64 {

	return float64(r.Width * r.Height)
}
func (r Rectangle) perimeter() float64 {
	return float64(r.Width + r.Height)
}
func (c Circle) area() float64 {
	return float64(c.Radius * c.Radius)
}
func (c Circle) perimeter() float64 {
	return float64(2 * c.Radius * c.Radius)
}
func (t Triangle) area() float64 {
	per := (t.SideA + t.SideB + t.SideC) / 2.0
	return float64((math.Sqrt(float64(per * (per - t.SideA) * (per - t.SideB) * (per - t.SideC)))))

}
func (t Triangle) perimeter() float64 {
	return float64((t.SideB + t.SideC + t.SideA) / 2.0)
}

type Circle struct {
	Radius float32
}

type Triangle struct {
	SideA float32
	SideB float32
	SideC float32
}

func PrintShapeDetails(s Shape) {

	fmt.Printf("Shape details:\n")
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())

}

func main() {
	shapes := []Shape{
		Rectangle{Width: 50, Height: 50},
		Circle{Radius: 5},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapes {
		PrintShapeDetails(shape)
	}
}
