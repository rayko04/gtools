package commands

import ( 
	"fmt"
	"os"
)

func Mkdir(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: gtools mkdir <dir>")
		return
	}

	dirname := args[0]
	err := os.Mkdir(dirname, 0755)
	
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}
}
