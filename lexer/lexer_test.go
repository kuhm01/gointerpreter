package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	!-*/5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return /# true #/ false;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"} 

	let float = 10.1;
	/# puts("hello World") #/ 
	10.1 < 10;

	if (5.5 < 10) {
		return true;
		/# 
		puts("hello")
		#/
	} else {
		return false;
	}
	/#
	comments
	#/
	/##########/

	10.1 == 10.1;
	10 != 9.1;
	[1, 2.1];  /# sdfse [1, 2, 3] #/
	!-/*5.5;
	!/5.5*-;
	5.5*!;

	5.555555555555555555555555555555555;
	5.;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.SEMICOOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "float"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.1"},
		{token.SEMICOOLON, ";"},
		{token.FLOAT, "10.1"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.SEMICOOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.FLOAT, "5.5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOOLON, ";"},
		{token.RBRACE, "}"},
		{token.FLOAT, "10.1"},
		{token.EQ, "=="},
		{token.FLOAT, "10.1"},
		{token.SEMICOOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.FLOAT, "9.1"},
		{token.SEMICOOLON, ";"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.FLOAT, "2.1"},
		{token.RBRACKET, "]"},
		{token.SEMICOOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.FLOAT, "5.5"},
		{token.SEMICOOLON, ";"},
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.FLOAT, "5.5"},
		{token.ASTERISK, "*"},
		{token.MINUS, "-"},
		{token.SEMICOOLON, ";"},
		{token.FLOAT, "5.5"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.SEMICOOLON, ";"},
		{token.FLOAT, "5.555555555555555555555555555555555"},
		{token.SEMICOOLON, ";"},
		{token.ILLEGAL, "5."},
		{token.SEMICOOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype worng. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
