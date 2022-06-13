package test

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
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

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort(address string) (int, error) {
	// "localhost:0"
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return 0, err
	}
	port, err := CheckAddressAvailable(net.ListenTCP, addr)
	return port, err
}

func CheckAddressAvailable(tcpHandler func(network string, laddr *net.TCPAddr) (*net.TCPListener, error),
	addr *net.TCPAddr) (int, error) {
	l, err := tcpHandler("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

type MockServer struct {
}

func (mock *MockServer) DoSimulation(handlerName func(w http.ResponseWriter, r *http.Request),
	muxRequestVars map[string]string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", "/test", http.NoBody)
	request = mux.SetURLVars(request, muxRequestVars)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerName)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func MockSuccessHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "text/json")
	responseWriter.WriteHeader(http.StatusOK)
	payLoad := `{"Message":"test simulation"}`
	_, err := io.WriteString(responseWriter, payLoad)
	handleErr(err, "MockSuccessHandler()")
}

func MockFailureHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "text/json")
	responseWriter.WriteHeader(http.StatusInternalServerError)
	payLoad := `{"Message":"simulated error"}`
	_, err := io.WriteString(responseWriter, payLoad)
	handleErr(err, "MockFailureHandler()")
}

func handleErr(err error, errMsg string) {
	if err != nil {
		log.Printf(errMsg+"Err = %v", err)
	}
}
