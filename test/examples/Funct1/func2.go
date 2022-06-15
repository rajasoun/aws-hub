package function

import (
	"fmt"
)

func Bill(price, number int) int {
	var totPrice = price * number
	return totPrice
}

func main() {
	price, number := 90, 6
	totPrice := Bill(price, number)
	fmt.Println("Total price is", totPrice)
}
