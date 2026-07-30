package main

import ( 
	"fmt"
	"os"
)

func cat(arg []string) {
	if len(arg) < 1 {
		fmt.Println("Usage: gtools cat <file>")
		return
	}

	file := arg[0]
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return	
	}

	fmt.Println(string(bytes))
}
