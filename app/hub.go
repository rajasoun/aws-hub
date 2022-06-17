// Package app is the cli based application for Analysis and management of AWS cost
package app

import (
	"io"
	"log"
	"time"

	"github.com/rajasoun/aws-hub/app/config/cmd"
	"github.com/rajasoun/aws-hub/app/config/flag"
	structs "github.com/rajasoun/go-ds"
	"github.com/urfave/cli/v2"
)

// Hub is CLI Application.
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

func (hub *Hub) setUpCommands() {
	handler := cmd.Handler{EnableShutdDown: false}
	startCommand := cmd.New("start", "Start Server", handler.StartCommand)
	commands := []*cli.Command{&startCommand}
	hub.cli.Commands = commands
	hub.cli.CommandNotFound = cmd.NewErr()
}

func (hub *Hub) setUpOutput(writer io.Writer) {
	log.SetOutput(writer)
	hub.cli.Writer = writer
}

// StructToMap to convert struct into a map[string]interface{}.
func (hub *Hub) StructToMap(ds interface{}) map[string]interface{} {
	s := structs.New(ds)
	m := s.Map()
	return m
}

// SliceToStrMap to convert slice of cli.Flag into a map[int]string{}.
func (hub *Hub) SliceToStrMap(elements []cli.Flag) map[int]string {
	elementMap := make(map[int]string)
	for index, s := range elements {
		elementMap[index] = s.String()
	}
	return elementMap
}

// New creates and returns cli hub application.
func New() *Hub {
	hub := Hub{&cli.App{}}
	hub.setUpInfo()
	hub.setUpAuthors()
	hub.setUpFlags()
	hub.setUpCommands()
	return &hub
}

func Execute(args []string, writer io.Writer) error {
	app := New()
	// Dependency Injection - Enable TDD
	app.setUpOutput(writer)
	// Start App
	err := app.cli.Run(args)
	return err
}
