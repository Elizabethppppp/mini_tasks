package _interface

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }

type Rectangle struct{ W, H float64 }

type Triangle struct{ A, H float64 }

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (t Triangle) Area() float64 {
	return 0.5 * t.A * t.H
}

func DescribeShape(s Shape) {
	fmt.Println(s.Area())
}

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func Largest(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}
	largest := shapes[0]
	for _, shape := range shapes {
		if shape.Area() > largest.Area() {
			largest = shape
		}
	}
	return largest
}
