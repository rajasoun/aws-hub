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
	}
	Start(stepslog)
	for _, tt := range steps {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.flow.GetMermaidFlow())
		})
	}
	End()
}
