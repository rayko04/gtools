package commands

import ( 
	"fmt"
	"os"
)

func Cat(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: gtools cat <file>")
		return
	}

	file := args[0]
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return	
	}

	fmt.Println(string(bytes))
}

