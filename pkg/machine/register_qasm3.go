package machine

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/Quant-Team/qvm/internal/qasm3"
	"github.com/Quant-Team/qvm/pkg/circuit"
	"github.com/Quant-Team/qvm/pkg/circuit/gates/one"
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
	ipath := path.Join(r.dir, strings.ReplaceAll(include, "\"", ""))
	r.ParseProgram(ipath)
}

func (r *register) EnterQuantumDeclaration(c *qasm3.QuantumDeclarationContext) {
	for _, ch := range c.IndexIdentifierList().GetChildren() {
		switch v := ch.(type) {
		case *qasm3.IndexIdentifierContext:
			for _, ch2 := range v.GetChildren() {
				switch v2 := ch2.(type) {
				case *qasm3.DesignatorContext:
					i, _ := strconv.Atoi(v2.Expression().GetText())
					r.qbName[v.Identifier().GetText()] = r.NewRegPos(i)
					for j := 0; j < i; j++ {
						r.q = append(r.q, nil)
					}
				}
			}
		}
	}
}

func (r *register) EnterQuantumGateCall(c *qasm3.QuantumGateCallContext) {
	v := c.QuantumGateName().GetText()
	switch v {
	case "reset":
		qb := c.IndexIdentifierList().GetText()
		rp := r.qbName[qb]
		for j := rp.GetStart(); j < rp.GetEnd(); j++ {
			r.q[j] = circuit.Zero()
		}
	case "U":
	case "CX":
	default:
		if _, ok := r.gates[v]; !ok {
			panic("not found defenition of gate: " + v)
		}
		switch v {
		case "h":
			ci := c.IndexIdentifierList().(*qasm3.IndexIdentifierListContext)
			ci2 := ci.IndexIdentifier(0).(*qasm3.IndexIdentifierContext)
			di := ci2.Designator().(*qasm3.DesignatorContext)
			i, _ := strconv.Atoi(di.Expression().GetText())
			r.q[i].Apply(one.H())
		case "cx":
		}
	}
}

func (r *register) EnterGlobalStatement(c *qasm3.GlobalStatementContext) {
	// fmt.Println(">>", c.GetText())
}

func (r *register) EnterQuantumGateDefinition(c *qasm3.QuantumGateDefinitionContext) {
	signature := c.QuantumGateSignature().(*qasm3.QuantumGateSignatureContext)
	i := signature.Identifier().GetText()
	if _, ok := r.gates[i]; !ok {
		r.gates[i] = struct{}{}
		return
	}
	panic("gate " + i + " exists")
}

func (r *register) ExitBitDeclaration(c *qasm3.BitDeclarationContext) {
	r.bits[c.Identifier().GetText()] = Zero
}

func (r *register) ParseProgram(path string) {
	qasmFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read schema failed", err.Error())
		os.Exit(2)
	}

	fmt.Println("parse:", path)
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
	var parse antlr.Tree = p.Program()
	if parse == nil {
		parse = p.Statement()
	}
	antlr.ParseTreeWalkerDefault.Walk(r, parse)
}
