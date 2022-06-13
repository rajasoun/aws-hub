package spike

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
