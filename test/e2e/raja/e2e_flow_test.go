package raja

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestE2E(t *testing.T) {
	t.Parallel()
	flow := Flow{}
	fileLog, _ := flow.OpenOrCreate(os.OpenFile)
	flow.Start(fileLog)
	defer fileLog.Close()

	t.Run("User To main", func(t *testing.T) {
		flowDoc := flow.GetFlow("User", "main", "aws-env go run main.go start")
		log.Println("\t" + flowDoc)
	})
	flow.End()
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

type MockIt struct {
	mock.Mock
}

func (c *MockIt) openFilefunc(name string, flag int, perm os.FileMode) (*os.File, error) {
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
		mockIt := new(MockIt)
		mockIt.
			On("openFilefunc", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("simulated error"))
		_, err := flow.OpenOrCreate(mockIt.openFilefunc)
		assert.Error(err, "Flow.OpenOrCreate() Err = %v", err)
	})
}
