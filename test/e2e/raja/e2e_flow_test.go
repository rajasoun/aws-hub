package raja

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Run("Check File Open Create For Error", func(t *testing.T) {
		assert := assert.New(t)
		flow := &Flow{}
		osOpenFile := func(name string, flag int, perm os.FileMode) (*os.File, error) {
			err := errors.New("Simulated Err")
			return nil, err
		}
		_, err := flow.OpenOrCreate(osOpenFile)
		assert.Error(err, "Flow.OpenOrCreate() Err = %v", err)
	})
}
