package ajith

import (
	"bytes"
	"log"
	"testing"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
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
				addopt:    "\t\topt\n",
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
				endopt:    "\n\t\tend",
			},
		},
		{
			name: "startCommand to GetCommand",
			flow: Flow{
				format:    "\t",
				sender:    "app.hub",
				direction: " ->> ",
				receiver:  "app.config.comd",
				message:   "GetCommand()",
			},
		},
		{
			name: "GetCommand to Create Command ",
			flow: Flow{
				format:    "\t",
				sender:    "app.config.cmd",
				direction: " ->> ",
				receiver:  "app.config.cmd",
				message:   "CreateCommand()",
			},
		},
		{
			name: "CreateCommand to urfave/cli/v2",
			flow: Flow{
				addopt:    "\topt\n",
				format:    "\t",
				sender:    "app.config.cmd",
				direction: "->>",
				receiver:  "urfave.cli",
				message:   "func(appCtx *cli.Context)",
			},
		},
		{
			name: "app.hub  to  app.hub SetupOutput",
			flow: Flow{
				format:    "\t",
				sender:    "app.hub",
				direction: "->>",
				receiver:  "app.hub",
				message:   "SetOutput()",
			},
		},
		/*	{
				name: "app.command to app",
				flow: Flow{
					format:    "\t",
					sender:    "app.hub",
					direction: "->>",
					receiver:  "app.config.arg",
					message:   "Urfavc.cli.run(args)",
				},

			},
		*/
		{
			name: "starting command using cli.run(args)",
			flow: Flow{
				format:    "\t",
				sender:    "app.cmd",
				direction: "->>",
				receiver:  "app.server",
				message:   "server.Start",
				endopt:    "\n\tend",
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

func TestCommandline(t *testing.T) {
	t.Run("flow simulation", func(t *testing.T) {
		assert := assert.New(t)
		var outputBuffer bytes.Buffer
		log.SetOutput(&outputBuffer)
		log.SetFlags(0)
		CliContext := cli.Context{}
		cmdhandler := cmd.CmdHandler{}
		cmdhandler.EnableShutdDown = true
		err := cmdhandler.StartCommand(&CliContext)
		got := outputBuffer.String()
		want := ":3000"
		assert.NoError(err, "err = %v", err)
		assert.Contains(got, want, "Server Start = %v, want =%v", got, want)

	})
}
func TestFlowOpenOrCreate(t *testing.T) {
	t.Run("Check Markdown Creation os.OpenFile", func(t *testing.T) {
		assert := assert.New(t)
		t.Parallel()

		//got, _ := flowManager.CreateMarkdown(os.OpenFile)
		got, _ := createMarkdown()
		want := "e2e.md"
		assert.Equal(want, got.Name(), "Flow.OpenOrCreate() = %v, want %v", got.Name(), want)
	})
}
