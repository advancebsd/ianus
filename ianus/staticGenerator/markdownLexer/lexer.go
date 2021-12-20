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
func (l *Lexer) InitializeLexer (in string) {
	l.input = in
	l.position = 0
	l.readPosition = 1
	l.ch = l.input[l.position]
}

/* Advances the position in input */
func (l *Lexer) readChar () {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhiteSpace() {
	if l.ch ==  ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
