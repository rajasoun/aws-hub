package ajith

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	stepslog, _ := createMarkdown()
	defer stepslog.Close()

	steps := []struct {
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
			name: "main to Ececute",
			flow: Flow{
				format:    "\t",
				sender:    "main",
				direction: " ->> ",
				receiver:  "app.hub",
				message:   "Execute()",
			},
		},
		{
			name: "Execute to newApp",
			flow: Flow{
				addopt:    "\topt\n",
				format:    "\t\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "app.hub",
				message:   "NewApp()",
			},
		},
		{
			name: "NewApp to setupInfo()",
			flow: Flow{
				format:    "\t\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "app.hub",
				message:   "setupInfo()",
			},
		},
		{
			name: "NewApp to setupAuthor",
			flow: Flow{
				format:    "\t\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "app.hub",
				message:   "SetupAuthor",
			},
		},
		{
			name: "NewApp() to Setup Command",
			flow: Flow{
				format:    "\t\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "aap.hub",
				message:   "setupCommand()",
				endopt:    "\n\tend",
			},
		},
		{
			name: "startCommand to GetCommand",
			flow: Flow{
				format:    "\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "app.config.comd.command.go",
				message:   "GetCommand",
			},
		},
	}
	Start(stepslog)
	for _, tt := range steps {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.flow.GetMermaidFlow())
		})
	}
	End()
}
