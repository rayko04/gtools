package commands

import (
	"fmt"
	"os"
)

func Pwd(args []string) {
	if len(args) > 0 {
		fmt.Println("Usage: gtools pwd")
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", err)
		return		
	}
	fmt.Println(dir)
}
