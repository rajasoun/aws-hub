package test

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"testing"
)

// Initialize go test
func init() {
	testing.Init()
	flag.Bool("isTest", true, "Returns true if run from go test")
	flag.Parse()
}

// Returns true if invoked with go test -v or go test
func IsTestRun() bool {
	fmt.Println()
	verbose := flag.Lookup("test.v").Value.(flag.Getter).Get().(bool)
	isTest := flag.Lookup("isTest").Value.(flag.Getter).Get().(bool)
	return verbose || isTest
}

func SetLogOutputToBuffer() *bytes.Buffer {
	var outputBuffer bytes.Buffer
	log.SetOutput(&outputBuffer)
	return &outputBuffer
}
