package spike

import (
	"encoding/json"
	"fmt"
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
func GetProductsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var products []Product
	for _, v := range storeDB {
		if v.InventoryCount > 0 {
			products = append(products, v)
		}
	}
	if len(products) > 0 {
		RespondWithJSON(responseWriter, products, http.StatusOK)
	} else { // Handle Empty Result
		errMsg := "Store is Empty"
		RespondWithErrorJSON(responseWriter, errMsg)
	}
}

// GetProduct returns the product with the specified title
func GetProductHandler(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["title"]
	for _, product := range storeDB {
		if product.Name == name {
			RespondWithJSON(responseWriter, product, http.StatusOK)
			return
		}
	}
	//if no product is found
	errMsg := fmt.Sprintf("Not Found Product for Name = %v", name)
	RespondWithErrorJSON(responseWriter, errMsg)
}

func RespondWithJSON(responseWriter http.ResponseWriter, payload interface{}, code int) {
	jsonPayLoad, _ := json.Marshal(payload)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(jsonPayLoad)
}

func RespondWithErrorJSON(responseWriter http.ResponseWriter, errMsg string) {
	code := http.StatusBadRequest
	payload := map[string]string{"error": errMsg}
	RespondWithJSON(responseWriter, payload, code)
}
