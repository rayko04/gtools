package registry

import "fmt"

const Version = "1.0.0"

func VersionCmd(args []string) {
	if len(args) > 0 {
		fmt.Println("Usage: gtools version")
		return
	}
    fmt.Println("gtools version", Version)
}
