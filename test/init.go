package test

import (
	"flag"
	"fmt"
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
