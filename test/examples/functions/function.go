package function

import (
	"fmt"
)

func DmartAdd(Soaps1, Cloths int) (result int) {
	return Soaps1 + Cloths
}
func DmartDelete(Soaps, Cloths int) (result int) {
	return Soaps - Cloths

}
func Dmart() {
	fmt.Println("DmartAdd(1, 2)")
}

//on 15th june I added an interface and a function to the existing code
type Shapes interface {
	Area() float64
}

// dimensions of a square.
type Square struct {
	Height float64
	Width  float64
}

// area of the square.
func (r Square) Area() float64 {
	return r.Height * r.Width
}

// dimension of a triangle.
type Triangle struct {
	Height float64
	Base   float64
}

// area of the triangle.
func (c Triangle) Area() float64 {
	return (c.Height * c.Base) * 0.5
}
