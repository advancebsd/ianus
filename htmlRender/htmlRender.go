package htmlRender

import (
	"errors"

	markdownLexer "github.com/advancebsd/ianus/markdownLexer"
)

type HtmlRender struct {
	token_stream  []markdownLexer.Token
	idx           int
	length        int
	current_token markdownLexer.Token
	page          string
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
func (h *HtmlRender) get_token_by_index(i int) markdownLexer.Token {
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
	return markdownLexer.Token{Type: markdownLexer.INVALID}, errors.New("Can not index an array less than 0")
}

// look at the token that comes after the current token
func (h *HtmlRender) peek_next_token() (markdownLexer.Token, error) {
	if h.idx < h.length-1 {
		return h.token_stream[h.idx+1], nil
	}
	return markdownLexer.Token{Type: markdownLexer.INVALID}, errors.New("Can not look at tokens that don't exist beyond bounds")
}

// Looks ahead and search ig a token similar to the current or a newline or
// eof token comes up
// returns the token
func (h *HtmlRender) peek_to_newline_or_token() (markdownLexer.Token, error) {
	if h.idx >= h.length-1 {
		return markdownLexer.Token{Type: markdownLexer.INVALID}, errors.New("Can not peek beyond array bounds")
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

func (h *HtmlRender) render_to_end_of_line() string {
	str := ""
	for {
		if h.current_token.Type == markdownLexer.NEW_LINE {
			return str
		}
		if h.current_token.Type == markdownLexer.EOF {
			return str
		}
		literal, err := h.render_token()
		if err != nil {
			continue
		}
		str += literal
		next_token, err := h.peek_next_token()
		if err != nil {
			continue
		}
		if next_token.Type == markdownLexer.NEW_LINE {
			break
		}
	}
	return str
}

func (h *HtmlRender) check_if_asterick_bullet() bool {
	prev_token, err := h.peek_next_token()
	if err != nil {
		// do nothing
	}
	if prev_token.Type != markdownLexer.NEW_LINE {
		return false
	}
	next_token, err := h.peek_next_token()
	if err != nil {
		return false
	}
	if next_token.Type != markdownLexer.WHITESPACE {
		return false
	}
	return true
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
		return "<ul>", "</ul>", nil
	case markdownLexer.BULLET_PLUS:
		return "<ul>", "</ul>", nil
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
	if h.check_if_asterick_bullet() {
		start_tag, end_tag, err := h.get_tags(markdownLexer.Token{
			Type: markdownLexer.NEW_LINE,
		})
		if err != nil {
			return "", err
		}
		str := ""
		str += start_tag
		next_token, err := h.peek_next_token()
		if err != nil {
			str += end_tag
			return str, err
		}
		if next_token.Type == markdownLexer.NEW_LINE || next_token.Type == markdownLexer.EOF {
			str += end_tag
			return str, nil
		}
		for {
			h.increment_token()
			literal, err := h.render_token()
			if err != nil {
				return "", nil
			}
			str += literal
			next_token, err := h.peek_next_token()
			if err != nil {
				return "", nil
			}
			if next_token.Type == markdownLexer.NEW_LINE || next_token.Type == markdownLexer.EOF {
				break
			}
		}
		str += end_tag
		return str, nil
	}
	// render italic
	str := ""
	start_tag, end_tag, err := h.get_tags((h.current_token))
	if err != nil {
		return "", errors.New("Issue getting appropraite tags for italics")
	}
	str += start_tag
	next_token, err := h.peek_next_token()
	if err != nil {
		str += end_tag
		return str, errors.New("Issue peeking ahead of next token in italics")
	}
	if next_token.Type == markdownLexer.EOF || next_token.Type == markdownLexer.NEW_LINE {
		str += end_tag
		return str, nil
	}
	for {
		h.increment_token()
		if h.current_token.Type == markdownLexer.ITALIC {
			break
		}
		literal, err := h.render_token()
		if err != nil {
			return "", err
		}
		str += literal
		next_token, err := h.peek_next_token()
		if err != nil {
			return "", err
		}
		if next_token.Type == markdownLexer.EOF || next_token.Type == markdownLexer.NEW_LINE {
			break
		}
	}
	str += end_tag
	return str, nil
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
	if h.check_if_asterick_bullet() == false {
		return h.current_token.Literal, nil
	}
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

func (h *HtmlRender) render_between(t markdownLexer.TokenType) (string, error) {
	h.increment_token()
	str := ""
	for h.current_token.Type != t {
		literal, err := h.render_token()
		if err != nil {
			return "", err
		}
		str += literal
		if h.is_next_newline_or_eof() {
			return str, errors.New("No matching ending token")
		}
		h.increment_token()
	}
	return str, nil
}

func (h *HtmlRender) render_left_bracket() (string, error) {
	if h.is_next_newline_or_eof() {
		return "[", nil
	}
	between_brackets, err := h.render_between(markdownLexer.RIGHT_BRACKET)
	if err != nil {
		return "[" + between_brackets, nil
	}
	next_token, err := h.peek_next_token()
	if err != nil {
		return "", err
	}
	if next_token.Type != markdownLexer.LEFT_PAREN {
		return "[" + between_brackets + "]", nil
	}
	if next_token.Type == markdownLexer.NEW_LINE || next_token.Type == markdownLexer.EOF {
		return "[" + between_brackets + "]", nil
	}
	h.increment_token()
	between_parens, err := h.render_between(markdownLexer.RIGHT_PAREN)
	if err != nil {
		return "[" + between_brackets + "]" + "(" + between_parens, nil
	}
	return h.render_link(between_parens, between_brackets), nil

}
