package markdownLexer

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int // next position in input
	runes        []rune
	ch           rune
	tokens       []Token
}

/* Return the token stream in this instance of the lexer */
func (l *Lexer) GetTokens() []Token {
	return l.tokens
}

/* Set the input of Lexer instance */
func (l *Lexer) InitializeLexer(in string) {
	l.input = in
	l.runes = []rune(in)
	l.position = 0
	l.readPosition = 1
	l.ch = l.runes[l.position]
}

/* Advances the position in input */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.runes[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhiteSpace() {
	if l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
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
func (l *Lexer) lexEscapeToken(id byte) Token {
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
func (l *Lexer) getRepeatCharToken(ch rune) string {
	pos := l.position
	for l.ch == ch {
		l.readChar()
	}

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
	} else {
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

/* Lex the dash tokens as either bullet points or horizontal rules */
func (l *Lexer) lexHorizontalRule() Token {
	var t Token

	str := l.getRepeatCharToken(l.ch)
	if str == "-" {
		t.Type = BULLET_MINUS
		t.Literal = str
	} else if str == "---" {
		t.Type = HORIZONTAL_RULE
		t.Literal = str
	} else {
		t.Type = INVALID
		t.Literal = str
	}

	return t
}

/* check if the character is a letter between A and Z, upper and lower case */
func isLetter(ch rune) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
}

/* check if the digit is a ascii number between 0 and 9 */
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

/* check for allowed punctuation in content block */
func isPunctuation(ch rune) bool {
	return ch == '.' || ch == ',' || ch == '_' || ch == ':' || ch == '/' || ch == '?' || ch == '!' || ch == '\'' || ch == '"'
}

func isContentWhiteSpace(ch rune) bool {
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

/* Checks for LEFT_BRACKET, CHECKED, and UNCHECKED tokens */
func (l *Lexer) lexBracketToken() Token {
	var token Token
	pos := l.position
	switch l.ch {
	case '[':
		l.readChar()
		switch l.ch {
		case ' ':
			l.readChar()
			switch l.ch {
			case ']':
				l.readChar()
				token.Type = UNCHECKED
				token.Literal = l.input[pos:l.position]
			default:
				token.Type = LEFT_BRACKET
				token.Literal = "["
			}
		case 'x':
			l.readChar()
			switch l.ch {
			case ']':
				l.readChar()
				token.Type = CHECKED
				token.Literal = l.input[pos:l.position]
			default:
				token.Type = INVALID
				token.Literal = l.input[pos:l.position]
			}
		default:
			token.Type = LEFT_BRACKET
			token.Literal = "["
		}
	default:
		token.Type = INVALID
		token.Literal = string(l.ch)
	}

	return token
}

func (l *Lexer) checkPrevTokenForBlock() bool {
	size := l.getNumberTokens()
	if size < 1 {
		return false
	}

	if l.tokens[size].Type == INLINE_CODE || l.tokens[size].Type == CODE_BLOCK {
		return true
	}

	return false
}

func (l *Lexer) checkForBlockEnd() bool {
	if l.ch == '`' {
		return true
	}

	return false
}

func (l *Lexer) readBlock() string {
	pos := l.position
	for l.checkForBlockEnd() != true {
		l.readChar()
	}

	return l.input[pos:l.position]
}

/* Get the previous Token read */
func (l *Lexer) getNumberTokens() int {
	return len(l.tokens)
}

/* Read the the next token in input */
func (l *Lexer) NextToken() Token {
	var token Token

	if l.checkPrevTokenForBlock() {
		token.Type = CONTENT
		token.Literal = l.readBlock()
		return token
	}

	// TODO: finish implementing state machine lexer
	switch l.ch {

	case '#':
		token = l.lexHeaderToken()
		return token
	case '*':
		token = l.lexEmphasisToken()
		return token
	case '[':
		token = l.lexBracketToken()
		return token
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
		var c []rune
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
		token = l.lexHorizontalRule()
		return token
	case '+':
		token.Type = BULLET_PLUS
		token.Literal = string(l.ch)
	case '\n':
		token.Type = NEW_LINE
		token.Literal = string(l.ch)
	case ' ' :
		token.Type = WHITESPACE
		token.Literal = string(l.ch)
	case 0:
		token.Type = EOF
		token.Literal = ""
	default:
		content := l.readContent()
		token.Type = CONTENT
		token.Literal = content
		return token
	}

	l.readChar()

	return token
}
