package gemtextRender

import (
	"testing"
	"github.com/advancebsd/ianus/ianus/staticGenerator/markdownLexer"
)

func testRenderHeaderTokens(t *testing.T) {
	var token Token
	token.Type = HEADER_ONE
	var str string
	str = renderMdTokenToGemtext(token)
	if str != "# " {
		t.Errorf("Did not properly render HEADER_ONE token")
	}
}
