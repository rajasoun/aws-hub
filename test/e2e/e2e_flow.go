package raja

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	DefaultFileName                = "e2e.md"
	DefaultFileOptions int         = os.O_RDWR | os.O_CREATE
	DefaultPermission  os.FileMode = 0666
	DefaultFileOpener              = os.OpenFile
)

type Markdown struct {
	fileName       string
	fileOptions    int
	filePermission os.FileMode
	FileOpener     func(string, int, os.FileMode) (*os.File, error)
}

type FlowManager struct {
	startDoc string
	endDoc   string
	markdown Markdown
}

type FlowAutomation interface {
	CreateMarkdown() (*os.File, error)
	Start(writer io.Writer)
	End()
}

func NewFlowManager() *FlowManager {
	markdown := Markdown{
		fileName:       DefaultFileName,
		fileOptions:    DefaultFileOptions,
		filePermission: DefaultPermission,
		FileOpener:     DefaultFileOpener,
	}

	flowManager := FlowManager{
		startDoc: "```mermaid\n\nsequenceDiagram\n\tactor User",
		endDoc:   "\n```",
		markdown: markdown,
	}
	return &flowManager
}

func (fm *FlowManager) CreateMarkdown() (*os.File, error) {
	mk := fm.markdown
	logFile, err := mk.FileOpener(mk.fileName, mk.fileOptions, mk.filePermission)
	if err != nil {
		log.Printf("Error Opening or Creating File %s Err = %v", mk.fileName, err)
		return nil, err
	}
	return logFile, nil
}

func (fm *FlowManager) Start(writer io.Writer) {
	log.SetOutput(writer)
	log.SetFlags(0)
	log.Println(fm.startDoc)
}

func (fm *FlowManager) End() {
	log.Println(fm.endDoc)
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
