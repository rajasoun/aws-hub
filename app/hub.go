package app

import (
	"io"
	"time"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/rajasoun/aws-hub/app/config/flag"
	structs "github.com/rajasoun/go-ds"
	"github.com/urfave/cli/v2"
)

type Hub struct {
	cli *cli.App
}

func (hub *Hub) setUpInfo() {
	hub.cli.Name = "AWS Hub"
	hub.cli.Usage = "AWS Cost Explorer"
	hub.cli.Version = "0.0.1"
	hub.cli.Compiled = time.Now()
}

func (hub *Hub) setUpAuthors() {
	authors := []*cli.Author{
		{
			Name:  "Raja Soundaramourty",
			Email: "rajasoun@cisco.com",
		},
	}
	hub.cli.Authors = authors
}

func (hub *Hub) setUpFlags() {
	hub.cli.Flags = flag.GetFlags()
}

func (hub *Hub) setUpCommands(handler func(appCtx *cli.Context) error) {
	hub.cli.Commands = cmd.GetCommands(handler)
}

func (hub *Hub) StructToMap(ds interface{}) map[string]interface{} {
	s := structs.New(ds)
	m := s.Map()
	return m
}

func (hub *Hub) SliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}

func NewApp() *Hub {
	hub := Hub{&cli.App{}}
	hub.setUpInfo()
	hub.setUpAuthors()
	hub.setUpFlags()
	hub.setUpCommands(cmd.StartCommandHandler)
	return &hub
}

func Execute(args []string, writer io.Writer) error {
	app := NewApp()
	app.cli.Writer = writer
	err := app.cli.Run(args)
	return err
}
