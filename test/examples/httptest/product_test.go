package spike

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setUpStoreDB() {
	paste := Product{Name: "Paste", Price: 10, InventoryCount: 10}
	soap := Product{Name: "Soap", Price: 20, InventoryCount: 5}
	storeDB = append(storeDB, paste)
	storeDB = append(storeDB, soap)
}

func TestGetProducts(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	setUpStoreDB()
	t.Run("Check Get Products", func(t *testing.T) {
		responseRecorder := executeHandler(GetProductsHandler, map[string]string{})
		got := responseRecorder.Code
		want := http.StatusOK
		assert.Equal(got, want, "handler returned wrong status code: got %v want %v", got, want)
	})
}

func TestGetProduct(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	setUpStoreDB()
	tests := []struct {
		name    string
		want    int
		muxVars map[string]string
	}{
		{
			name:    "GetProduct: valid test case",
			muxVars: map[string]string{"title": "Soap"},
			want:    http.StatusOK,
		},
		{
			name:    "GetProduct: invalid test case",
			muxVars: map[string]string{"title": "invalid_value"},
			want:    http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run("Check Get Products", func(t *testing.T) {
			responseRecorder := executeHandler(GetProductHandler, tt.muxVars)
			got := responseRecorder.Code
			assert.Equal(tt.want, got, "handler returned wrong status code: got %v want %v", got, tt.want)
		})
	}
}

func executeHandler(handlerName func(w http.ResponseWriter, r *http.Request),
	muxRequestVars map[string]string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", "/test", nil)
	request = mux.SetURLVars(request, muxRequestVars)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerName)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
