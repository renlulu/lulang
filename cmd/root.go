package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(interpreterCmd)
	RootCmd.AddCommand(compileCmd)
}

var RootCmd = &cobra.Command{
	Use:   "lulang",
	Short: "language for fun",
	Long:  `language for fun`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
