package raja

import (
	"fmt"
	"log"
	"os"
)

var startDoc = "```mermaid \nsequenceDiagram\n	actor User"
var endDoc = "```"

type Flow struct{}

func (flow *Flow) GetFlow(sender string, reciever string, message string) string {
	flowDoc := fmt.Sprintf("%s ->> %s : %s ", sender, reciever, message)
	return flowDoc
}

func (flow *Flow) Start() *os.File {
	logFile, err := os.OpenFile("e2e.md", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(0)
	log.Println(startDoc)
	return logFile
}

func (flow *Flow) End() {
	log.Println(endDoc)
}
