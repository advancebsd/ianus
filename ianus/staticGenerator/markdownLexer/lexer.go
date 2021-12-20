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
	var token Token

	l.skipWhiteSpace()

	// TODO: finish implementing state machine lexer
	switch l.ch {
	case '#':
		switch l.lookAheadNextChar() {
		case '#':
			l.readChar()
			switch l.lookAheadNextChar() {
			case '#':
				token.Type = HEADER_THREE
				token.Literal = "###"

			case ' ':
				token.Type = HEADER_TWO
				token.Literal = "##"
			}
		case ' ' :
			token.Type = HEADER_ONE
			token.Literal = string(l.ch)
		}
	}

	l.readChar()
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

func (l *Lexer) lookAheadNextChar() byte {
	return l.input[l.readPosition]
}
