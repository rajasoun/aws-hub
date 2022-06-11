package raja

import (
	"bytes"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli/v2"
)

// Mocking References
// 1. https://medium.com/swlh/golangs-interfaces-explained-with-mocks-886f69eca6f0
// 2. https://www.myhatchpad.com/insight/mocking-techniques-for-go/
func TestE2E(t *testing.T) {
	t.Parallel()
	var flowByCoding FlowAutomation = NewFlowManager()
	//flowLog, _ := flowManager.CreateMarkdown(os.OpenFile)
	flowLog, _ := flowByCoding.CreateMarkdown()
	defer flowLog.Close()
	flows := []struct {
		name string
		flow Flow
	}{
		{
			name: "User To main",
			flow: Flow{
				format:    "\t",
				sender:    "User",
				direction: " ->> ",
				receiver:  "main",
				message:   "aws-env go run main.go start",
			},
		},
		{
			name: "main to app.hub",
			flow: Flow{
				format:    "\t",
				sender:    "main",
				direction: " ->> ",
				receiver:  "app.hub",
				message:   "Execute()",
			},
		},
		{
			name: "urfave.cli to start command",
			flow: Flow{
				format:    "\t",
				sender:    "urfave.cli",
				direction: " ->> ",
				receiver:  "app.config.cmd.startCmd",
				message:   "StartCommand(appCtx *cli.Context)",
			},
		},
		{
			name: "start command to server start",
			flow: Flow{
				format:    "\t",
				sender:    "app.config.cmd.startCmd",
				direction: " ->> ",
				receiver:  "app.server",
				message:   "Start(port, enableShutdown)",
			},
		},
	}
	flowByCoding.Start(flowLog)
	for _, tt := range flows {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.flow.GetMermaidFlow())
		})
	}
	log.Println("\t\tNote right of app.server: Server Started!")
	flowByCoding.End()
}

func TestSimulateExecute(t *testing.T) {
	t.Run("Flow Execute Simulation", func(t *testing.T) {
		assert := assert.New(t)
		var outputBuffer bytes.Buffer
		log.SetOutput(&outputBuffer)
		log.SetFlags(0)

		set := flag.NewFlagSet("test", 0)
		port, _ := test.GetFreePort("localhost:0")
		portString := strconv.Itoa(port)
		_ = set.Parse([]string{"start", "--port", portString})

		mockApp := &cli.App{Writer: ioutil.Discard}
		context := cli.NewContext(mockApp, set, nil)
		cmdhandler := cmd.CmdHandler{}
		cmdhandler.EnableShutdDown = true
		startCommand := cmd.GetCommand(cmdhandler.StartCommand)
		err := startCommand.Run(context)
		assert.NoError(err, "err = %v ", err)

		got := outputBuffer.String()
		want := portString
		assert.Contains(got, want, "Server Start = %v, want = %v", got, want)
	})
}

func TestFlowOpenOrCreate(t *testing.T) {
	t.Run("Check Markdown Creation os.OpenFile", func(t *testing.T) {
		assert := assert.New(t)
		t.Parallel()
		flowManager := NewFlowManager()
		//got, _ := flowManager.CreateMarkdown(os.OpenFile)
		got, _ := flowManager.CreateMarkdown()
		want := flowManager.markdown.fileName
		assert.Equal(want, got.Name(), "Flow.OpenOrCreate() = %v, want %v", got.Name(), want)
	})
}

type MockOs struct {
	mock.Mock
}

func (c *MockOs) FileOpener(name string, flag int, perm os.FileMode) (*os.File, error) {
	args := c.Called(name, flag, perm)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*os.File), args.Error(1)
}

func TestFlowOpenOrCreateErr(t *testing.T) {
	t.Run("Check Markdown Creation for Err", func(t *testing.T) {
		assert := assert.New(t)
		mockos := new(MockOs)
		flowManager := NewFlowManager()
		mockErr := errors.New("simulated error")
		mockos.
			On("FileOpener", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, mockErr)
		// Inject Mock FileOpener Function
		//_, err := flowManager.CreateMarkdown(mockos.FileOpener)
		flowManager.markdown.FileOpener = mockos.FileOpener
		_, err := flowManager.CreateMarkdown()
		assert.Error(err, "Flow.OpenOrCreate() Err = %v", err)
		assert.Equal(err, mockErr, " err = %v , mockErr = %v ", err, mockErr)
	})
}
