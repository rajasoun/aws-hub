package raja

import (
	"fmt"
	"io"
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

func (flow *Flow) Start(writer io.Writer) {
	log.SetOutput(writer)
	log.SetFlags(0)
	log.Println(startDoc)
}

func (*Flow) OpenOrCreate(osOpenFile func(name string, flag int, perm os.FileMode) (*os.File, error)) (*os.File, error) {
	Options := os.O_RDWR | os.O_CREATE
	FileName := "e2e.md"
	logFile, err := osOpenFile(FileName, Options, 0666)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func (flow *Flow) End() {
	log.Println(endDoc)
}
