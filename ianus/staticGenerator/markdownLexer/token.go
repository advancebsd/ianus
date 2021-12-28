package markdownLexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	/* Escape sequences */
	// TODO: Not complete
	ESCAPE_EXCLAMATION = "\\!"
	ESCAPE_QUOTE       = "\""
	ESCAPE_POUND       = "\\#"
	ESCAPE_DOLLAR      = "\\$"
	ESCAPE_PERCENT     = "\\%"

	/* Single character tokens */
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"
	LEFT_PAREN    = "("
	RIGHT_PAREN   = ")"
	INLINE_CODE   = "`"
	CODE_BLOCK    = "```"
	EXCLAMATION   = "!"
	NEW_LINE      = "\n"

	/* Emphasis Tokens */
	ITALIC      = "*"
	BOLD        = "**"
	BOLD_ITALIC = "***"
	QUOTE       = ">"

	/* Header tokens */
	HEADER_ONE   = "#"
	HEADER_TWO   = "##"
	HEADER_THREE = "###"

	/* List Tokens */
	// No support for asteric as bullet
	// Only support for unordered bullets currently
	BULLET_MINUS = "-"
	BULLET_PLUS  = "+"

	/* Check box tokens */
	UNCHECKED = "[ ]"
	CHECKED   = "[x]"

	HORIZONTAL_RULE = "---"

	CONTENT = "content"

	INVALID = "INVALID"
	EOF     = "EOF"
)
