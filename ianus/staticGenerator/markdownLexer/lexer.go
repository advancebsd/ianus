package markdownLexer

/**
  * TODO: Implement lexer
  * TODO: Write test cases
 */

type Lexer struct {
	input string
	position int // current position in input
	readPosition int // next position in input
	ch byte
	tokens []Token
}

/* Return the token stream in this instance of the lexer */
func (l *Lexer) GetTokens() []Token {
	return l.tokens
}

/* Read the the next token in input */
func (l *Lexer) NextToken() {
	// TODO: implement
}

/* Set the input of Lexer instance */
func (l *Lexer) SetInput (in string) {
	l.input = in
}
