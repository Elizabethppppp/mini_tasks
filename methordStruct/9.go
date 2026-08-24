package methordStruct

import "math"

type Point struct {
	X, Y int
}

func (p Point) Dist() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func Dist2(p Point) float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
