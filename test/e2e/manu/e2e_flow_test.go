package manu

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()
	var flowByCoding FlowAutomation = NewFlowManager()

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

/*
func TestE2E(t *testing.T) {

	got := Createfile(file)
	want := "flow.md"

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
*/
