package lexer

import (
	"go_interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatal(
				"tests[%d] - tokentype wrong. expected=%q, received=%q", i, tt.expectedType, toke.Type
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatal(
				"tests[%d] - tokenliteral wrong. expected=%q, received=%q", i, tt.expectedLiteral, toke.Literal
			)
		}
	}
}
