package htmlRender

import (
  "errors"
  markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

type HtmlRender struct {
	token_stream []markdownLexer.Token
	idx int
	length int
	current_token markdownLexer.Token
	page string
}

// initializ a HTML render
// passing in a stream of tokens from the markdownLexer
func InitializeHtmlRender(stream []markdownLexer.Token) *HtmlRender {
	var render = new(HtmlRender)
	render.token_stream = stream
	render.idx = 0
	render.length = len(stream)
	render.page = ""
	render.current_token = render.token_stream[0]
	return render
}

// Renders a document in html from the markdown tokens provided
func (h *HtmlRender) RenderDocument() (string, error) {
	for h.current_token.Type != markdownLexer.EOF {
		str, err := h.render_token()
		if err != nil {
			return "", err
		}
		h.page += str
		if h.is_next_eof() {
			break
		}
		h.increment_token()
	}
	return h.page, nil
}

// get the token at a given index
func (h *HtmlRender) get_token_by_index(i int) (markdownLexer.Token) {
	return h.token_stream[i]
}

// increment the current token being looked at
func (h *HtmlRender) increment_token() bool {
	if h.idx >= h.length {
		return false
	}
	h.idx += 1
	h.current_token = h.token_stream[h.idx]
	return true
}

// look at the token previous to the current token
func (h *HtmlRender) peek_prev_token() (markdownLexer.Token, error) {
	if h.idx > 0 {
		return h.token_stream[h.idx-1], nil
	}
	return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Can not index an array less than 0")
}

// look at the token that comes after the current token
func (h *HtmlRender) peek_next_token() (markdownLexer.Token, error) {
	if h.idx < h.length-1 {
		return h.token_stream[h.idx+1], nil
	}
	return markdownLexer.Token{ Type: markdownLexer.INVALID }, errors.New("Can not look at tokens that don't exist beyond bounds")
}

// Looks ahead and search ig a token similar to the current or a newline or
// eof token comes up
// returns the token
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

// looks at the next token and returns whether or not is a newline token or an EOF token
func (h *HtmlRender) is_next_newline_or_eof() bool {
	token, err := h.peek_next_token()
	if err != nil {
		return false
	}
	if token.Type == markdownLexer.EOF {
		return true
	}
	if token.Type == markdownLexer.NEW_LINE {
		return true
	}
	return false
}

// looks at the next token and returns if it is a end of file token
func (h *HtmlRender) is_next_eof() bool {
	if h.idx < h.length-1 {
		if h.token_stream[h.idx+1].Type == markdownLexer.EOF {
			return true
		}
		return false
	}
	return true
}

// [desc](link) => <a href="link">Desc</a>
func (h *HtmlRender) render_link(link string, desc string) string {
	return "<a href=\"" + link + "\">" + desc + "</a>"
}

func (h *HtmlRender) get_tags(token markdownLexer.Token) (string, string, error) {
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
		return "<hr>", nil
	case markdownLexer.CHECKED:
		return "<input type=\"checkbox\" checked>", nil
	case markdownLexer.UNCHECKED:
		return "<input type=\"checkbox\" >", nil
	case markdownLexer.HEADER_ONE:
		return h.render_headers()
	case markdownLexer.HEADER_TWO:
		return h.render_headers()
	case markdownLexer.HEADER_THREE:
		return h.render_headers()
	case markdownLexer.ITALIC:
		return h.render_italic_token()
	case markdownLexer.BOLD:
		return h.render_bold_text()
	case markdownLexer.BOLD_ITALIC:
		return h.render_bold_italic_text()
	case markdownLexer.QUOTE:
		return h.render_quote()
	case markdownLexer.INLINE_CODE:
		return h.render_inline_code()
	case markdownLexer.CODE_BLOCK:
		return h.render_code_block()
	case markdownLexer.BULLET_MINUS:
		return h.render_bullet_points()
	case markdownLexer.BULLET_PLUS:
		return h.render_bullet_points()
	case markdownLexer.LEFT_BRACKET:
		return h.render_left_bracket()
	default:
		return "", errors.New("Issue rendering HTML")
	}
}

