package main

import ( 
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: gtools <command>")
		return
	}

	cmd := os.Args[1]
	arg := os.Args[2:]

	switch cmd {
	case "pwd": 
		pwd(arg)
	
	case "ls" : 
		ls(arg)

	case "cat" : 
		cat(arg)

	default: 
		fmt.Println("unknown cmd:", cmd)
	}	
}
