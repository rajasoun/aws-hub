package main

import (
	"log"
	"os"

	"github.com/rajasoun/aws-hub/app"
)

func main() {
	err := app.Execute(os.Args, os.Stdout)
	if err != nil {
		log.Println("Error in Starting Application")
		log.Fatal(err)
	}
}