func (h *HtmlRender) render_italic_token() (string, error) {
	str := ""
	// Reference previous and next token to see it italic token is at the end of the line
	prev_token, err := h.peek_prev_token()
	if err != nil {
		return str, err
	}
	if prev_token.Type == markdownLexer.NEW_LINE {
		next_token, err := h.peek_next_token()
		if err != nil {
			return str, err
		}
		if next_token.Type == markdownLexer.WHITESPACE {
			// Since previous token is newline and next token is a white space
			// make the assumption that the asterick is actually meant to be
			// a bullt point
			start_tag, end_tag, err := h.get_tags(h.current_token)
			if err != nil {
				return "", err
			}
			str += start_tag
			for {
				if h.is_next_newline_or_eof() {
					break
				}
				h.increment_token()
				literal, err := h.render_token()
				if err != nil {
					return "", errors.New("Issue rendering bullet point")
				}
				str += literal
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
			start_tag, end_tag, err := h.get_tags(h.current_token)
			if err != nil {
				return "", err
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
	token, err := h.peek_to_newline_or_token()
	if err != nil {
		return "", err
	}
	// significant tokens to come up are the newline or end of file token
	// so just render the asterick as is in place
	if token.Type == markdownLexer.NEW_LINE || token.Type == markdownLexer.EOF {
		return h.current_token.Literal, nil
	}

	// next significant token is an italics
	// so render as expected for an italics tag
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Issue getting tags for rendering")
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

func (h *HtmlRender) render_quote() (string, error) {
	var str = ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Could not identify the correct tag for block quotes")
	}
	str += start_tag
	if h.is_next_newline_or_eof() {
		str += end_tag
		return str, nil
	}
	// look at the next token
	next_token, err := h.peek_next_token()
	if err != nil {
		return "", err
	}
	// if the next token is white space, we want to just consume that
	if next_token.Type == markdownLexer.WHITESPACE {
		h.increment_token()
	}
	h.increment_token()
	for {
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Unable to render a token")
		}
		str += literal
		if h.is_next_newline_or_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_inline_code() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Could not identify tag for code block tokens")
	}
	str += start_tag
	if h.is_next_newline_or_eof() {
		str += end_tag
		return str, nil
	}
	h.increment_token()
	for {
		if h.current_token.Type == markdownLexer.INLINE_CODE {
			break
		}
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering content within code blocks")
		}
		str += literal
		// go ahead and render if there is
		if h.is_next_newline_or_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_code_block() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Issue getting html tags for code blocks")
	}
	str += start_tag
	if h.is_next_eof() {
		str += end_tag
		return str, nil
	}
	h.increment_token()
	for {
		if h.current_token.Type == markdownLexer.CODE_BLOCK {
			break
		}
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering content in code blocks")
		}
		str += literal
		if h.is_next_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_bullet_points() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return "", errors.New("Could not identify the proper tag for bullet points")
	}
	str += start_tag
// 	if h.is_next_newline_or_eof() {
// 		str += end_tag
// 		return str, nil
// 	}
	next_token, err := h.peek_next_token()
	if err != nil {
		return "", err
	}
	if next_token.Type == markdownLexer.EOF || next_token.Type == markdownLexer.NEW_LINE {
		str += end_tag
		return str, nil
	}
	// if next token is a white space, consume it
	if next_token.Type == markdownLexer.WHITESPACE {
		h.increment_token()
	}
	h.increment_token()
	for {
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue render token in a bullet point")
		}
		str += literal
		if h.is_next_newline_or_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_headers() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Issue getting html tags for headers")
	}
	str += start_tag
	if h.is_next_newline_or_eof() {
		str += end_tag
		return str, nil
	}
	next_token, err := h.peek_next_token()
	if err != nil {
		return "", err
	}
	if next_token.Type != markdownLexer.WHITESPACE {
		return h.current_token.Literal, nil
	}
	// skip the white space
	h.increment_token()
	// not get to the start of the header
	h.increment_token()
	for {
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering text in header")
		}
		str += literal
		if h.is_next_newline_or_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_bold_text() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Issue getting html tags for bold")
	}
	str += start_tag
	if h.is_next_newline_or_eof() {
		str += end_tag
		return str, nil
	}
	h.increment_token()
	for {
		if h.current_token.Type == markdownLexer.BOLD {
			break
		}
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering content in bold block")
		}
		str += literal
		if h.is_next_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_bold_italic_text() (string, error) {
	str := ""
	start_tag, end_tag, err := h.get_tags(h.current_token)
	if err != nil {
		return str, errors.New("Issue getting html tags for bold")
	}
	str += start_tag
	if h.is_next_newline_or_eof() {
		str += end_tag
		return str, nil
	}
	h.increment_token()
	for {
		if h.current_token.Type == markdownLexer.BOLD_ITALIC {
			break
		}
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue rendering content in bold block")
		}
		str += literal
		if h.is_next_eof() {
			break
		}
		h.increment_token()
	}
	str += end_tag
	return str, nil
}

func (h *HtmlRender) render_left_bracket() (string, error) {
	// need to keep this just in case
	left_bracket := "["
	// between brackets
	content := ""
	// get content after right left bracket '[' and place in a temporary string
	// called content. This could be literal content or a part of a markdown link
	for {
		// check if next token is new line or eof
		if h.is_next_newline_or_eof() {
			// if so, just return the literal of the bracket and the content read so far
			return left_bracket + content, nil
		}
		// check if current token is right bracket
		if h.current_token.Type == markdownLexer.RIGHT_BRACKET {
			break
		}
		// increment the token
		h.increment_token()
		// render token
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue getting content after left bracket")
		}
		// add to the content for what comes after the left bracket
		content += literal
	}
	// current token is a right bracket at this point
	// save this just in case this is not a link
	right_bracket := "]"
	// if there is a new line or eof next, render what we have in brackets
	if h.is_next_newline_or_eof()  {
		return left_bracket + content + right_bracket, nil
	}
	next_token, err := h.peek_next_token()
	if err != nil {
		return "", err
	}
	if next_token.Type != markdownLexer.LEFT_PAREN {
		return left_bracket + content + right_bracket, nil
	}
	link := ""
	left_paren := "("
	// established that the next token can only be left parenthesis
	h.increment_token()
	// skip over that right parenthesis token
	for {
		// check if next token is eof or newline
		if h.is_next_newline_or_eof() {
			// return the literals of what we have so far
			return left_bracket + content + right_bracket + left_paren, nil
		}
		h.increment_token()
		// check for exit condition of right parenthsus
		if h.current_token.Type == markdownLexer.RIGHT_PAREN {
			// break since exit condition is met
			break
		}
		// get content inside of parenthesis
		literal, err := h.render_token()
		if err != nil {
			return "", errors.New("Issue trying to render possible link")
		}
		link += literal
	}
	// link present in link
	// hit right parenthsis token
	return h.render_link(link, content), nil
}
