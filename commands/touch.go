package commands

import ( 
	"fmt"
	"os"
)

func Touch(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: gtools touch <filename>")
		return
	}

	for _, filename := range args {
		file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0644)
	
		if err != nil {
			fmt.Println("Failed to create file:", err)
			continue
		}
		file.Close()
	}
	
}
