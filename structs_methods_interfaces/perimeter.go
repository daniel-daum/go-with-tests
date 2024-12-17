package structsmethodsinterfaces

import "math"

func Perimeter(length, width float64) (perimeter float64) {
	return 2 * (length + width)
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

type Shape interface {
	Area() float64
}
