package parser

func (p *Parser) parseVariableAfterTypeAndName(typ Type, name string) {
	if typ.Name == "extern" {
		var ok bool
		typ, ok = p.parseType(false)
		if !ok {
			p.skipUntil(";")
			p.match(";")
			return
		}
		name, ok = p.ident()
		if !ok {
			p.skipUntil(";")
			p.match(";")
			return
		}
	}
	p.skipUntil(";")
	p.expect(";")
	p.api.Variables = append(p.api.Variables, Variable{Name: name, Type: typ})
}
