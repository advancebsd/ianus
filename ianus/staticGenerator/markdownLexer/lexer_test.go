package markdownLexer

import (
	"testing"
)

func TestInitializerLexer(t *testing.T) {
	testString := "This is a test string"
	l := new(Lexer)
	l.InitializeLexer(testString)
	if l.input != testString {
		t.Errorf("Testing strings do not match.\n\texpected: %s\n\tactual: %s\n", testString, l.input)
	}
	if l.readPosition != 1 {
		t.Errorf("Read position not properly initialized\n\texpected: 1\n\tactual: %d\n", l.readPosition)
	}
	if l.position != 0 {
		t.Errorf("Position not properly initialized\n\texpected: 0\n\tactual: %d\n", l.position)
	}
}

func TestReadChar(t *testing.T) {
	l := new(Lexer)
	l.InitializeLexer("hello")
	if l.ch != byte(l.input[0]) {
		t.Errorf("Improper Initialization")
	}
	l.readChar()
	if l.ch != byte(l.input[1]) {
		t.Errorf("Did no advance properly")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	testString := " is a string"
	l := new(Lexer)
	l.InitializeLexer(testString)
	l.skipWhiteSpace()
	if l.ch != byte(testString[1]) {
		t.Errorf("White space not skipped")
	}
}

/* Test for header tokens */
func TestHeaderTokens(t *testing.T) {
	headerOne := "# "
	headerTwo := "## "
	headerThree := "### "

	var tok Token
	l := new(Lexer)

	l.InitializeLexer(headerOne)
	tok = l.NextToken()
	if tok.Type != HEADER_ONE {
		t.Errorf("\nHeader token '#' not read properly")
	}
	l.InitializeLexer(headerTwo)
	tok = l.NextToken()
	if tok.Type != HEADER_TWO {
		t.Errorf("\nHeader token '##' not properly read")
	}
	l.InitializeLexer(headerThree)
	tok = l.NextToken()
	if tok.Type != HEADER_THREE {
		t.Errorf("\nHeader token '###' not properly read %s\n", tok.Type)
	}

}