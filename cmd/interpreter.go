package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lulang/repl"
	"os"
	user2 "os/user"
)

var interpreterCmd = &cobra.Command{
	Use:   "interpreter",
	Short: "run lulang interpreter",
	Long:  "run lulang interpreter",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user2.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s!\n",
			user.Username)
		repl.StartInterpreter(os.Stdin, os.Stdout)
	},
}
