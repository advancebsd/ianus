package gemtextRender

import (
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
	"testing"
	"os"
)

/* Test rendering  header tokens from markdown tokens */
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

/* Test rendering gemtext bullet points from markdown tokens */
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

/* Testing the rendering of strings, newlines, text emphasis and headers */
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
	}
}

/* testing generating a gemtext link from a markdown link */
func TestLinkGeneration(t *testing.T) {
	str := "[NetBSD Website](netbsd.org)"
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
	expected := "=> netbsd.org NetBSD Website\n"
	if result != expected {
		t.Errorf("Issue with rendering a link")
	}
}

// need to troubleshoot test case
func TestLeftBracketRender(t *testing.T) {
	str := "[netbsd.org(one)"
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
	expected := "[netbsd.org(one)"
	if result != expected {
		t.Errorf("Issue with rendering bracket token followed by content")
	}
}

/* Run rendering against a sample file */
func TestSampleFile(t *testing.T) {
	file := "sample/SampleMD.md"
	content, err := os.ReadFile(file)
	if err != nil {
		t.Errorf("Make sure that file exists as sample file could not be found")
	}
	var l markdownLexer.Lexer
	l.InitializeLexer(string(content))
	var token markdownLexer.Token
	var tokens []markdownLexer.Token
	token = l.NextToken()
	for token.Type != markdownLexer.EOF {
		tokens = append(tokens, token)
		token = l.NextToken()
	}

	var g GemtextRender
	g.InitializeGemtextRender(tokens)
	result, err := g.RenderDocument()
	//fmt.Println(result)

	file_expected := "sample/expected.gmi"
	expected, err := os.ReadFile(file_expected)
	if err != nil {
		t.Errorf("Could not find the file that holds the expected results")
	}

	result_bytes := []byte(result)
	expected_bytes := []byte(string(expected))

	// if len(result_bytes) != len(expected_bytes) {
	// 	t.Errorf("Size of output file does not match expected file")
	// 	t.Errorf("Size result: %d, size expected: %d", len(result_bytes), len(expected_bytes))
	// }

	for i := 0; i < len(expected_bytes); i++ {
		if result_bytes[i] != expected_bytes[i] {
			t.Errorf("Non matching characters as index: %d. Expected: %c, Actual: %c", i, expected_bytes[i], result_bytes[i])
		}
	}
}
