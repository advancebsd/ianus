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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold italic case")
	}
	expected := "<b><i>helloworld</i></b>"
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<blockquote>This is a quote</blockquote>\n"
	if result != expected {
		t.Errorf("Block quote rendering.\nExpected: %s\nResult: %s\n", expected, result)
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
	h := InitializeHtmlRender(tokens)
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
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document for bold case")
	}
	expected := "<ul>This is an item</ul>\n"
	if result != expected {
		t.Errorf("Bold rendering.\nExpected: %s\nResult: %s\n", expected, result)
	}
}

func TestLink(t *testing.T) {
	str := "[NetBSD](http://netbsd.org)"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	for token := l.NextToken(); token.Type != markdownLexer.EOF; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue render document for link")
	}
	expected := "<a href=\"http://netbsd.org\">NetBSD</a>"
	if result != expected {
		t.Errorf("Did not properly render link")
	}
}

func TestCheckBoxChecked(t *testing.T) {
	str := "[x]"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var token markdownLexer.Token
	for {
		token = l.NextToken()
		tokens = append(tokens, token)
		if token.Type == markdownLexer.EOF {
			break
		}
	}
	fmt.Println(tokens)
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document with checked checkbox")
	}
	expected := "<input type=\"checkbox\" checked>"
	if result != expected {
		t.Errorf("Fail to render checked checkbox\nExpected %s\nActual: %s\n", expected, result)
	}
}

func TestUnCheckBoxChecked(t *testing.T) {
	str := "[ ]"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var token markdownLexer.Token
	for {
		token = l.NextToken()
		tokens = append(tokens, token)
		if token.Type == markdownLexer.EOF {
			break
		}
	}
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering document with checked checkbox")
	}
	expected := "<input type=\"checkbox\" >"
	if result != expected {
		t.Errorf("Fail to render checked checkbox\nExpected %s\nActual: %s\n", expected, result)
	}
}

func TestHorizontalLineRule(t *testing.T) {
	str := "---"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var token markdownLexer.Token
	for {
		token = l.NextToken()
		tokens = append(tokens, token)
		if token.Type == markdownLexer.EOF {
			break
		}
	}
	h := InitializeHtmlRender(tokens)
	result, err := h.RenderDocument()
	if err != nil {
		t.Errorf("Issue rendering tokens for horizonatal rule test")
	}
	expected := "<hr>"
	if result != expected {
		t.Errorf("Failed to render horizontal rule\nExpected: %s\nActual:%s\n", expected, result)
	}
}

func TestSemicolon(t *testing.T) {
	str := "testing string;"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var token markdownLexer.Token
	var tokens []markdownLexer.Token
	token = l.NextToken()
	for token.Type != markdownLexer.EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}
	g := InitializeHtmlRender(tokens)
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Issue with rendering document from test string")
	}
	expected := "testing string;"
	if expected != result {
		t.Errorf("Expected and result do not match.\nExpected: %s\nResult: %s", expected, result)
	}
}
