package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/KadyrPoyraz/monkeygo/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    text := "Hello %s! This is the Monkey programming language!\n"
    fmt.Printf(text, user.Username)
    fmt.Printf("Feel free to type in commands\n")

    repl.Start(os.Stdin, os.Stdout)
}
