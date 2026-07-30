package main

import (
	"fmt"
	"os"
)

func pwd(arg []string) {
	if len(arg) > 0 {
		fmt.Println("Usage: gtools pwd")
	}
	
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get working directory:", err)
		return		
	}
	fmt.Println(dir)
}
