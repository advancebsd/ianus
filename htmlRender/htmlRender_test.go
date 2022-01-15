package htmlRender

import (
	"testing"
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

func TestH1Rendering (t *testing.T) {
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


func TestH2Rendering (t *testing.T) {
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

func TestH3Rendering (t *testing.T) {
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
