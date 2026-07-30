package commands

import ( 
	"fmt"
	"os"
)

func Rm(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: gtools rm <path>")
		return
	}

	for _, path := range args {
		err := os.Remove(path)
		
		if err != nil {
			fmt.Println("Failed to remove:", path, "-", err)
		}
	}
}
