package markdownLexer

import (
	"os"
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
	if l.ch != rune(l.input[0]) {
		t.Errorf("Improper Initialization")
	}
	l.readChar()
	if l.ch != rune(l.input[1]) {
		t.Errorf("Did no advance properly")
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	testString := " is a string"
	l := new(Lexer)
	l.InitializeLexer(testString)
	l.skipWhiteSpace()
	if l.ch != rune(testString[1]) {
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
	if len(tok) != 8 {
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
	if tok[6].Type != WHITESPACE {
		t.Errorf("Did not read the white space token")
	}
	if tok[7].Type != CODE_BLOCK {
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
	if tok.Type != WHITESPACE {
		t.Errorf("\nWhite space not read")
	}
	tok = lex.NextToken()
	if tok.Type != CONTENT {
		t.Errorf("Did not properly read the content 'Hello'")
	}
	tok = lex.NextToken()
	if tok.Type != BOLD_ITALIC {
		t.Errorf("Did not read the bold italic token at between new line token and 'BSD',  recv:  |%d|", lex.position)
	}
}

func TestBulletPoints(t *testing.T) {
	str := "+ Hello world"
	l := new(Lexer)
	l.InitializeLexer(str)
	var token Token
	token = l.NextToken()
	if token.Type != BULLET_PLUS {
		t.Errorf("Did not properly parse plus signed used as a bullet")
	}
	token = l.NextToken()
	if token.Type != WHITESPACE {
		t.Errorf("Did not parse the white space properly after the plus bullet token")
	}
	token = l.NextToken()
	if token.Type != CONTENT {
		t.Errorf("Content not properly parsed after bullt point plus")
	}

	str1 := "- Hello World"
	l1 := new(Lexer)
	l1.InitializeLexer(str1)
	var token1 Token
	token1 = l1.NextToken()
	if token1.Type != BULLET_MINUS {
		t.Errorf("Did not properly parse the minus bullet point")
	}
	token1 = l1.NextToken()
	if token1.Type != WHITESPACE {
		t.Errorf("Did not properly parse the white space after the minus bullet token")
	}
	token1 = l1.NextToken()
	if token1.Type != CONTENT {
		t.Errorf("Did not properly parse content after minus bullet point")
	}
}

/* Testing the parsing of bracket tokens for LEFT_BRACKET, UNCHECKED, and CHECKED tokens */
func TestBracketTokens(t *testing.T) {
	str := "[ ] [x] ["
	l := new(Lexer)
	l.InitializeLexer(str)
	var token Token
	token = l.NextToken()
	if token.Type != UNCHECKED {
		t.Errorf("Failed to parse the unchecked markdown token properly")
	}
	token = l.NextToken()
	if token.Type != WHITESPACE {
		t.Errorf("Could not parse white space after unchecked token")
	}
	token = l.NextToken()
	if token.Type != CHECKED {
		t.Errorf("Failed to parse the checked markdown token properly")
	}
	token = l.NextToken()
	if token.Type != WHITESPACE {
		t.Errorf("Could not parse the whitespace after the checked token")
	}
	token = l.NextToken()
	if token.Type != LEFT_BRACKET {
		t.Errorf("Failed to parse the left bracket token properly")
	}
}

/* Testing the tokens generated from a link markdown style input */
func TestMarkdownLinks(t *testing.T) {
	str := "[netbsd](http://netbsd.com)"
	l := new(Lexer)
	l.InitializeLexer(str)
	var token Token
	token = l.NextToken()
	if token.Type != LEFT_BRACKET {
		t.Errorf("Did not properly parse left bracket token")
	}
	token = l.NextToken()
	if token.Type != CONTENT {
		t.Errorf("Did not properly parse link description as content")
	}
	token = l.NextToken()
	if token.Type != RIGHT_BRACKET {
		t.Errorf("Did not properly parse the right bracket token")
	}
	token = l.NextToken()
	if token.Type != LEFT_PAREN {
		t.Errorf("Did not properly parse the left parenthesis token")
	}
	token = l.NextToken()
	if token.Type != CONTENT {
		t.Errorf("Did not properly parse the link as a content token")
	}
	token = l.NextToken()
	if token.Type != RIGHT_PAREN {
		t.Errorf("Did not properly parse the right parenthesis as a token")
	}
}

/* test against a local source located in the sample folder of this directory */
func TestAgainstRemoteSource(t *testing.T) {
	file := "sample/SampleMD.md"
	content, err := os.ReadFile(file)
	if err != nil {
		t.Errorf("Could not run test due to unable to locate sample file")
	}
	l := new(Lexer)
	l.InitializeLexer(string(content))
	var token Token
	var tokens []Token

	token = l.NextToken()
	for token.Type != EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	for i, curr := range tokens {
		if curr.Type == INVALID {
			t.Errorf("Found an invalid token! [%d] %s : %s", i, curr.Type, curr.Literal)
		}
	}
}

/* test punctuation */
func TestContentWithPunctuation(t *testing.T) {
	str := "This, is a string with \"some punctuation \" to check that there isn't any issue reading it"
	var lex Lexer
	lex.InitializeLexer(str)
	var token Token
	var tokens []Token
	token = lex.NextToken()
	for token.Type != EOF {
		tokens = append(tokens, token)
		token = lex.NextToken()
	}
	if len(tokens) != 1 {
		t.Errorf("To many tokens were read in")
	}
	if tokens[0].Type != CONTENT {
		t.Errorf("The incorrect token type was generated")
	}
}

/* test for unicode characters */
func TestUnicodeCharacters(t *testing.T) {
	str := "£YË"
	var lex Lexer
	lex.InitializeLexer(str)
	var tokens []Token
	var token Token
	token = lex.NextToken()
	for token.Type != EOF {
		tokens = append(tokens, token)
		token = lex.NextToken()
	}
	if len(tokens) != 1 {
		t.Errorf("Did not read unicode content properly")
	}
	if tokens[0].Type != CONTENT {
		t.Errorf("Did not type the token properly")
	}
	if tokens[0].Literal != "£YË" {
		t.Errorf("Did not catpure the literal of token correctly")
	}
}
