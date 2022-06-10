package spike

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var storeDB []Product

// Product represents a product in the store
type Product struct {
	Name           string  `json:"title,omitempty"`
	Price          float64 `json:"price,omitempty"`
	InventoryCount int     `json:"inventory_count,omitempty"`
}

// GetProducts returns all products in the store
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, v := range storeDB {
		if v.InventoryCount > 0 {
			products = append(products, v)
		}
	}
	json.NewEncoder(w).Encode(products)

}

// GetProduct returns the product with the specified title
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["title"]
	for _, v := range storeDB {
		if v.Name == name {
			log.Printf("SAME -> v.Title = %v, title = %v", v.Name, name)
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	//if no product is found
	log.Printf("Not Found Product for Name = %v", name)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Not Found")
}
