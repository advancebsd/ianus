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

/* Read the the next token in input */
func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhiteSpace()

	// TODO: finish implementing state machine lexer
	switch l.ch {

	case '#':
		var c []byte
		for l.ch == '#' {
			c = append(c, l.ch)
			l.readChar()
		}
		if string(c) == "#" {
			token.Type = HEADER_ONE
			token.Literal = string(c)
		} else if string(c) == "##" {
			token.Type = HEADER_TWO
			token.Literal = string(c)
		} else if string(c) == "###" {
			token.Type = HEADER_THREE
			token.Literal = string(c)
		} else {
			token.Type = INVALID
			token.Literal = string(c)
		}
	case '[':
		token.Type = LEFT_BRACKET
		token.Literal = string(l.ch)
	case ']':
		token.Type = RIGHT_BRACKET
		token.Literal = string(l.ch)
	case '(':
		token.Type = LEFT_PAREN
		token.Literal = string(l.ch)
	case ')':
		token.Type = RIGHT_PAREN
		token.Literal = string(l.ch)
	case '`':
		var c []byte
		for l.ch == '`' {
			c = append(c, l.ch)
			l.readChar()
		}
		if string(c) == "`" {
			token.Type = INLINE_CODE
			token.Literal = string(c)
		} else if string(c) == "```" {
			token.Type = CODE_BLOCK
			token.Literal = string(c)
		} else {
			token.Type = INVALID
			token.Literal = string(c)
		}
	case '!':
		token.Type = EXCLAMATION
		token.Literal = string(l.ch)
	case '\n':
		token.Type = NEW_LINE
		token.Literal = string(l.ch)
	case 0 :
		token.Type = EOF
		token.Literal = ""
	}
	l.readChar()

	return token
}
