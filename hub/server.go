package hub

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rajasoun/aws-hub/handlers/aws"
	"github.com/rajasoun/aws-hub/services/cache"
	"github.com/rs/cors"
)

type Server struct {
	name        string
	httpHandler http.Handler
	awsHandler  *aws.AWSHandler
	routes      *mux.Router
	cors        *cors.Cors
}

var awsHandler *aws.AWSHandler

func setUpCache(cache cache.Cache, multiple bool) *aws.AWSHandler {
	cache.Connect()
	awsHandler = aws.NewAWSHandler(cache, multiple)
	return awsHandler
}

func setUpCors() *cors.Cors {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	return corsOptions
}

func NewServer(cache cache.Cache, multiple bool) *Server {
	server := Server{}
	server.name = "Mux Server 0.1"
	server.awsHandler = setUpCache(cache, multiple)
	server.routes = awsHandler.SetUpRoutes()
	server.cors = setUpCors()
	server.httpHandler = handlers.LoggingHandler(os.Stdout, server.cors.Handler(server.routes))
	return &server
}

func (server *Server) start(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), server.httpHandler)
	if err != nil {
		log.Println("Error in Starting Application")
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}
