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

// dimensions of a square.
type Square struct {
	Height float64
	Width  float64
}

// area of the square.
func (r Square) Area() float64 {
	return r.Height * r.Width
}
