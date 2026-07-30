package commands

import (
	"fmt"
	"os"
)

type FileInfo struct {
	Name string
	IsDir bool
}

func (f FileInfo) Display() {
    if f.IsDir {
		fmt.Println("DIR", f.Name)
	} else {
		fmt.Println("FILE", f.Name)
	}
}

func Ls(args []string) {
	
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return
	}

	for _, entry := range dirEntries {
		info := FileInfo{
			Name: entry.Name(), 
			IsDir: entry.IsDir(),
		}
		info.Display()
	}
}
