package rohini

import (
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	flowLog, _ := CreateMarkdown()
	defer flowLog.Close()

	Start(flowLog)
	///flow := Flow{}

}
