package registry 

import ( 
	"fmt"
)

func Help(args []string) {
	if len(args) > 0 {
		fmt.Println("Usage: gtools help")
		return
	}
	if len(args) == 0 {
		fmt.Println("gtools - A collection of filesystem utilities")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("    gtools <command> [arguments]")
		fmt.Println()
		fmt.Println("Commands:")

		for name, cmd := range CmdMap {
			fmt.Printf("    %-10s %s\n", name, cmd.Description)
		}
		return
	}

	command := args[0]

	cmd, ok := CmdMap[command]
	if !ok {
		fmt.Println("Unknown command:", command)
		return
	}

	fmt.Println("Usage:")
	fmt.Printf("    %s\n\n", cmd.Usage)

	fmt.Println("Description:")
	fmt.Printf("    %s\n", cmd.Description)
}
