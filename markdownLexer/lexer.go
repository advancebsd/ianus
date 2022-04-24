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
	if l.readPosition >= len(l.runes) {
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
func (l *Lexer) lookAheadNextChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.runes[l.readPosition]
}

/* Lex the character being excaped from in token */
func (l *Lexer) lexEscapeToken(id rune) Token {
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

	return string(l.runes[pos:l.position])
}

/* Check for italic token or bullet token */
func (l *Lexer) checkIfBulletOrItalic() string {
	curr := l.position
	if len(l.runes) <= 2 {
		return ITALIC
	}
	curr++;
	for l.runes[curr] != l.ch && curr < len(l.runes) {
		if l.runes[curr] == '\n' {
			return NEW_LINE
		}
		curr++;
	}
	return ITALIC
}

/* Lex the type of emphasis toke */
func (l *Lexer) lexEmphasisToken() Token {
	var t Token

	// if l.checkIfBulletOrItalic() == NEW_LINE {
	// 	t.Type = BULLET_MINUS
	// 	t.Literal = "+"
	// 	return t
	// }

	str := l.getRepeatCharToken(l.ch)

	if str == "*" || str == "_" {
		t.Type = ITALIC
		t.Literal = str
	} else if str == "**" || str == "__" {
		t.Type = BOLD
		t.Literal = str
	} else if str == "***" || str == "___" {
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

/* check if the character is a letter between A and Z, upper and lower case */
func (l *Lexer) isLetter() bool {
	return l.ch >= 'A' && l.ch <= 'Z' || l.ch >= 'a' && l.ch <= 'z' || l.ch >= 0x7E
}

/* check if the digit is a ascii number between 0 and 9 */
func (l *Lexer) isDigit() bool {
	return l.ch >= '0' && l.ch <= '9'
}

/* check for allowed punctuation in content block */
func (l *Lexer) isPunctuation() bool {
	return l.ch == '.' || l.ch == ',' || l.ch == '_' || l.ch == ':' || l.ch == '/' || l.ch == '?' || l.ch == '!' || l.ch == '\'' || l.ch == '"' || l.ch == '>' || l.ch == '<' || l.ch == ';' || l.ch == '%' || l.ch == '$' || l.ch == '=' || l.ch == '{' || l.ch == '}'
}

/* Checks if the current character is white space */
func (l *Lexer) isContentWhiteSpace() bool {
	return l.ch == ' '
}

/* Continuously read so long as markdown tokens are not read in */
func (l *Lexer) readContent() string {
	pos := l.position
	for l.isPunctuation() || l.isDigit() || l.isLetter() || l.isPunctuation() || l.isContentWhiteSpace() {
		l.readChar()
	}
	// for l.IsToken(l.runes[l.position]) == false {
	// 	l.readChar()
	// }
	return string(l.runes[pos:l.position])
}

/* Check is rune read is one of the markdown tokens */
func (l *Lexer) IsToken(ch rune) bool {
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

func (l *Lexer) lexDash() Token {
	pos := l.position
	var r []rune
	var t Token
	for l.ch == '-' {
		l.readChar()
	}
	r = l.runes[pos:l.position]
	if len(r) == 1 {
		t.Type = BULLET_MINUS
		t.Literal = string(r)
	} else if len(r) >= 3 {
		t.Type = HORIZONTAL_RULE
		t.Literal = string(r)
	} else {
		t.Type = INVALID
		t.Literal = string(r)
	}
	return t

}

/* Read the the next token in input */
func (l *Lexer) NextToken() Token {
	var token Token

	if l.checkPrevTokenForBlock() {
		token.Type = CONTENT
		token.Literal = l.readBlock()
		return token
	}

	switch l.ch {

	case '#':
		token = l.lexHeaderToken()
		return token
	case '_':
		token = l.lexEmphasisToken()
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
		token = l.lexDash()
		return token
	case '+':
		token.Type = BULLET_PLUS
		token.Literal = string(l.ch)
	case '\n':
		token.Type = NEW_LINE
		token.Literal = string(l.ch)
	case ' ':
		token.Type = WHITESPACE
		token.Literal = string(l.ch)
	case '>':
		token.Type = QUOTE
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
