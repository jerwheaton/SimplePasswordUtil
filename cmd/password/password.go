package main

import (
	"os"

	"github.com/therounds/PasswordService/pkg/cmd"
)

func main() {
	command := cmd.NewCommand()
	command.Execute()

	os.Exit(0)
}
