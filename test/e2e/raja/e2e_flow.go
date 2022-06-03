package raja

import (
	"fmt"
	"io"
	"log"
	"os"

	iface "github.com/rajasoun/aws-hub/test/e2e/raja/interface"
)

var startDoc = "```mermaid \nsequenceDiagram\n	actor User"
var endDoc = "```"

type Flow struct{}

func (flow *Flow) GetFlow(sender string, reciever string, message string) string {
	flowDoc := fmt.Sprintf("%s ->> %s : %s ", sender, reciever, message)
	return flowDoc
}

func (flow *Flow) Start(writer io.Writer) {
	log.SetOutput(writer)
	log.SetFlags(0)
	log.Println(startDoc)
}

func (*Flow) OpenOrCreate(openFile iface.OpenFile) (*os.File, error) {
	Options := os.O_RDWR | os.O_CREATE
	fileName := "e2e.md"
	logFile, err := openFile(fileName, Options, 0666)
	if err != nil {
		log.Printf("Error Opening File %s Err = %v", fileName, err)
		return nil, err
	}
	return logFile, nil
}

func (flow *Flow) End() {
	log.Println(endDoc)
}
