package main

import (
	"os"

	"github.com/rajasoun/aws-hub/hub"
)

func main() {
	hub.Execute(os.Args, os.Stdout)
}
