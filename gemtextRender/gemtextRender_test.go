package gemtextRender

import (
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
	"testing"
)

func TestRenderHeaderTokens(t *testing.T) {
	var err error
	str := "# ## ###"
	l := new(markdownLexer.Lexer)
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
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "# " {
		t.Errorf("Did not properly render HEADER_ONE token")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "## " {
		t.Errorf("Did not properly render HEADER_TWO token")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "### " {
		t.Errorf("Did not properly render HEADER_THREE token")
	}
}

func TestRenderBulletPoints(t *testing.T) {
	var err error
	str := "+ -"
	l := new(markdownLexer.Lexer)
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
		t.Errorf("Issue reading token during test")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "* " {
		t.Errorf("Did not properly render BULLET_PLUS token")
	}
	tok, err = g.readToken()
	if g.renderMdTokenToGemtext(tok) != "* " {
		t.Errorf("Did not properly render BULLET_MINUS token")
	}
}

func TestRenderOnString(t *testing.T) {
	str := "# HeaderOne\nSomeInformation about test\n *HelloWorld*"
	l := new(markdownLexer.Lexer)
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
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Not properly reading tokens from lexer")
	}
	expected := "# HeaderOne\nSomeInformation about test\nHelloWorld"
	if result != expected {
		t.Errorf("Did no properly render to gemtext the test string")
	}
}

func TestLinkGeneration(t *testing.T) {
	str := "[netbsd.org](NetBSD Website)"
	l := new(markdownLexer.Lexer)
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
	result, err := g.RenderDocument()
	if err != nil {
		t.Errorf("Issue with rendering tokens")
	}
	expected := "=> netbsd.org NetBSD Website"
	if result != expected {
		t.Errorf("Issue with rendering a link")
	}
}
