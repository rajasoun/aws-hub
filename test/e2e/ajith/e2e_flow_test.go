package ajith

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	flowLog, _ := createMarkdown()
	defer flowLog.Close()

	Start(flowLog)
	steps := Flow{}
	t.Run("User To main", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "User"
		steps.direction = " ->> "
		steps.receiver = "main"
		steps.message = "aws-env go run main.go start"
		flowDoc := steps.GetMermaidFlow()
		log.Println(flowDoc)
	})
	t.Run("main to app.hub", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "main"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "Execute()"
		flowDoc := steps.GetMermaidFlow()
		log.Println(flowDoc)

	})
	t.Run("app.hub.execute() to app.hub.newApp()", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "NewApp()"
		flowDoc := steps.GetMermaidFlow()
		log.Println(flowDoc)
	})
	End()

}
