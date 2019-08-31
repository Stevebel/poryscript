package lexer

import (
	"testing"

	"github.com/huderlem/poryscript/token"
)

func TestNextToken(t *testing.T) {
	input := `script BugContestOfficer_EnterContest_23
{
	if (var(VAR_BUG_CONTEST_PRIZE) != ITEM_NONE) {
		giveitem_std(VAR_BUG_CONTEST_PRIZE)
		if (flag(FLAG_TEST) == TRUE) {
			setvar(VAR_BUG_CONTEST_PRIZE, ITEM_NONE)
		} elif (var(VAR_TEST) <= 5) {
		} else { ##
		#}
		<
		>
		>=
		=
		!
		/
		do
		while
		("Hello\n"
		"I'm glad to see$")
		raw RawTest ` + "`" + `
	step
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.SCRIPT, "script"},
		{token.IDENT, "BugContestOfficer_EnterContest_23"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.VAR, "var"},
		{token.LPAREN, "("},
		{token.IDENT, "VAR_BUG_CONTEST_PRIZE"},
		{token.RPAREN, ")"},
		{token.NEQ, "!="},
		{token.IDENT, "ITEM_NONE"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "giveitem_std"},
		{token.LPAREN, "("},
		{token.IDENT, "VAR_BUG_CONTEST_PRIZE"},
		{token.RPAREN, ")"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.FLAG, "flag"},
		{token.LPAREN, "("},
		{token.IDENT, "FLAG_TEST"},
		{token.RPAREN, ")"},
		{token.EQ, "=="},
		{token.TRUE, "TRUE"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "setvar"},
		{token.LPAREN, "("},
		{token.IDENT, "VAR_BUG_CONTEST_PRIZE"},
		{token.COMMA, ","},
		{token.IDENT, "ITEM_NONE"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.ELSEIF, "elif"},
		{token.LPAREN, "("},
		{token.VAR, "var"},
		{token.LPAREN, "("},
		{token.IDENT, "VAR_TEST"},
		{token.RPAREN, ")"},
		{token.LTE, "<="},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.GTE, ">="},
		{token.ILLEGAL, "="},
		{token.ILLEGAL, "!"},
		{token.ILLEGAL, "/"},
		{token.DO, "do"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.STRING, "Hello\\n\nI'm glad to see$"},
		{token.RPAREN, ")"},
		{token.RAW, "raw"},
		{token.IDENT, "RawTest"},
		{token.RAWSTRING, "\tstep"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokenType wrong. Expected=%q, Got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. Expected=%q, Got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
