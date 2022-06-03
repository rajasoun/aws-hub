package raja

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp() *os.File {
	logFile, err := os.OpenFile("e2e.md", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return logFile
}

func LogFlow(flow string) {
	log.Println(flow)
}

func Flow(sender string, reciever string, message string) string {
	flowDoc := fmt.Sprintf("%s ->> %s : %s ", sender, reciever, message)
	return flowDoc
}

var startDoc = "```mermaid \nsequenceDiagram\n	actor User"
var endDoc = "```"

func TestE2E(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	fileLog := setUp()
	log.SetOutput(fileLog)
	log.SetFlags(0)
	defer fileLog.Close()
	LogFlow(startDoc)
	flowDoc := Flow("User", "main", "aws-env go run main.go start")
	LogFlow("\t" + flowDoc)
	assert.Contains(flowDoc, "User", "")
	LogFlow(endDoc)
}
