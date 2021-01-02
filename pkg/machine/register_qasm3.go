package machine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Quant-Team/qvm/internal/qasm3"
	"github.com/Quant-Team/qvm/pkg/circuit"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func (r *register) ExitVersion(c *qasm3.VersionContext) {
	fmt.Print("Version: ")
	fmt.Print(c.Integer(0).GetText())
	if c.DOT() != nil {
		fmt.Print("." + c.Integer(1).GetText())
	}
	fmt.Println()
}

func (r *register) ExitInclude(c *qasm3.IncludeContext) {
	include := c.StringLiteral().GetText()
	fmt.Println("Include:", include)
	r.ParseProgramm(strings.ReplaceAll(include, "\"", ""))
}

func (r *register) EnterQuantumDeclaration(c *qasm3.QuantumDeclarationContext) {
	for _, ch := range c.IndexIdentifierList().GetChildren() {
		switch v := ch.(type) {
		case *qasm3.IndexIdentifierContext:
			for _, ch2 := range v.GetChildren() {
				switch v2 := ch2.(type) {
				case *qasm3.DesignatorContext:
					i, _ := strconv.Atoi(v2.Expression().GetText())
					for j := 0; j < i; j++ {
						r.q = append(r.q, circuit.Zero())
					}
				}
			}
		}
	}
}

func (r *register) EnterQuantumGateDefinition(c *qasm3.QuantumGateDefinitionContext) {
	signature := c.QuantumGateSignature().GetText()
	if _, ok := r.gates[signature]; !ok {
		r.gates[signature] = struct{}{}
		return
	}
	panic("gate " + signature + " exists")
}

func (r *register) ExitBitDeclaration(c *qasm3.BitDeclarationContext) {
	r.bits[c.Identifier().GetText()] = Zero
}

func (r *register) ParseProgramm(path string) {
	qasmFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read schema failed", err.Error())
		os.Exit(2)
	}

	// Setup the input
	is := antlr.NewInputStream(string(qasmFile))

	// Create the Lexer
	lexer := qasm3.Newqasm3Lexer(is)
	if r.debug {
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n",
				lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
		return
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := qasm3.Newqasm3Parser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(r, p.Program())
}
