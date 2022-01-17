package main

//Heaven R.I.P Avicii

import (
	"fmt"
	"monkey/repl"
	"monkey/routeoption"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		routeoption.RouteOption(args)
	}
}
