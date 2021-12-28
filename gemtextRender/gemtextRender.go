package gemtextRender

import (
	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

/**
 * TODO: Implement render to gemtext
 * TODO: Manage module imports
 */

func renderMdTokenToGemtext(t markdownLexer.Token) string {
	var str string
	switch t.Type {
	case "HEADER_ONE":
		str = "# "
	case "HEADER_TWO":
		str =  "## "
	case "Header_THREE":
		str = "### "
	}

	return str
}
