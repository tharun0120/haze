package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tharun0120/haze/evaluator"
	"github.com/tharun0120/haze/lexer"
	"github.com/tharun0120/haze/object"
	"github.com/tharun0120/haze/parser"
	"github.com/tharun0120/haze/token"
)

const PROMPT string = ">> "

func Start(in io.Reader, out io.Writer, dev bool) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
			continue
		}

		if dev {
			io.WriteString(out, "Parsed program:\n")
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			// io.WriteString(out, "Evaluated program:\n")
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "There are some errors in your code:\n")
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
