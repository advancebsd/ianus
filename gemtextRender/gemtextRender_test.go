package gemtextRender

import (
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
	"testing"
	"fmt"
)

func TestRenderHeaderTokens(t *testing.T) {
	var err error
	str := "# ## ###"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var tok markdownLexer.Token
	tok = l.NextToken()
	for tok.Type != markdownLexer.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}
	g := new(GemtextRender)
	g.InitializeGemtextRender(tokens)
	tok, err = g.readToken()
	if err != nil {
		t.Errorf("Issue reading token in sample")
	}
	if g.renderMdTokenToGemtext(tok) != "#" {
		t.Errorf("Did not properly render HEADER_ONE token")
	}
	tok, _ = g.readToken()
	if g.renderMdTokenToGemtext(tok) != " " {
		t.Errorf("Could not render white space between HEADER_ONE and HEADER_TWO tokens")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "##" {
		t.Errorf("Did not properly render HEADER_TWO token")
	}
	tok, _ = g.readToken()
	if g.renderMdTokenToGemtext(tok) != " "{
		t.Errorf("Could not render whitespace between HEADER_TWO and HEADER_THREE tokens")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "###" {
		t.Errorf("Did not properly render HEADER_THREE token")
	}
}

func TestRenderBulletPoints(t *testing.T) {
	var err error
	str := "+ -"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var tok markdownLexer.Token
	tok = l.NextToken()
	for tok.Type != markdownLexer.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}
	var g GemtextRender
	g.InitializeGemtextRender(tokens)
	tok, err = g.readToken()
	if err != nil {
		t.Errorf("Issue reading token during test")
	}
	tok, _ = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "*" {
		t.Errorf("Did not properly render BULLET_PLUS token")
	}
	tok, _ = g.readToken()
	g.renderMdTokenToGemtext(tok)
	tok, _ = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "*" {
		t.Errorf("Did not properly render BULLET_MINUS token")
	}
}

func TestRenderOnString(t *testing.T) {
	str := "# HeaderOne\nSomeInformation about test\n *HelloWorld*"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var tok markdownLexer.Token
	tok = l.NextToken()
	for tok.Type != markdownLexer.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}
	var g GemtextRender
	g.InitializeGemtextRender(tokens)
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Not properly reading tokens from lexer")
	}
	expected := "# HeaderOne\nSomeInformation about test\n HelloWorld"
	if result != expected {
		t.Errorf("Did no properly render to gemtext the test string")
		fmt.Println(result)
	}
}

func TestLinkGeneration(t *testing.T) {
	str := "[netbsd.org](NetBSD Website)"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var tok markdownLexer.Token
	tok = l.NextToken()
	for tok.Type != markdownLexer.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}

	var g GemtextRender
	g.InitializeGemtextRender(tokens)
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Issue with rendering tokens")
	}
	expected := "=> netbsd.org NetBSD Website"
	if result != expected {
		t.Errorf("Issue with rendering a link")
	}
}

// need to troubleshoot test case
func TestLeftBracketRender(t *testing.T) {
	str := "[ netbsd.org(one)"
	var l markdownLexer.Lexer
	l.InitializeLexer(str)
	var tokens []markdownLexer.Token
	var tok markdownLexer.Token
	tok = l.NextToken()
	for tok.Type != markdownLexer.EOF {
		tokens = append(tokens, tok)
		tok = l.NextToken()
	}

	fmt.Println(tokens)

	var g GemtextRender
	g.InitializeGemtextRender(tokens)
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Issue with rendering tokens")
	}
	expected := "[ netbsd.org(one)"
	if result != expected {
		t.Errorf("Issue with rendering bracket token followed by content")
		fmt.Println(result)
	}
}
