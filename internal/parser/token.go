package parser

type TokenKind int

const (
	TokenEOF TokenKind = iota
	TokenIdent
	TokenNumber
	TokenString
	TokenSymbol
)

type Token struct {
	Kind TokenKind
	Text string
	Line int
	Col  int
}

func (t Token) Is(text string) bool {
	return t.Text == text
}
