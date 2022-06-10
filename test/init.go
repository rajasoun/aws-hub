package test

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Initialize go test
func init() {
	testing.Init()
	flag.Bool("isTest", true, "Returns true if run from go test")
	flag.Parse()
}

// Returns true if invoked with go test -v or go test
func IsTestRun() bool {
	fmt.Println()
	verbose := flag.Lookup("test.v").Value.(flag.Getter).Get().(bool)
	isTest := flag.Lookup("isTest").Value.(flag.Getter).Get().(bool)
	return verbose || isTest
}

func ExecuteHandler(handlerName func(w http.ResponseWriter, r *http.Request),
	muxRequestVars map[string]string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", "/test", nil)
	request = mux.SetURLVars(request, muxRequestVars)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerName)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
