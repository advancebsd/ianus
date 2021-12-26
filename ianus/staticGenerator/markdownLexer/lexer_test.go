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
	headerOne := "#"
	headerTwo := "##"
	headerThree := "###"

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

/* Test for emphasis tokens */
func TestEmphasisTokens(t *testing.T) {
	headerOne := "*"
	headerTwo := "**"
	headerThree := "***"

	var tok Token
	l := new(Lexer)

	l.InitializeLexer(headerOne)
	tok = l.NextToken()
	if tok.Type != ITALIC {
		t.Errorf("\nHeader token '#' not read properly")
	}
	l.InitializeLexer(headerTwo)
	tok = l.NextToken()
	if tok.Type != BOLD {
		t.Errorf("\nHeader token '##' not properly read")
	}
	l.InitializeLexer(headerThree)
	tok = l.NextToken()
	if tok.Type != BOLD_ITALIC {
		t.Errorf("\nHeader token '###' not properly read %s\n", tok.Type)
	}
}

/* Testing single character tokens */
func TestSingleCharTokenLexer(t *testing.T) {
	testString := "[]()!` ```"
	l := new(Lexer)
	l.InitializeLexer(testString)
	var tok []Token
	var token Token = l.NextToken()
	for token.Type != EOF {
		tok = append(tok, token)
		token = l.NextToken()
	}
	if len(tok) != 7 {
		t.Errorf("Did not read the proper amount of tokens")
	}
	if tok[0].Type != LEFT_BRACKET {
		t.Errorf("Did not read left bracket token")
	}
	if tok[1].Type != RIGHT_BRACKET {
		t.Errorf("Did not read right bracket token")
	}
	if tok[2].Type != LEFT_PAREN {
		t.Errorf("Did not read left parenthesis token")
	}
	if tok[3].Type != RIGHT_PAREN {
		t.Errorf("Did not read right parenthesis token")
	}
	if tok[4].Type != EXCLAMATION {
		t.Errorf("Did not read exclamation token")
	}
	if tok[5].Type != INLINE_CODE {
		t.Errorf("Did not read in line code token")
	}
	if tok[6].Type != CODE_BLOCK {
		t.Errorf("Did not read code block token")
	}

}

/* Test for lexing escape tokens */
func TestEscapeTokens(t *testing.T) {
	str := "\n"
	l := new(Lexer)
	l.InitializeLexer(str)
	var tokens []Token
	token := l.NextToken()
	for token.Type != EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	if tokens[0].Type != NEW_LINE {
		t.Errorf("\nDid not properly read new line token")
	}
}

/* Test for reading in content */
func TestReadContent(t *testing.T) {
	str := "hello . world \n one# two * three"
	l := new(Lexer)
	l.InitializeLexer(str)
	var tokens []Token
	var token Token
	token = l.NextToken()
	for token.Type != EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	if tokens[0].Type != CONTENT {
		t.Errorf("Could not parse 'hello . world ' as content")
	}
	if tokens[1].Type != NEW_LINE {
		t.Errorf("Could not parse the new line as a token")
	}

	str_one := "### Hello World ***BSD***"
	var lex Lexer
	lex.InitializeLexer(str_one)
	var tok Token
	tok = lex.NextToken()
	if tok.Type != HEADER_THREE {
		t.Errorf("Did not properly read in token for header in the beginning")
	}
	tok = lex.NextToken()
	if tok.Type != CONTENT {
		t.Errorf("Did not properly read the content 'Hello World'")
	}
	tok = lex.NextToken()
	if tok.Type != BOLD_ITALIC {
		t.Errorf("Did not read the bold italic token at between new line token and 'BSD',  recv:  |%d|", lex.position)
	}
}
