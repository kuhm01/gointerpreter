package routeoption

import (
	"fmt"
	"strings"
)

func RouteOption(args []string) {
	arg := args[1]
	switch arg {
	case "-h", "help":
		printHelpOption()
	case "-c", "start":
		if len(args) == 2 {
			fmt.Printf("Not input file. Monkey <command> [argument]\n")
		} else {
			startLPEmonkeyinterpreting(args[2])
		}
	}
}

func printHelpOption() {
	fmt.Printf("Monkey is a tool for managing Monkey source code.\n\nUsage:\n\n")
	fmt.Printf("         monkey <command> [argument]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("         -h help       print Information")
	fmt.Printf("\n\n\n")
}

func startLPEmonkeyinterpreting(argv string) {
	splitedFileName := strings.Split(argv, ".")
	if splitedFileName[1] != "monkey" {
		fmt.Printf("Not Monkey File. got=%s\n", argv)
		return
	}
	//파일 읽기 및 lexing, parsing, evaluating 구현 해야 함
}
