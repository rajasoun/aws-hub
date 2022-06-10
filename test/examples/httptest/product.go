package spike

import (
	"encoding/json"
	"net/http"
)

var storeDB []Product

// Product represents a product in the store
type Product struct {
	Title          string  `json:"title,omitempty"`
	Price          float64 `json:"price,omitempty"`
	InventoryCount int     `json:"inventory_count,omitempty"`
}

// GetProducts returns all products in the store
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, v := range storeDB {
		if v.InventoryCount > 0 {
			products = append(products, v)
		}
	}
	json.NewEncoder(w).Encode(products)

}
