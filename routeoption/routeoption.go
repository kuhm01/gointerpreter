package routeoption

import (
	"fmt"
	"io"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
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
	default:
		fmt.Printf("Monkey. Wrong Option\n")
	}
}

func printHelpOption() {
	fmt.Printf("Monkey is a tool for managing Monkey source code.\n\nUsage:\n\n")
	fmt.Printf("         monkey <command> [argument]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("         -h help       print Information\n")
	fmt.Printf("         -c start      Start Interpreting\n")
	fmt.Printf("\n\n")
}

func startLPEmonkeyinterpreting(argv string) {
	if !isFile(argv) {
		fmt.Printf("Enter the monkey file.\n")
		return
	}

	dat, err := ioutil.ReadFile(argv)
	if err != nil {
		fmt.Printf("Monkey File Error. We don't Interpreting this\n")
		return
	}
	input := string(dat)

	out := os.Stdout
	env := object.NewEnvironment()
	macroenv := object.NewEnvironment()

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluator.DefineMacros(program, macroenv)
	expanded := evaluator.ExpandMacros(program, macroenv)

	evaluated := evaluator.Eval(expanded, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func isFile(name string) bool {
	if !strings.Contains(name, ".") {
		fmt.Printf("Wrong File. got=%s\n", name)
		return false
	}

	splitedFileName := strings.Split(name, ".")
	if splitedFileName[1] != "monkey" {
		fmt.Printf("Not Monkey File. got=%s\n", name)
		return false
	}

	return true
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Interpreter Error. don't ran Interpreting\n")
	io.WriteString(out, " parser Error:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
