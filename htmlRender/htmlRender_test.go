package htmlRender

import (
	"fmt"
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
	"testing"
)

func TestH1Rendering(t *testing.T) {
	var str string
	var l markdownLexer.Lexer
	var token markdownLexer.Token
	var tokens []markdownLexer.Token
	var h HtmlRender
	var result string
	var expected string
	var err error

	str = "# Hello world\n"
	l.InitializeLexer(str)
	token = l.NextToken()
	for token.Type != markdownLexer.EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	h.InitializeHtmlRender(tokens)
	result, err = h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document")
	}
	expected = "<h1>Hello world</h1>\n"
	if result != expected {
		t.Errorf("Issue with rendering header token.\nExpected: %s\nActual: %s", expected, result)
	}
}

func TestH2Rendering(t *testing.T) {
	var str string
	var l markdownLexer.Lexer
	var token markdownLexer.Token
	var tokens []markdownLexer.Token
	var h HtmlRender
	var result string
	var expected string
	var err error

	str = "## Hello world\n"
	l.InitializeLexer(str)
	token = l.NextToken()
	for token.Type != markdownLexer.EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	h.InitializeHtmlRender(tokens)
	result, err = h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document")
	}
	expected = "<h2>Hello world</h2>\n"
	if result != expected {
		t.Errorf("Issue with rendering header token.\nExpected: %s\nActual: %s", expected, result)
	}
}

func TestH3Rendering(t *testing.T) {
	var str string
	var l markdownLexer.Lexer
	var token markdownLexer.Token
	var tokens []markdownLexer.Token
	var h HtmlRender
	var result string
	var expected string
	var err error

	str = "### Hello world\n"
	l.InitializeLexer(str)
	token = l.NextToken()
	for token.Type != markdownLexer.EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	h.InitializeHtmlRender(tokens)
	result, err = h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document")
	}
	expected = "<h3>Hello world</h3>\n"
	if result != expected {
		t.Errorf("Issue with rendering header token.\nExpected: %s\nActual: %s", expected, result)
	}
}

func TestBoldEmphasis(t *testing.T) {
	str := "**helloworld**"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<b>helloworld</b>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestItalicEmphasis(t *testing.T) {
	str := "*helloworld*"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for italic case")
	}
	expected := "<i>helloworld</i>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestBoldItalicEmphasis(t *testing.T) {
	str := "***helloworld***"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold italic case")
	}
	expected := "<b><i>helloworld</b></i>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestBoldAndItalicSeperate(t *testing.T) {
	str := "**In this sentence, *this* is italic**"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold italic case")
	}
	expected := "<b>In this sentence, <i>this</i> is italic</b>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestInLineCode(t *testing.T) {
	str := "`helloworld`"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for in_line code case")
	}
	expected := "<code>helloworld</code>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestCodeBlock(t *testing.T) {
	str := "```helloworld```"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<code>helloworld</code>"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestQuote(t *testing.T) {
	str := "> This is a quote\n"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<blockquote>This is a quote</blockquote>\n"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestQuote2(t *testing.T) {
	str := "> This is a quote\nHello World"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<blockquote>This is a quote</blockquote>\nHello World"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestUnorderedList(t *testing.T) {
	str := "- This is an item\n"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<ul>This is an item</ul>\n"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestLink (t *testing.T) {
	str := "[NetBSD](http://netbsd.org)"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
	var h HtmlRender
	h.InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue render document for link")
	}
	expected := "<a href=\"http://netbsd.org\">NetBSD</a>"
	if result != expected {
		t.Errorf("Did not properly render link")
	}
}
