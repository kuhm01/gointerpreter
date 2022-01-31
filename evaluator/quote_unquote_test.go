package evaluator

import (
	"monkey/object"
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(5)`,
			`5`,
		},
		{
			`quote(5 + 8)`,
			`(5 + 8)`,
		},
		{
			`quote(foobar)`,
			`foobar`,
		},
		{
			`quote(foobar + barfoo)`,
			`(foobar + barfoo)`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)

		if !ok {
			t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
		}

		if quote.Node == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
		}
	}
}

func TestQuoteUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(unquote(4))`,
			`4`,
		},
		{
			`quote(unquote(5 + 8))`,
			`13`,
		},
		{
			`quote(8 + unquote(4 + 4))`,
			`(8 + 8)`,
		},
		{
			`quote(unquote(4 + 4) + 8)`,
			`(8 + 8)`,
		},
		{
			`let foobar = 8; quote(foobar)`,
			`foobar`,
		},
		{
			`let foobar = 8; quote(unquote(foobar))`,
			`8`,
		},
		{
			`quote(unquote(4.1))`,
			`4.100000`,
		},
		{
			`quote(unquote(5.1 + 8.1))`,
			`13.200000`,
		},
		{
			`quote(8.1 + unquote(4.1 + 4.1))`,
			`(8.1 + 8.200000)`,
		},
		{
			`quote(unquote(4.1 + 4.1) + 8.1)`,
			`(8.200000 + 8.1)`,
		},
		{
			`let foobar = 8.1; quote(foobar)`,
			`foobar`,
		},
		{
			`let foobar = 8.1; quote(unquote(foobar))`,
			`8.100000`,
		},
		{
			`quote(unquote(true))`,
			`true`,
		},
		{
			`quote(unquote(true == false))`,
			`false`,
		},
		{
			`quote(unquote(quote(4 + 4)))`,
			`(4 + 4)`,
		},
		{
			`let qie = quote(4 + 4); quote(unquote(4 + 4) + unquote(qie))`,
			`(8 + (4 + 4))`,
		},
		{
			`qie := quote(4 + 4); quote(unquote(4 + 4) + unquote(qie))`,
			`(8 + (4 + 4))`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)

		if !ok {
			t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
		}

		if quote.Node == nil {
			t.Fatalf("quote.Node is nil")
		}

		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", quote.Node.String(), tt.expected)
		}
	}
}
