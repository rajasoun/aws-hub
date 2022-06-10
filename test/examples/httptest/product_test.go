package spike

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpStoreDB() {
	paste := Product{Title: "Paste", Price: 10, InventoryCount: 10}
	soap := Product{Title: "Soap", Price: 20, InventoryCount: 5}
	storeDB = append(storeDB, paste)
	storeDB = append(storeDB, soap)
}

func TestGetProductsHandler(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	setUpStoreDB()
	t.Run("Check Get Products", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/test", nil)
		assert.NoError(err, "http.NewRequest() = %v ", err)
		responseRecorder := httptest.NewRecorder()

		// Set Handler
		handler := http.HandlerFunc(GetProducts)
		// ServeHTTP calls
		handler.ServeHTTP(responseRecorder, request)

		got := responseRecorder.Code
		want := http.StatusOK
		assert.Equal(got, want, "handler returned wrong status code: got %v want %v", got, want)
	})

}
