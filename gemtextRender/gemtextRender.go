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

/* Initialize the render by setting the token stream and idx to 0 */
func (g *GemtextRender) InitializeGemtextRender(ts []markdownLexer.Token) {
	g.tokenStream = ts
	g.idx = 0
}

/* Move the GemtextRender Struct's idx one forward so long as it is not at the end of the document */
func (g *GemtextRender) incrementIndex() bool {
	if g.idx >= len(g.tokenStream) {
		return false
	}
	g.idx += 1
	return true
}

/* Read the token currently pointed to by GemtextRender's index */
func (g *GemtextRender) readToken() (markdownLexer.Token, error) {
	if g.idx >= len(g.tokenStream) {
		var t markdownLexer.Token
		t.Type = markdownLexer.INVALID
		t.Literal = ""
		return t, errors.New("No more tokens to read")
	}
	return g.tokenStream[g.idx], nil
}

/* Called to process the input token stream to a string output of gemtext */
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

/* Look ahead to the next token */
func (g *GemtextRender) peekNextToken() (markdownLexer.Token, error) {
	if g.idx >= len(g.tokenStream) {
		var t markdownLexer.Token
		t.Type = markdownLexer.INVALID
		t.Literal = ""
		return t, errors.New("End of token stream")
	}
	return g.tokenStream[g.idx+1], nil
}

func (g *GemtextRender) hasPrevToken () bool {
	return g.idx > 0
}

func (g *GemtextRender) getPrevTokenType() markdownLexer.Token {
	return g.tokenStream[g.idx-1];
}

func (g *GemtextRender) resetTokenRenderForLinks(old_idx int) string {
	g.idx = old_idx
	return g.tokenStream[g.idx].Literal
}

/* State machine to handle the decision tree of rendering either a left bracket token */
/* or different forms of links that can exist from markdown tokens */
func (g *GemtextRender) renderLeftBracket() string {
	old_idx := g.idx
	var isLink bool = false

	var str string
	var link string
	var desc string
	var token markdownLexer.Token

	token, _ = g.readToken()
	g.incrementIndex()
	switch token.Type {
	case markdownLexer.LEFT_BRACKET :
		token, _ = g.readToken()
		g.incrementIndex()
		switch token.Type {
		case markdownLexer.CONTENT:
			desc = token.Literal
			token, _ = g.readToken()
			g.incrementIndex()
			switch token.Type {
			case markdownLexer.RIGHT_BRACKET:
				token, _ = g.readToken()
				g.incrementIndex()
				switch token.Type {
				case markdownLexer.LEFT_PAREN:
					token, _ = g.readToken()
					g.incrementIndex()
					switch token.Type {
					case markdownLexer.CONTENT:
						link = token.Literal
						token, _ = g.readToken()
						g.incrementIndex()
						switch token.Type {
						case markdownLexer.RIGHT_PAREN:
							str = "=> " + link + " " + desc + "\n"
							isLink = true
						default:
							str = g.resetTokenRenderForLinks(old_idx)
						}
					case markdownLexer.RIGHT_PAREN:
						str = "=> " + link + "\n"
						isLink = true
					default:
						str = g.resetTokenRenderForLinks(old_idx)
					}
				default:
					str = g.resetTokenRenderForLinks(old_idx)
				}
			default:
				str = g.resetTokenRenderForLinks(old_idx)
			}
		default:
			str = g.resetTokenRenderForLinks(old_idx)
		}
	default:
		str = g.resetTokenRenderForLinks(old_idx)
	}

	if isLink {
		if g.tokenStream[old_idx - 1].Type != markdownLexer.NEW_LINE {
			str = "\n" + str
		}
	}

	return str
}

func (g *GemtextRender) renderBulletMinus () string {
	if g.hasPrevToken() {
		if g.tokenStream[g.idx-1].Type == markdownLexer.NEW_LINE {
			return "*"
		}
	}
	return "-"
}

/* Takes a token and renders that token to gemtext */
func (g *GemtextRender) renderMdTokenToGemtext(t markdownLexer.Token) string {
	var str string
	switch t.Type {
	case markdownLexer.HEADER_ONE:
		str = "#"
	case markdownLexer.HEADER_TWO:
		str = "##"
	case markdownLexer.HEADER_THREE:
		str = "###"
	case markdownLexer.INLINE_CODE:
		str = t.Literal
	case markdownLexer.CODE_BLOCK:
		str = t.Literal
	case markdownLexer.BULLET_MINUS:
		str = g.renderBulletMinus()
	case markdownLexer.BULLET_PLUS:
		str = "*"
	case markdownLexer.UNCHECKED:
		str = t.Literal
	case markdownLexer.CHECKED:
		str = t.Literal
	case markdownLexer.QUOTE:
		str = t.Literal
	case markdownLexer.RIGHT_BRACKET:
		str = t.Literal
	case markdownLexer.LEFT_BRACKET:
		str = g.renderLeftBracket()
	case markdownLexer.RIGHT_PAREN:
		str = t.Literal
	case markdownLexer.LEFT_PAREN:
		str = t.Literal
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
	case markdownLexer.WHITESPACE:
		str = t.Literal
	default:
		str = ""
	}

	g.incrementIndex()

	return str
}
