package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	aws "github.com/rajasoun/aws-hub/handlers"
	"github.com/rajasoun/aws-hub/service/cache"
	"github.com/rs/cors"
)

type Server struct {
	name             string
	httpHandler      http.Handler
	awsHandler       *aws.AWSHandler
	routes           *mux.Router
	cors             *cors.Cors
	shutdownDuration time.Duration
}

func setUpCache(cacheHandler cache.Cache, multiple bool) *aws.AWSHandler {
	cacheHandler.Connect()
	return aws.NewAWSHandler(cacheHandler, multiple)
}

func setUpCors() *cors.Cors {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	return corsOptions
}

func NewServer(cacheHandler cache.Cache, multiple bool) (*Server, *mux.Router) {
	server := Server{}
	server.name = "Mux Server 0.1"
	server.awsHandler = setUpCache(cacheHandler, multiple)
	// Connects Routes to Handlers
	router := server.awsHandler.SetUpRoutes()
	server.routes = router
	server.cors = setUpCors()
	server.httpHandler = handlers.LoggingHandler(os.Stdout, server.cors.Handler(server.routes))
	return &server, router
}

func (server *Server) GetAWSHandler() *aws.AWSHandler {
	return server.awsHandler
}

func (server *Server) Start(port int, enableShutdown bool) error {
	portString := ":" + strconv.Itoa(port)
	httpServer := server.NewHTTPServer(portString)
	if enableShutdown {
		server.RegisterShutdown(httpServer)
	}
	err := httpServer.StartHTTPServer()
	return err
}

func (server *Server) RegisterShutdown(httpServer HTTPServer) {
	go func() {
		duration := server.shutdownDuration
		time.Sleep(duration * time.Second)
		defer httpServer.Shutdown(context.Background())
	}()
}

type HTTPServer struct {
	*http.Server
}

func (server *Server) NewHTTPServer(adr string) HTTPServer {
	log.Println("adr ", adr)
	return HTTPServer{
		&http.Server{
			Addr:    adr,
			Handler: server.httpHandler,
		},
	}
}

func (httpServer HTTPServer) StartHTTPServer() error {
	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
