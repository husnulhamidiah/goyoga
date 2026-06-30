package parser

func (p *Parser) parseExport() {
	for p.peek().Text == "YG_DEPRECATED" {
		p.skipBalanced("(", ")")
	}
	typ, ok := p.parseType(false)
	if !ok {
		p.skipUntil(";")
		p.match(";")
		return
	}
	name, ok := p.ident()
	if !ok {
		p.skipUntil(";")
		p.match(";")
		return
	}
	if p.match("(") {
		params := p.parseParams()
		p.expect(";")
		p.api.Functions = append(p.api.Functions, Function{Name: name, Return: typ, Params: params})
		return
	}
	p.parseVariableAfterTypeAndName(typ, name)
}

func (p *Parser) parseParams() []Param {
	if p.match("void") {
		p.expect(")")
		return nil
	}
	var params []Param
	for !p.eof() && !p.match(")") {
		typ, ok := p.parseType(false)
		if !ok {
			p.skipUntil(",", ")")
			p.match(",")
			continue
		}
		name := ""
		if p.peek().Kind == TokenIdent {
			name = p.next().Text
		}
		params = append(params, Param{Name: name, Type: typ})
		if !p.match(",") {
			p.expect(")")
			break
		}
	}
	return params
}
