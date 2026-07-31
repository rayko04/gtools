package registry

import (
	"gtools/commands"
)


var CmdMap = map[string]Command {

	"pwd": {
		Description: "Print working directory",
		Usage: "Usage: gtools pwd",
		Run: commands.Pwd,
	},
	"ls": {
		Description: "List directory content",
		Usage: "Usage: gtools ls <directory>",
		Run: commands.Ls,
	},
	"cat": {
		Description: "Print file content",
		Usage: "Usage: gtools cat <file>",
		Run: commands.Cat,
	},
	"touch": {
		Description: "Create a file if it does not exist",
		Usage: "Usage: gtools touch <filename>",
		Run: commands.Touch,
	},
	"mkdir": {
		Description: "Create directory",
		Usage: "Usage: gtools mkdir <directory>",
		Run: commands.Mkdir,
	},
	"rm": {
		Description: "Remove a file or an empty directory",
		Usage: "Usage: gtools rm <path>",
		Run: commands.Rm,
	},
	"mv": {
		Description: "Move/Rename a file or empty directory",
		Usage: "Usage: gtools mv <source> <destination>",
		Run: commands.Mv,
	},
	"cp": {
		Description: "Copy a file",
		Usage: "Usage: gtools cp <source> <destination>",
		Run: commands.Cp,
	},
	"version": {
		Usage:       "gtools version",
		Description: "Show version information",
		Run:         VersionCmd,
	},
}
