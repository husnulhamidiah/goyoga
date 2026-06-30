package stdgen

import (
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func publicName(name string) string {
	return strings.TrimPrefix(name, "YG")
}

func enumValueName(name string) string {
	return strings.TrimPrefix(name, "YG")
}

func stdType(t parser.Type) string {
	if t.Name == "char" && t.Pointers == 1 {
		return "string"
	}
	if t.Name == "void" && t.Pointers > 0 {
		return "unsafe.Pointer"
	}
	if t.Pointers > 0 {
		switch t.Name {
		case "YGNode":
			return "*Node"
		case "YGConfig":
			return "*Config"
		}
		return strings.Repeat("*", t.Pointers) + baseStdType(t)
	}
	return baseStdType(t)
}

func baseStdType(t parser.Type) string {
	switch t.Name {
	case "YGNodeRef", "YGNodeConstRef":
		return "*Node"
	case "YGConfigRef", "YGConfigConstRef":
		return "*Config"
	}
	switch t.Name {
	case "void":
		return ""
	case "bool":
		return "bool"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "int":
		return "int"
	case "uint32_t":
		return "uint32"
	case "size_t":
		return "uint"
	case "char":
		return "byte"
	default:
		return publicName(t.Name)
	}
}

func zeroValue(t parser.Type) string {
	switch t.Name {
	case "YGNodeRef", "YGNodeConstRef", "YGConfigRef", "YGConfigConstRef":
		return "nil"
	}
	if strings.HasSuffix(stdType(t), "Func") || stdType(t) == "Logger" {
		return "nil"
	}
	switch stdType(t) {
	case "bool":
		return "false"
	case "float32", "float64", "int", "uint32", "uint":
		return "0"
	case "string":
		return `""`
	case "*Node", "*Config":
		return "nil"
	case "Value", "Size", "Layout":
		return stdType(t) + "{}"
	default:
		if t.IsVoid() {
			return ""
		}
		return stdType(t) + "(0)"
	}
}

func safeIdent(name string) string {
	if name == "" {
		return "value"
	}
	switch name {
	case "break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var":
		return name + "Value"
	default:
		return name
	}
}
