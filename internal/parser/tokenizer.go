package parser

import "strings"

type Tokenizer struct {
	src  string
	pos  int
	line int
	col  int
}

func Tokenize(src string) []Token {
	t := &Tokenizer{src: src, line: 1, col: 1}
	var out []Token
	for {
		tok := t.next()
		out = append(out, tok)
		if tok.Kind == TokenEOF {
			return out
		}
	}
}

func (t *Tokenizer) next() Token {
	t.skipSpaceAndIgnored()
	if t.pos >= len(t.src) {
		return Token{Kind: TokenEOF, Line: t.line, Col: t.col}
	}

	line, col := t.line, t.col
	ch := t.peek()
	if isIdentStart(ch) {
		start := t.pos
		t.advance()
		for t.pos < len(t.src) && isIdentPart(t.peek()) {
			t.advance()
		}
		return Token{Kind: TokenIdent, Text: t.src[start:t.pos], Line: line, Col: col}
	}
	if isDigit(ch) {
		start := t.pos
		t.advance()
		for t.pos < len(t.src) && (isDigit(t.peek()) || t.peek() == '.') {
			t.advance()
		}
		return Token{Kind: TokenNumber, Text: t.src[start:t.pos], Line: line, Col: col}
	}
	if ch == '"' {
		start := t.pos
		t.advance()
		for t.pos < len(t.src) {
			c := t.peek()
			t.advance()
			if c == '\\' && t.pos < len(t.src) {
				t.advance()
				continue
			}
			if c == '"' {
				break
			}
		}
		return Token{Kind: TokenString, Text: t.src[start:t.pos], Line: line, Col: col}
	}

	t.advance()
	return Token{Kind: TokenSymbol, Text: string(ch), Line: line, Col: col}
}

func (t *Tokenizer) skipSpaceAndIgnored() {
	for t.pos < len(t.src) {
		if strings.HasPrefix(t.src[t.pos:], "//") {
			t.skipLine()
			continue
		}
		if strings.HasPrefix(t.src[t.pos:], "/*") {
			t.advance()
			t.advance()
			for t.pos < len(t.src) && !strings.HasPrefix(t.src[t.pos:], "*/") {
				t.advance()
			}
			if t.pos < len(t.src) {
				t.advance()
				t.advance()
			}
			continue
		}
		if t.peek() == '#' {
			t.skipLine()
			continue
		}
		if isSpace(t.peek()) {
			t.advance()
			continue
		}
		break
	}
}

func (t *Tokenizer) skipLine() {
	for t.pos < len(t.src) {
		continued := false
		for t.pos < len(t.src) && t.peek() != '\n' {
			continued = t.peek() == '\\'
			t.advance()
		}
		if t.pos < len(t.src) && t.peek() == '\n' {
			t.advance()
		}
		if !continued {
			return
		}
	}
}

func (t *Tokenizer) peek() byte {
	return t.src[t.pos]
}

func (t *Tokenizer) advance() {
	if t.src[t.pos] == '\n' {
		t.line++
		t.col = 1
	} else {
		t.col++
	}
	t.pos++
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func isIdentStart(ch byte) bool {
	return ch == '_' || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func isIdentPart(ch byte) bool {
	return isIdentStart(ch) || isDigit(ch)
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
