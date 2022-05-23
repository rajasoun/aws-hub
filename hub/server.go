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
	"github.com/robfig/cron"
	"github.com/rs/cors"
	"github.com/urfave/cli/v2"
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

func setUpRoutes(awsHandler *aws.AWSHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/aws/profiles", awsHandler.ConfigProfilesHandler)
	router.HandleFunc("/aws/iam/users", awsHandler.IAMUsersHandler)
	router.HandleFunc("/aws/iam/account", awsHandler.IAMUserHandler)
	router.HandleFunc("/aws/iam/organization", awsHandler.DescribeOrganizationHandler)
	router.HandleFunc("/aws/cost/current", awsHandler.CurrentCostHandler)
	router.HandleFunc("/aws/cost/history", awsHandler.CostAndUsageHandler)
	router.HandleFunc("/aws/cost/forecast", awsHandler.DescribeForecastPriceHandler)
	router.HandleFunc("/aws/cost/instance_type", awsHandler.CostAndUsagePerInstanceTypeHandler)
	router.HandleFunc("/health", awsHandler.HealthCheckHandler)
	return router
}

func setUpCors() *cors.Cors {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"profile", "X-Requested-With", "Content-Type", "Authorization"},
	})
	return corsOptions
}

func setUpServer(cache cache.Cache, multiple bool) http.Handler {
	awsHandler := setUpCache(cache, multiple)
	setUpCron()
	awsRoutes := setUpRoutes(awsHandler)
	//awsRoutes.PathPrefix("/").Handler(http.FileServer(assetFS()))
	corsOptions := setUpCors()
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsOptions.Handler(awsRoutes))
	return loggedRouter
}

func startServer(port int, loggedRouter http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Println("Error in Starting Application")
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}

func New() *cli.App {
	app := &cli.App{}
	setUpApp(app)
	setFlags(app)
	setUpCommands(app)
	return app
}

func Execute(args []string) {
	app := New()
	app.Run(args)
}
