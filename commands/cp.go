package commands

import ( 
	"fmt"
	"os"
	"io"
)

func Cp(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: gtools cp <source> <destination>")
		return
	}

	src, dst := args[0], args[1]
	
	srcfile, srcerr := os.Open(src)
	if srcerr != nil {
		fmt.Println("Failed to open src file:", srcerr)
		return
	}
	defer srcfile.Close()

	dstfile, dsterr := os.Create(dst)
	if dsterr != nil {
		fmt.Println("Failed to create dst file:", dsterr)
		return
	}
	defer dstfile.Close()

	_, err := io.Copy(dstfile, srcfile)
	
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}

}
