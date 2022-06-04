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
)

type FlowManager struct {
	fileName       string
	fileOptions    int
	filePermission os.FileMode
	startDoc       string
	endDoc         string
	FileOpener     func(string, int, os.FileMode) (*os.File, error)
}

//type FileOpener func(string, int, os.FileMode) (*os.File, error)

func NewFlowManager() *FlowManager {
	flowManager := FlowManager{
		fileName:       DefaultFileName,
		fileOptions:    DefaultFileOptions,
		filePermission: DefaultPermission,
		startDoc:       "```mermaid\n\nsequenceDiagram\n\tactor User",
		endDoc:         "\n```",
		FileOpener:     os.OpenFile,
	}
	return &flowManager
}

//func (flowmanager *FlowManager) CreateMarkdown(fileOpener FileOpener) (*os.File, error) {
func (fm *FlowManager) CreateMarkdown() (*os.File, error) {
	//logFile, err := fileOpener(flowmanager.fileName, flowmanager.fileOptions, flowmanager.filePermission)
	logFile, err := fm.FileOpener(fm.fileName, fm.fileOptions, fm.filePermission)
	if err != nil {
		log.Printf("Error Opening or Creating File %s Err = %v", fm.fileName, err)
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
