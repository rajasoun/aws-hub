package main

import (
	"log"
	"os"

	"github.com/rajasoun/aws-hub/hub"
)

func main() {
	err := hub.Execute(os.Args, os.Stdout)
	if err != nil {
		log.Println("Error in Starting Application")
		log.Fatal(err)
	}
}
