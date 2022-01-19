package htmlRender

import (
	"errors"
	"os"
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

func (h *HtmlRender) readToken() (markdownLexer.Token, error) {
	if h.idx >= len(h.tokenStream) {
		var t markdownLexer.Token
		t.Type = markdownLexer.INVALID
		t.Literal = ""
		return t, errors.New("No more tokens to read")
	}
	return h.tokenStream[h.idx], nil
}

func (h *HtmlRender) RenderDocument() (string, error) {
	var token markdownLexer.Token
	var err error
	token, err = h.readToken()
	if err != nil {
		return "", errors.New("No tokens to read")
	}
	for token.Type != markdownLexer.EOF {
		h.page += h.renderMdTokenToHtml(token)
		token, err = h.readToken()
		if err != nil {
			break
		}
	}
	return h.page, nil
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
	h.incrementIndex()
	if h.tokenStream[h.idx].Type == markdownLexer.WHITESPACE {
		h.incrementIndex()
	}
	if err != nil {
		return "", errors.New("Issue with getting the closing tag for a given token")
	}
	for { 
		currStr := h.renderMdTokenToHtml(h.tokenStream[h.idx])
 		if currStr == markdownLexer.NEW_LINE {
			str = str + endHeader + "\n"
			return str, nil
		} else if currStr == markdownLexer.EOF {
			return "", errors.New("Header did not terminate before end of file")
		}
		str = str + currStr
	}
	return str, nil
}

func (h *HtmlRender) handleEmphasis() string {
	var str string
	var err error
	var end_tag string
	tag := h.tokenStream[h.idx].Type

	str, end_tag, err = h.getTags()
	if err != nil {
		panic("Tag not recognized")
		os.Exit(1)
	}

	h.incrementIndex()
	for token, err := h.readToken(); token.Type != tag; token, err = h.readToken()  {
		str = str + h.renderMdTokenToHtml(token)
		if err != nil {
			panic("Reached a token that is not recognized")
			os.Exit(1)
		}
	}

	str = str + end_tag

	return str

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
			panic("Document ended before terminating a header")
			os.Exit(1)
		}
		return str
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
		str = t.Literal
	case markdownLexer.EOF:
		str = ""
	case markdownLexer.BOLD:
		str = h.handleEmphasis()
	case markdownLexer.ITALIC:
		str = h.handleEmphasis()
	case markdownLexer.BOLD_ITALIC:
		str = h.handleEmphasis()
	case markdownLexer.INLINE_CODE:
		str = h.handleEmphasis()
	case markdownLexer.CODE_BLOCK:
		str = h.handleEmphasis()
	}

	h.incrementIndex()

	return str
}
