package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path"

	"github.com/tharun0120/haze/evaluator"
	"github.com/tharun0120/haze/lexer"
	"github.com/tharun0120/haze/object"
	"github.com/tharun0120/haze/parser"
	"github.com/tharun0120/haze/repl"
	"github.com/tharun0120/haze/token"
)

func main() {
	if len(os.Args) > 1 {
		runFile(os.Args[1])
	} else {
		runREPL()
	}
}

func runREPL() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s, Welcome to the Haze Interpreter!\n", user.Username)
	fmt.Printf("This is the Haze REPL, Type away!\n")
	fmt.Printf("To exit, type in \"exit\"!\n")
	repl.Start(os.Stdin, os.Stdout, false)
}

func runFile(pathStr string) {
	f, err := os.Open(pathStr)

	if err != nil {
		fmt.Printf("error opening file")
		os.Exit(1)
	}

	extension := path.Ext(pathStr)

	if extension != ".haze" {
		fmt.Printf(".haze file is required")
		os.Exit(1)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	env := object.NewEnvironment()

	for scanner.Scan() {
		l := lexer.New(scanner.Text())
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			repl.PrintParserErrors(os.Stdout, p.Errors())
			break
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err.Error())
	}
}
