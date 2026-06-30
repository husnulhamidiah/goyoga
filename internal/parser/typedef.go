package parser

func (p *Parser) parseTypedef() {
	if p.match("enum") {
		p.parseEnum("")
		return
	}
	if p.match("struct") {
		name := ""
		if p.peek().Kind == TokenIdent {
			name = p.next().Text
		}
		if p.peek().Text == "{" {
			st, ok := p.parseStruct(name)
			if !ok {
				p.skipUntil(";")
				p.match(";")
				return
			}
			if alias, ok := p.ident(); ok {
				st.Name = alias
			}
			p.expect(";")
			p.api.Structs = append(p.api.Structs, st)
			p.api.Typedefs = append(p.api.Typedefs, Typedef{Name: st.Name, Type: Type{Name: st.Name, Struct: true}})
			return
		}
		typ := Type{Name: name, Struct: true}
		for p.match("*") {
			typ.Pointers++
		}
		alias, ok := p.ident()
		if ok {
			p.api.Typedefs = append(p.api.Typedefs, Typedef{Name: alias, Type: typ})
		}
		p.skipUntil(";")
		p.match(";")
		return
	}

	ret, ok := p.parseType(false)
	if !ok {
		p.skipUntil(";")
		p.match(";")
		return
	}
	if p.match("(") && p.match("*") {
		name, ok := p.ident()
		if !ok {
			p.skipUntil(";")
			p.match(";")
			return
		}
		p.expect(")")
		if !p.expect("(") {
			p.skipUntil(";")
			p.match(";")
			return
		}
		params := p.parseParams()
		p.expect(";")
		fn := Function{Name: name, Return: ret, Params: params}
		p.api.Typedefs = append(p.api.Typedefs, Typedef{Name: name, FunctionPointer: &fn})
		return
	}
	alias, ok := p.ident()
	if ok {
		p.api.Typedefs = append(p.api.Typedefs, Typedef{Name: alias, Type: ret})
	}
	p.skipUntil(";")
	p.match(";")
}
