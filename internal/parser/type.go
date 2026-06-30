package parser

func (p *Parser) parseType(stopNames bool) (Type, bool) {
	var typ Type
	for p.match("const") {
		typ.Const = true
	}
	if p.match("struct") {
		typ.Struct = true
		name, ok := p.ident()
		if !ok {
			return Type{}, false
		}
		typ.Name = name
	} else {
		name, ok := p.ident()
		if !ok {
			return Type{}, false
		}
		typ.Name = name
	}
	for {
		if p.match("*") {
			typ.Pointers++
			continue
		}
		break
	}
	return typ, true
}

func (t Type) IsVoid() bool {
	return t.Name == "void" && t.Pointers == 0
}
