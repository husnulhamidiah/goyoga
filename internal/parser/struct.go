package parser

func (p *Parser) parseStruct(name string) (Struct, bool) {
	if name == "" {
		name, _ = p.ident()
	}
	if !p.match("{") {
		return Struct{Name: name}, false
	}
	var fields []Field
	for !p.eof() && !p.match("}") {
		typ, ok := p.parseType(false)
		if !ok {
			p.skipUntil(";")
			p.match(";")
			continue
		}
		fieldName, ok := p.ident()
		if !ok {
			p.skipUntil(";")
			p.match(";")
			continue
		}
		p.skipUntil(";")
		p.match(";")
		fields = append(fields, Field{Name: fieldName, Type: typ})
	}
	return Struct{Name: name, Fields: fields}, true
}
