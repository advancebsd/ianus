package htmlRender

import (
  "errors"
  markdownLexer "github.com/advancebsd/ianus/markdownLexer"
  "os"
)

type HtmlRender struct {
	token_stream []markdownLexer.Token
	idx int
	length int
	current_token markdownLexer.Token
	page string
}

func InitializeHtmlRender(stream []markdownLexer.Token) *HtmlRender {
	var render = new(HtmlRender)
	render.token_stream = stream
	render.idx = 0
	render.length = len(stream)
	render.page = ""
	render.current_token = render.token_stream[0]
	return render
}

func (h *HtmlRender) RenderDocument() (string, error) {
	for h.current_token.Type != markdownLexer.EOF {
		str, err := h.render_token()
		if err != nil {
			return "", err
		}
		h.page += str
		h.increment_token()
	}
	return h.page, nil
}

func (h *HtmlRender) get_token_by_index(i int) (markdownLexer.Token) {
	return h.token_stream[i]
}

func (h *HtmlRender) increment_token() bool {
	if h.idx >= h.length {
		return false
	}
	h.idx += 1
	h.current_token = h.token_stream[h.idx]
	return true
}

func (h *HtmlRender) peek_prev_token() (markdownLexer.Token, error) {
	if h.idx > 0 {
		return h.token_stream[h.idx-1], nil
	}
	return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Can not index an array less than 0")
}

func (h *HtmlRender) peek_next_token() (markdownLexer.Token, error) {
	if h.idx < h.length-1 {
		return h.token_stream[h.idx+1], nil
	}
	return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Can not look at tokens that don't exist beyond bounds")
}

func (h *HtmlRender) peek_to_newline_or_token() (markdownLexer.Token, error) {
	if h.idx >= h.length-1 {
		return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Can not peek beyond array bounds")
	}
	position := h.idx + 1
	token := h.token_stream[position]
	for {
		if token.Type == markdownLexer.NEW_LINE {
			return token, nil
		}
		if token.Type == h.current_token.Type {
			return token, nil
		}
		if token.Type == markdownLexer.EOF {
			return token, nil
		}
		position++
		token = h.token_stream[position]
	}
	return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Could not peek ahead")
}

func get_tags(token markdownLexer.Token) (string, string, error) {
	switch token.Type {
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
		return "<b><i>", "</i></b>", nil
	case markdownLexer.BULLET_MINUS:
		return "<li>", "</li>", nil
	case markdownLexer.BULLET_PLUS:
		return "<li>", "</li>", nil
	default:
		return "", "", errors.New("Could not match appropraite tag")
	}
}

func (h *HtmlRender) render_token() (string, error) {
	switch h.current_token.Type {
	case markdownLexer.CONTENT:
		return h.current_token.Literal, nil
	case markdownLexer.NEW_LINE:
		return h.current_token.Literal, nil
	case markdownLexer.WHITESPACE:
		return h.current_token.Literal, nil
	case markdownLexer.LEFT_PAREN:
		return h.current_token.Literal, nil
	case markdownLexer.RIGHT_PAREN:
		return h.current_token.Literal, nil
	case markdownLexer.RIGHT_BRACKET:
		return h.current_token.Literal, nil
	case markdownLexer.HORIZONTAL_RULE:
		return "<hr>", nil //TODO
	case markdownLexer.CHECKED:
		return "<input type=\"checkbox\" checked>", nil
	case markdownLexer.UNCHECKED:
		return "<input type=\"checkbox\" >", nil
	case markdownLexer.HEADER_ONE:
		return "", nil // TODO
	case markdownLexer.HEADER_TWO:
		return "", nil // TODO
	case markdownLexer.HEADER_THREE:
		return "", nil // TODO
	case markdownLexer.ITALIC:
		return h.render_italic_token()
	case markdownLexer.BOLD:
		return "", nil // TODO
	case markdownLexer.BOLD_ITALIC:
		return "", nil // TODO
	case markdownLexer.QUOTE:
		return "", nil // TODO
	case markdownLexer.INLINE_CODE:
		return "", nil // TODO
	case markdownLexer.CODE_BLOCK:
		return "", nil // TODO
	case markdownLexer.BULLET_MINUS:
		return "", nil // TODO
	case markdownLexer.BULLET_PLUS:
		return "", nil // TODO
	case markdownLexer.LEFT_BRACKET:
		return "", nil // TODO
	default:
		return "", errors.New("Issue rendering HTML")
	}
}

