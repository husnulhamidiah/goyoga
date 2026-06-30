package parser

import "fmt"

type Parser struct {
	tokens []Token
	pos    int
	api    API
}

func Parse(src string) (API, error) {
	p := &Parser{tokens: Tokenize(src)}
	for !p.eof() {
		before := p.pos
		switch {
		case p.match("YG_ENUM_DECL"):
			p.parseYGEnumDecl()
		case p.match("typedef"):
			p.parseTypedef()
		case p.match("enum"):
			p.parseEnum("")
		case p.match("YG_EXPORT"):
			p.parseExport()
		default:
			p.next()
		}
		if p.pos == before {
			p.next()
		}
	}
	return p.api, nil
}

func (p *Parser) peek() Token {
	return p.tokens[p.pos]
}

func (p *Parser) next() Token {
	tok := p.tokens[p.pos]
	if p.pos < len(p.tokens)-1 {
		p.pos++
	}
	return tok
}

func (p *Parser) eof() bool {
	return p.peek().Kind == TokenEOF
}

func (p *Parser) match(text string) bool {
	if p.peek().Text != text {
		return false
	}
	p.next()
	return true
}

func (p *Parser) expect(text string) bool {
	return p.match(text)
}

func (p *Parser) ident() (string, bool) {
	if p.peek().Kind != TokenIdent {
		return "", false
	}
	return p.next().Text, true
}

func (p *Parser) skipUntil(texts ...string) {
	for !p.eof() {
		for _, text := range texts {
			if p.peek().Text == text {
				return
			}
		}
		p.next()
	}
}

func (p *Parser) skipBalanced(open, close string) {
	if !p.match(open) {
		return
	}
	depth := 1
	for !p.eof() && depth > 0 {
		switch p.next().Text {
		case open:
			depth++
		case close:
			depth--
		}
	}
}

func (p *Parser) unsupportedAt() error {
	tok := p.peek()
	return fmt.Errorf("unsupported token %q at %d:%d", tok.Text, tok.Line, tok.Col)
}
