package main

import (
	"gtools/registry"
	"os"
)


func main() {
	if len(os.Args) == 1 {
		registry.Help(nil)
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "-v", "--version":
		cmd = "version"
	case "-h", "--help":
		cmd = "help"
	}

	if cmd == "help" {
		registry.Help(args)
		return
	}

	command, ok := registry.CmdMap[cmd]
	if !ok {
		registry.Help(nil)
		return	
	}
	command.Run(args)
}
