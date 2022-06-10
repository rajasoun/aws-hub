package rohini

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	flowLog, _ := CreateMarkdown()
	defer flowLog.Close()

	Start(flowLog)
	flow := Flow{}
	t.Run("User To main", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "User"
		flow.direction = " ->> "
		flow.receiver = "main"
		flow.message = "aws-env go run main.go start"
		flowDoc := flow.GetMermaidFlow()
		log.Println(flowDoc)
	})

	t.Run("main to app.hub", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "main"
		flow.direction = " ->> "
		flow.receiver = "app.hub"
		flow.message = "Execute()"
		flowDoc := flow.GetMermaidFlow()
		log.Println(flowDoc)

	})

	t.Run("main to app.hub", func(t *testing.T) {
		flow.format = "\t"
		flow.sender = "app.hub"
		flow.direction = " ->> "
		flow.receiver = "urfave.cli"
		flow.message = "app.cli.Run(args)"
		flowDoc := flow.GetMermaidFlow()
		log.Println(flowDoc)
	})

	End()
}
