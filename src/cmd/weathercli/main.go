package main

import (
	"os"
)

func main() {
	command := newRootCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
