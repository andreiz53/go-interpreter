package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("could not read user: ", err)
	}
	fmt.Printf("Hello %s! This is my programming language.\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