func (h *HtmlRender) render_header_tokens() (string, error) {
	return "", errors.New("Method not implemented")
}

func (h *HtmlRender) render_italic_token() (string, error) {
	var str string = ""
	// Reference previous and next token to see it italic token is at the end of the line
	prev_token, err := h.peek_prev_token()
	if err != nil {
		return "", err
	}
	if prev_token.Type == markdownLexer.NEW_LINE {
		next_token, err := h.peek_next_token()
		if err != nil {
			return "", err
		}
		if next_token.Type == markdownLexer.WHITESPACE {
			// Since previous token is newline and next token is a white space
			// make the assumption that the asterick is actually meant to be
			// a bullt point
			start_tag, end_tag, error := get_tags(h.current_token)
			if err != nil {
				return "", error
			}
			str += start_tag
			h.increment_token();
			for h.current_token.Type != markdownLexer.NEW_LINE || h.current_token.Type != markdownLexer.EOF {
				render, error := h.render_token()
				if error != nil {
					return "", errors.New("Could not render token")
				}
				str += render
				h.increment_token()
			}
			str += end_tag
			return str, nil
		}
		// Not facing a situation where a asterick is supposed to be a bullet point
		// look ahead an see which tokens are encountered first between
		// italic tokens, newline tokens, or end of file tokens
		token, err := h.peek_to_newline_or_token()
		if err != nil {
			return "", errors.New("Issue looking ahead when rendering italic token")
		}
		// if look ahead and see new line or EOF instead of ending italic token
		// render the asterick as is in place
		if token.Type == markdownLexer.NEW_LINE || token.Type == markdownLexer.EOF {
			return h.current_token.Literal, nil
		}
		// look ahead showed that next token we care about is italics token
		// render all content between the two astericks and surrounf them with html
		// italic tags
		if token.Type == markdownLexer.ITALIC {
			start_tag, end_tag, error := get_tags(h.current_token)
			if error != nil {
				return "", error
			}
			str += start_tag
			h.increment_token()
			for {
				if h.current_token.Type == markdownLexer.ITALIC {
					str += end_tag
					return str, nil
				}
				render, err := h.render_token()
				if err != nil {
					return "", errors.New("Issue rendering tokens after italic start")
				}
				str += render
				h.increment_token()
			}
		}
	}

	// This is the case where the previous token is not a new line token
	// so the assumption can be made that the token is not meant to be a bullet point
	// look ahead to see if the token stream will lead to a italic token, new line token
	// or end of file token next
	token, error := h.peek_to_newline_or_token()
	if error != nil {
		return "", error
	}
	// significant tokens to come up are the newline or end of file token
	// so just render the asterick as is in place
	if token.Type == markdownLexer.NEW_LINE || token.Type == markdownLexer.EOF {
		return h.current_token.Literal, nil
	}

	// next significant token is an italics
	// so render as expected for an italics tag
	start_tag, end_tag, err := get_tags(h.current_token)
	if err != nil {
		return "", errors.New("Issue getting tags for rendering")
	}
	str += start_tag
	h.increment_token()
	for {
		if h.current_token.Type == markdownLexer.ITALIC {
			str += end_tag
			return str, nil
		}
		render, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering tokens in between italics")
		}
		str += render
		h.increment_token()
	}
	return "", errors.New("Issue rendering italic tokens")
}
