package raja

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()
	flow := Flow{}
	fileLog := flow.Start()

	t.Run("User To main", func(t *testing.T) {
		flowDoc := flow.GetFlow("User", "main", "aws-env go run main.go start")
		log.Println("\t" + flowDoc)
	})
	defer fileLog.Close()
}
