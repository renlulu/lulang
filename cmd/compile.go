package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"lulang/complier"
	lexer2 "lulang/lexer"
	parser2 "lulang/parser"
	"os"
)

var source string

func init() {
	compileCmd.Flags().StringVarP(&source, "source", "s", "", "path of source file")
}

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compile lulang source code into bytecode",
	Long:  "compile lulang source code into bytecode",
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			fmt.Println("source file is empty")
			os.Exit(1)
		}
		input, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Println("read source file error: " + err.Error())
			os.Exit(1)
		}

		lexer := lexer2.New(string(input))
		parser := parser2.New(lexer)
		compiler := complier.New()
		err = compiler.Compile(parser.ParseProgram())
		if err != nil {
			fmt.Println("compile error: " + err.Error())
			os.Exit(1)
		}
		bytecode := compiler.Bytecode()
		// todo constant
		_ = ioutil.WriteFile(fmt.Sprintf("%s.code", source), []byte(bytecode.Instructions.String()), 0644)
	},
}
