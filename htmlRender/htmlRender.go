package htmlRender

import (
	"errors"
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

// TODO: Implement rest of render
// TODO: Test render
// TODO: Write test cases for render

type HtmlRender struct {
	tokenStream []markdownLexer.Token
	idx         int
	page        string
}

// Initialize the Html Render using a token stream from the markdownLexer
func (h *HtmlRender) InitializeHtmlRender(ts []markdownLexer.Token) {
	h.tokenStream = ts
	h.idx = 0
}

// Increment the location of the token stream one forward
func (h *HtmlRender) incrementIndex() bool {
	if h.idx >= len(h.tokenStream) {
		return false
	}
	h.idx += 1
	return true
}

// Using the current token, get the opening and closing tags in html
func (h *HtmlRender) getTags() (string, string, error) {
	switch h.tokenStream[h.idx].Type {
	case markdownLexer.HEADER_ONE:
		return "<h1>", "</h1>", nil
	case markdownLexer.HEADER_TWO:
		return "<h2>", "</h2>", nil
	case markdownLexer.HEADER_THREE:
		return "<h3>", "</h3>", nil
	case markdownLexer.INLINE_CODE:
		return "<code>", "</code>", nil
	case markdownLexer.CODE_BLOCK:
		return "<code>", "</code>", nil
	case markdownLexer.QUOTE:
		return "<blockquote>", "</blockquote>", nil
	case markdownLexer.BOLD:
		return "<b>", "</b>", nil
	case markdownLexer.ITALIC:
		return "<i>", "</i>", nil
	case markdownLexer.BOLD_ITALIC:
		return "<b><i>", "</b></i>", nil
	case markdownLexer.BULLET_MINUS:
		return "<ul>", "</ul>", nil
	case markdownLexer.BULLET_PLUS:
		return "<ul>", "</ul>", nil
	default:
		return "", "", errors.New("Undefined token for getting end tag")

	}
}

func (h *HtmlRender) handleHeaderTokens() (string, error) {
	var str string
	var endHeader string
	var err error
	str, endHeader, err = h.getTags()
	if err != nil {
		return "", errors.New("Issue with getting the closing tag for a given token")
	}
	for h.incrementIndex() {
		currStr := h.renderMdTokenToHtml(h.tokenStream[h.idx])
		if currStr == markdownLexer.NEW_LINE {
			str = str + endHeader
			return str, nil
		} else if currStr == markdownLexer.EOF {
			return "", errors.New("Header did not terminate before end of file")
		}
		str = str + currStr
	}
	return str, nil
}

// TODO: Write out the render to handle cases
// Render the token stream from lexing markdown to HTML text
func (h *HtmlRender) renderMdTokenToHtml(t markdownLexer.Token) string {
	var str string
	var err error
	switch t.Type {
	case markdownLexer.HEADER_ONE:
		str, err = h.handleHeaderTokens()
		if err != nil {
			// handle some error here
		}
	case markdownLexer.HEADER_TWO:
		str, err = h.handleHeaderTokens()
		if err != nil {
			// handle some error here
		}
	case markdownLexer.HEADER_THREE:
		str, err = h.handleHeaderTokens()
		if err != nil {
			// handle some error here
		}
	case markdownLexer.NEW_LINE:
		str = "\n"
	case markdownLexer.WHITESPACE:
		str = " "
	case markdownLexer.CONTENT:
		str = "<p>" + t.Literal + "</p>"
	}

	h.incrementIndex()

	return str
}
