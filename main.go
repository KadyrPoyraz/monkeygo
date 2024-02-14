package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/KadyrPoyraz/monkeygo/helpers"
	"github.com/KadyrPoyraz/monkeygo/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    text := fmt.Sprintf("Howdy %s! This is the Monkeygo programming language!", user.Username)
    fmt.Println(helpers.GetTextAlert(text) + "\n")
    fmt.Printf("Feel free to type in commands\n")

    repl.Start(os.Stdin, os.Stdout)
}
