package raja

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli/v2"
)

func TestE2E(t *testing.T) {
	t.Parallel()
	flow := Flow{}
	fileLog, _ := flow.OpenOrCreate(os.OpenFile)
	flow.Start(fileLog)
	defer fileLog.Close()

	t.Run("User To main", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "User"
		flow.direction = " ->> "
		flow.receiver = "main"
		flow.message = "aws-env go run main.go start"
		flowDoc := flow.GetFlow()
		log.Println(flowDoc)
	})
	t.Run("main to app.hub", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "main"
		flow.direction = " ->> "
		flow.receiver = "app.hub"
		flow.message = "Execute()"
		flowDoc := flow.GetFlow()
		log.Println(flowDoc)
	})
	t.Run("main to app.hub", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "app.hub"
		flow.direction = " ->> "
		flow.receiver = "urfave.cli"
		flow.message = "app.cli.Run(args)"
		flowDoc := flow.GetFlow()
		log.Println(flowDoc)
	})
	t.Run("urfave.cli to start command", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "urfave.cli"
		flow.direction = " ->> "
		flow.receiver = "app.config.cmd.startCmd"
		flow.message = "StartCommand(appCtx *cli.Context)"
		flowDoc := flow.GetFlow()
		log.Println(flowDoc)
	})
	t.Run("start command to server start", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "app.config.cmd.startCmd"
		flow.direction = " ->> "
		flow.receiver = "app.server"
		flow.message = "Start(port, enableShutdown)"
		flowDoc := flow.GetFlow()
		log.Println(flowDoc)
		log.Println("\t\tNote right of app.server: Server Started!")
	})
	flow.End()
}

func TestSimulateExecute(t *testing.T) {
	t.Run("Flow Execute Simulation", func(t *testing.T) {
		assert := assert.New(t)
		var outputBuffer bytes.Buffer
		log.SetOutput(&outputBuffer)
		log.SetFlags(0)
		//Simulate app.cli.Run(args) for command start
		defaultCliContext := cli.Context{}
		cmdhandler := cmd.CmdHandler{}
		cmdhandler.EnableShutdDown = true
		err := cmdhandler.StartCommand(&defaultCliContext)
		got := outputBuffer.String()
		want := ":3000"
		assert.NoError(err, "err = %v ", err)
		assert.Contains(got, want, "Server Start = %v, want = %v", got, want)
		// var outputBuffer bytes.Buffer
		// log.SetOutput(&outputBuffer)
		// log.SetFlags(0)
		// args := os.Args[0:1]
		// args = append(args, "start")
		// err := app.Execute(args, &outputBuffer)
		// assert.NoError(err, "err = %v ", err)
	})
}

func TestFlowOpenOrCreate(t *testing.T) {
	t.Run("Check File Open Create", func(t *testing.T) {
		assert := assert.New(t)
		t.Parallel()
		flow := &Flow{}
		got, _ := flow.OpenOrCreate(os.OpenFile)
		want := "e2e.md"
		assert.Equal(want, got.Name(), "Flow.OpenOrCreate() = %v, want %v", got.Name(), want)
	})
}

type MockOs struct {
	mock.Mock
}

func (c *MockOs) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	args := c.Called(name, flag, perm)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*os.File), args.Error(1)
}

func TestFlowOpenOrCreateErr(t *testing.T) {
	t.Run("Check File Open Create For Error with Framework Mock", func(t *testing.T) {
		assert := assert.New(t)
		flow := &Flow{}
		mockos := new(MockOs)
		mockos.
			On("OpenFile", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("simulated error"))
		_, err := flow.OpenOrCreate(mockos.OpenFile)
		assert.Error(err, "Flow.OpenOrCreate() Err = %v", err)
	})
}
