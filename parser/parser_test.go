package parser

import (
	"testing"

	"go-interpreter/ast"
	"go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	// input with errors
	// 	input := `
	// let x 5;
	// let = 10;
	// let 838383;
	// `

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements has incorrect length. expected: 3, got: %d", program.Statements)
	}

	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for idx, tc := range testCases {
		statement := program.Statements[idx]
		if !testLetStatement(t, statement, tc.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("statement token literal not 'let', got %v", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not *ast.Statement, got %T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("let statement name value not '%s', got '%s'", name, letStatement.Name.Value)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errs))
	for _, err := range errs {
		t.Errorf("parser error: %v", err)
	}
	t.Fail()
}
