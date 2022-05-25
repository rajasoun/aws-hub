package hub

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/rajasoun/aws-hub/handlers/aws"
	"github.com/rajasoun/aws-hub/services/cache"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

var awsHandler *aws.AWSHandler

func setUpCache(cache cache.Cache, multiple bool) *aws.AWSHandler {
	cache.Connect()
	awsHandler = aws.NewAWSHandler(cache, multiple)
	return awsHandler
}

func setUpCron() {
	c := cron.New()
	c.Start()
}

func setUpCors() *cors.Cors {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	return corsOptions
}

func setUp(cache cache.Cache, multiple bool) http.Handler {
	awsHandler := setUpCache(cache, multiple)
	setUpCron()
	awsRoutes := awsHandler.SetUpRoutes()
	//awsRoutes.PathPrefix("/").Handler(http.FileServer(assetFS()))
	corsOptions := setUpCors()
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsOptions.Handler(awsRoutes))
	return loggedRouter
}

func start(port int, loggedRouter http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Println("Error in Starting Application")
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}
