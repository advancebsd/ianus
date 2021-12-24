package markdownLexer
import "fmt"
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

/* Look ahead to the next character in the stream */
func (l *Lexer) lookAheadNextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

/* Lex the character being excaped from in token */
func (l *Lexer) lexEscapeToken(id byte) Token{
	var t Token

	switch id {
	case '!':
		t.Type = ESCAPE_EXCLAMATION
		t.Literal = string(l.ch)
	case '$':
		t.Type = ESCAPE_DOLLAR
		t.Literal = string(l.ch)
	case '#':
		t.Type = ESCAPE_POUND
		t.Literal = string(l.ch)
	case '%':
		t.Type = ESCAPE_PERCENT
		t.Literal = string(l.ch)
	case '"':
		t.Type = ESCAPE_QUOTE
		t.Literal = string(l.ch)
	default:
		t.Type = INVALID
		t.Literal = string(l.ch)
	}

	return t
}

/* Used to get repeacted characters that may be a token */
func (l *Lexer) getRepeatCharToken(ch byte) string {
	pos := l.position
	for l.ch == ch {
		l.readChar()
	}
	fmt.Printf(l.input[pos:l.position])
	return l.input[pos:l.position]
}

/* Lex the type of emphasis toke */
func (l *Lexer) lexEmphasisToken() Token {
	var t Token

	str := l.getRepeatCharToken(l.ch)

	if str == "*" {
		t.Type = ITALIC
		t.Literal = str
	} else if str == "**" {
		t.Type = BOLD
		t.Literal = str
	} else if str == "***" {
		t.Type = BOLD_ITALIC
		t.Literal = str
	} else  {
		t.Type = INVALID
		t.Literal = str
	}

	return t
}

/* Lex the type of header token */
func (l *Lexer) lexHeaderToken() Token {
	var t Token

	str := l.getRepeatCharToken(l.ch)

	if str == "#" {
		t.Type = HEADER_ONE
		t.Literal = str
	} else if str == "##" {
		t.Type = HEADER_TWO
		t.Literal = str
	} else if str == "###" {
		t.Type = HEADER_THREE
		t.Literal = str
	} else {
		t.Type = INVALID
		t.Literal = str
	}

	return t
}

func isLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isPunctuation(ch byte) bool {
	return ch == '.' || ch == ',' || ch == '_'
}

func isContentWhiteSpace(ch byte) bool {
	return ch == ' '
}

func (l *Lexer) readContent() string {
	pos := l.position
	for isLetter(l.ch) || isDigit(l.ch) || isPunctuation(l.ch) || isContentWhiteSpace(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isToken(ch byte) bool {
	switch ch {
	case '#':
		return true
	case '*':
		return true
	case '[':
		return true
	case ']':
		return true
	case '(':
		return true
	case ')':
		return true
	case '`':
		return true
	case '!':
		return true
	case '-':
		return true
	case '+':
		return true
	case '\n':
		return true
	default:
		return false
	}
	return false
}

/* Read the the next token in input */
func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhiteSpace()

	// TODO: finish implementing state machine lexer
	switch l.ch {

	case '#':
		token = l.lexHeaderToken()
		return token
	case '*':
		token = l.lexEmphasisToken()
		return token
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
		return token
	case '!':
		token.Type = EXCLAMATION
		token.Literal = string(l.ch)
	case '-':
		token.Type = BULLET_MINUS
		token.Literal = string(l.ch)
	case '+':
		token.Type = BULLET_PLUS
		token.Literal = string(l.ch)
	case '\n':
		token.Type = NEW_LINE
		token.Literal = string(l.ch)
	case 0:
		token.Type = EOF
		token.Literal = ""
	default:
		content := l.readContent()
		token.Type = CONTENT
		token.Literal = content
	}
	l.readChar()

	return token
}
