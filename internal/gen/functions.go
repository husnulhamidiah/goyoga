package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeFunctions(dir string, api parser.API) error {
	var w Writer
	writePreamble(&w, true)

	funcs := append([]parser.Function(nil), api.Functions...)
	sort.Slice(funcs, func(i, j int) bool { return funcs[i].Name < funcs[j].Name })
	for _, fn := range funcs {
		writeFunction(&w, fn)
	}
	return writeGoFile(dir, "functions.go", &w)
}

func writeFunction(w *Writer, fn parser.Function) {
	params := make([]string, 0, len(fn.Params))
	args := make([]string, 0, len(fn.Params))
	for i, param := range fn.Params {
		name := param.Name
		if name == "" {
			name = fmt.Sprintf("arg%d", i)
		}
		name = safeIdent(name)
		params = append(params, fmt.Sprintf("%s %s", name, goType(param.Type)))
		args = append(args, cArg(name, param.Type))
	}

	ret := goType(fn.Return)
	if fn.Return.IsVoid() {
		w.Line("func %s(%s) {", publicName(fn.Name), strings.Join(params, ", "))
		w.Line("C.%s(%s)", fn.Name, strings.Join(args, ", "))
		w.Line("}")
		w.Line("")
		return
	}
	w.Line("func %s(%s) %s {", publicName(fn.Name), strings.Join(params, ", "), ret)
	w.Line("return %s", goReturn(fmt.Sprintf("C.%s(%s)", fn.Name, strings.Join(args, ", ")), fn.Return))
	w.Line("}")
	w.Line("")
}

func safeIdent(name string) string {
	switch name {
	case "break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var":
		return name + "Value"
	default:
		return name
	}
}

func goType(t parser.Type) string {
	if t.Name == "char" && t.Pointers == 1 {
		return "string"
	}
	if t.Name == "void" && t.Pointers > 0 {
		return "unsafe.Pointer"
	}
	base := baseGoType(t)
	if t.Pointers == 0 {
		return base
	}
	return strings.Repeat("*", t.Pointers) + base
}

func baseGoType(t parser.Type) string {
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

func cArg(name string, t parser.Type) string {
	if t.Name == "void" && t.Pointers > 0 {
		return name
	}
	switch t.Name {
	case "bool":
		return "C.bool(" + name + ")"
	case "float":
		return "C.float(" + name + ")"
	case "double":
		return "C.double(" + name + ")"
	case "int":
		return "C.int(" + name + ")"
	case "uint32_t":
		return "C.uint32_t(" + name + ")"
	case "size_t":
		return "C.size_t(" + name + ")"
	}
	if t.Pointers > 0 {
		return fmt.Sprintf("(%s)(%s)", cPointerType(t), name)
	}
	return name
}

func goReturn(expr string, t parser.Type) string {
	if t.Name == "char" && t.Pointers == 1 {
		return "C.GoString(" + expr + ")"
	}
	switch t.Name {
	case "bool":
		return "bool(" + expr + ")"
	case "float":
		return "float32(" + expr + ")"
	case "double":
		return "float64(" + expr + ")"
	case "int":
		return "int(" + expr + ")"
	case "uint32_t":
		return "uint32(" + expr + ")"
	case "size_t":
		return "uint(" + expr + ")"
	default:
		return expr
	}
}

func cPointerType(t parser.Type) string {
	base := "C." + t.Name
	if t.Struct {
		base = "C.struct_" + t.Name
	}
	return strings.Repeat("*", t.Pointers) + base
}
