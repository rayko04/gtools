package commands

import ( 
	"fmt"
	"os"
)

func Mv(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: gtools mv <source> <destination>")
		return
	}

	src, dst := args[0], args[1]
	err := os.Rename(src, dst)
	
	if err != nil {
		fmt.Println("Failed to move:", err)
		return
	}
}
