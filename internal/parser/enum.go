package parser

func (p *Parser) parseYGEnumDecl() {
	if !p.expect("(") {
		return
	}
	name, ok := p.ident()
	if !ok {
		p.skipBalanced("(", ")")
		return
	}
	p.match(",")
	var values []EnumValue
	for !p.eof() && !p.match(")") {
		valueName, ok := p.ident()
		if !ok {
			p.next()
			continue
		}
		ev := EnumValue{Name: valueName}
		if p.match("=") {
			if p.peek().Kind == TokenNumber || p.peek().Kind == TokenIdent {
				ev.Value = p.next().Text
			}
		}
		values = append(values, ev)
		p.match(",")
	}
	p.api.Enums = append(p.api.Enums, Enum{Name: name, Values: values})
	p.api.Functions = append(p.api.Functions, Function{
		Name:   name + "ToString",
		Return: Type{Name: "char", Const: true, Pointers: 1},
		Params: []Param{{Name: "value", Type: Type{Name: name}}},
	})
}

func (p *Parser) parseEnum(name string) {
	if name == "" {
		name, _ = p.ident()
	}
	if !p.match("{") {
		p.skipUntil(";")
		p.match(";")
		return
	}
	var values []EnumValue
	for !p.eof() && !p.match("}") {
		valueName, ok := p.ident()
		if !ok {
			p.next()
			continue
		}
		ev := EnumValue{Name: valueName}
		if p.match("=") && (p.peek().Kind == TokenNumber || p.peek().Kind == TokenIdent) {
			ev.Value = p.next().Text
		}
		values = append(values, ev)
		p.match(",")
	}
	if alias, ok := p.ident(); ok {
		name = alias
	}
	p.expect(";")
	if name != "" {
		p.api.Enums = append(p.api.Enums, Enum{Name: name, Values: values})
	}
}
