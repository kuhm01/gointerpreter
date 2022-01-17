package routeoption

import (
	"fmt"
)

func RouteOption(arg string) {
	switch arg {
	case "-h", "help":
		printHelpOption()
	}
}

func printHelpOption() {
	fmt.Printf("Monkey is a tool for managing Monkey source code.\n\nUsage:\n\n")
	fmt.Printf("         monkey <command> [argument]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("         -h help       print Information")
}
