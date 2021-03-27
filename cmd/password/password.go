package main

import (
	"os"

	"github.com/jerwheaton/SimplePasswordUtil/pkg/cmd"
)

func main() {
	command := cmd.NewCommand()
	command.Execute()

	os.Exit(0)
}
