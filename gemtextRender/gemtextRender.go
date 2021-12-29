package gemtextRender

import (
	"errors"
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

type GemtextRender struct {
	tokenStream []markdownLexer.Token
	idx int
	page string
}

func (g *GemtextRender) InitializeGemtextRender(ts []markdownLexer.Token) {
	g.tokenStream = ts
	g.idx = 0
}

func (g *GemtextRender) incrementIndex() bool {
	if g.idx >= len(g.tokenStream) {
		return false
	}
	g.idx += 1
	return true
}

func (g *GemtextRender) readToken() (markdownLexer.Token, error) {
	if g.idx >= len(g.tokenStream) {
		var t markdownLexer.Token
		t.Type = markdownLexer.INVALID
		t.Literal = ""
		return t, errors.New("No more tokens to read")
	}
	return g.tokenStream[g.idx], nil
}

func (g *GemtextRender) RenderDocument() (string, error) {
	var gemtext string
	var curr_token markdownLexer.Token
	var err error
	curr_token, err = g.readToken()
	if err != nil {
		return "", errors.New("No tokens to read")
	}
	for curr_token.Type != markdownLexer.EOF {
		gemtext += g.renderMdTokenToGemtext(curr_token)
		curr_token, err = g.readToken()
		if err != nil {
			break
		}
	}
	return gemtext, nil
}

func (g *GemtextRender) peekNextToken() (markdownLexer.Token, error) {
	if g.idx >= len(g.tokenStream) {
		var t markdownLexer.Token
		t.Type = markdownLexer.INVALID
		t.Literal = ""
		return t, errors.New("End of token stream")
	}
	return g.tokenStream[g.idx+1], nil
}

/**
 * TODO: Implement render to gemtext
 *       TODO: Handle rendering of brackets for links
 */

func (g *GemtextRender) renderMdTokenToGemtext(t markdownLexer.Token) string {
	var str string
	switch t.Type {
	case markdownLexer.HEADER_ONE:
		str = "# "
	case markdownLexer.HEADER_TWO:
		str =  "## "
	case markdownLexer.HEADER_THREE:
		str = "### "
	case markdownLexer.INLINE_CODE:
		str = t.Literal
	case markdownLexer.CODE_BLOCK:
		str = t.Literal
	case markdownLexer.BULLET_MINUS:
		str = "* "
	case markdownLexer.BULLET_PLUS:
		str = "* "
	case markdownLexer.UNCHECKED:
		str = t.Literal
	case markdownLexer.CHECKED:
		str = t.Literal
	case markdownLexer.QUOTE:
		str = t.Literal
	case markdownLexer.RIGHT_BRACKET:

	case markdownLexer.LEFT_BRACKET:

	case markdownLexer.RIGHT_PAREN:

	case markdownLexer.LEFT_PAREN:

	case markdownLexer.ITALIC:
		str = ""
	case markdownLexer.BOLD:
		str = ""
	case markdownLexer.BOLD_ITALIC:
		str = ""
	case markdownLexer.HORIZONTAL_RULE:
		str = t.Literal
	case markdownLexer.CONTENT:
		str = t.Literal
	case markdownLexer.NEW_LINE:
		str = string('\n')
	default:
		str = ""
	}

	g.incrementIndex()

	return str
}
