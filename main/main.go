package main

import (
	"fmt"
	"lulang/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
