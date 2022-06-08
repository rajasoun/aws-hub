package ajith

import (
	"log"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	stepslog, _ := createMarkdown()
	defer stepslog.Close()

	Start(stepslog)
	steps := Flow{}
	t.Run("User To main", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "User"
		steps.direction = " ->> "
		steps.receiver = "main"
		steps.message = "aws-env go run main.go start"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	log.Println(optflow("\topt"))
	t.Run("main to app.hub", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "main"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "Execute()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)

	})

	t.Run("app.hub.execute() to app.hub.newApp()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "NewApp()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})

	t.Run("app.hub.newApp() to app.hub.setupinfo()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->>"
		steps.receiver = "app.hub"
		steps.message = "setupInfo()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	t.Run("app.hub.newApp() to app.hub.setupFlag()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "setupFlags()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	t.Run("app.hub.setupFlags() to hub.cli.Flags.GetFlags()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "hub.cli.Flags"
		steps.message = "Getflags()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)

	})
	t.Run("app.hub.newApp() to App.hub.setUPAuthors()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "setupAuthors"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	t.Run("app.hub.newApp() to App.hub.setUPcommands()", func(t *testing.T) {
		steps.format = "\t\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "setupACommmand()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	log.Println(optflow("\tend"))
	t.Run("app.hub.setupcommnads() to app.config.cmd.command.go", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.config.cmd"
		steps.message = "GetCommand()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	t.Run("app.config.cmd.GetCommands() to app.config.cmd.CreateCommand()", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "app.config.cmd"
		steps.direction = " ->> "
		steps.receiver = "app.config.cmd"
		steps.message = "CreateCommand()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)

	})
	t.Run("app.config.cmd.GetCommands() to app.config.cmd.CreateCommand()", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "app.config.cmd"
		steps.direction = " ->> "
		steps.receiver = "app.config.cmd"
		steps.message = "CreateCommand()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	t.Run("app.hub.setOutput() to app.hub.setOutput()  ", func(t *testing.T) {
		steps.format = "\t"
		steps.sender = "app.hub"
		steps.direction = " ->> "
		steps.receiver = "app.hub"
		steps.message = "app.setOutput()"
		docs := steps.GetMermaidFlow()
		log.Println(docs)
	})
	End()

}
