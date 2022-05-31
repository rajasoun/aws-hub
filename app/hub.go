package app

import (
	"io"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/rajasoun/aws-hub/app/config/flag"
	structs "github.com/rajasoun/go-ds"
	"github.com/urfave/cli/v2"
)

func NewApp() *App {
	app := App{&cli.App{}}
	app.SetUpInfo()
	app.cli.Flags = flag.GetFlags()
	app.cli.Commands = cmd.GetCommands(cmd.StartCommandHandler)
	return &app
}

func Execute(args []string, writer io.Writer) error {
	app := NewApp()
	app.cli.Writer = writer
	err := app.cli.Run(args)
	return err
}

func (app *App) StructToMap() map[string]interface{} {
	s := structs.New(app.cli)
	m := s.Map()
	return m
}

func SliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}
