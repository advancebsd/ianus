package gemtextRender

import (
	"testing"
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

func TestRenderHeaderTokens(t *testing.T) {
	var token markdownLexer.Token
	token.Type = markdownLexer.HEADER_ONE
	var str string
	str = renderMdTokenToGemtext(token)
	if str != "# " {
		t.Errorf("Did not properly render HEADER_ONE token")
	}
}
