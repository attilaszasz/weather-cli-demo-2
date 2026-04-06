package main

import (
	"os"

	"weather-cli/src/internal/exitcode"
)

func main() {
	os.Exit(execute())
}

func execute() int {
	command := newRootCommand()
	if err := command.Execute(); err != nil {
		return exitcode.FromError(err)
	}

	return exitcode.Success
}
