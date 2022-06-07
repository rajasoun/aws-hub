package ajith

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

var startDoc string = "```mermaid\n\nsequenceDiagram\n\tactor User"
var endDoc string = "\n```"

func createMarkdown() (*os.File, error) {
	DefaultFileName := "e2e.md"
	DefaultFileOptions := os.O_RDWR | os.O_CREATE
	DefaultPermission := 0666

	logFile, err := os.OpenFile(DefaultFileName, DefaultFileOptions, fs.FileMode(DefaultPermission))
	if err != nil {
		log.Printf("Error Opening or Creating File %s Err = %v", DefaultFileName, err)
		return nil, err
	}

	return logFile, nil

}
func Start(writer io.Writer) {
	log.SetOutput(writer)
	log.SetFlags(0)
	log.Println(startDoc)
}

func End() {
	log.Println(endDoc)
}

type Flow struct {
	format    string
	sender    string
	receiver  string
	direction string
	message   string
}

func (flow *Flow) GetMermaidFlow() string {
	flowDoc := fmt.Sprintf("%s\t%s %s %s : %s ",
		flow.format, flow.sender, flow.direction, flow.receiver, flow.message)
	return flowDoc
}
