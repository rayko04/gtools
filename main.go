package main

import (
	"fmt"
	"gtools/commands"
	"os"
)

type Command struct {
	Description		string
	Run				func([]string)
}

var cmdMap = map[string]Command {
	"pwd": {
		Description: "Print working directory",
		Run: commands.Pwd,
	},
	"ls": {
		Description: "List directory content",
		Run: commands.Ls,
	},
	"cat": {
		Description: "Print file content",
		Run: commands.Cat,
	},
	"touch": {
		Description: "Create a file if it does not exist",
		Run: commands.Touch,
	},
	"mkdir": {
		Description: "Create directory",
		Run: commands.Mkdir,
	},
	"rm": {
		Description: "Remove a file or an empty directory",
		Run: commands.Rm,
	},
	"mv": {
		Description: "Move/Rename a file or empty directory",
		Run: commands.Mv,
	},
	"cp": {
		Description: "Copy a file",
		Run: commands.Cp,
	},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: gtools <command>")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	command, ok := cmdMap[cmd]
	if !ok {
		fmt.Println("Unknown command:", cmd)
		return
	}
	command.Run(args)
}
