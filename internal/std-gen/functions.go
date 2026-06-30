package stdgen

import (
	"fmt"
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeFunctions(dir string, api parser.API) error {
	var w Writer
	writeHeader(&w)
	w.Line("import (")
	w.Line(`"math"`)
	w.Line(`yoga "github.com/husnulhamidiah/goyoga/gen"`)
	w.Line(")")
	w.Line("")
	for _, fn := range api.Functions {
		if !isTopLevelFunction(fn) || unsupportedFunction(fn) {
			continue
		}
		writeTopLevelFunction(&w, fn)
	}
	w.Line("func If[T any](expr bool, a, b T) T {")
	w.Line("if expr { return a }")
	w.Line("return b")
	w.Line("}")
	w.Line("")
	w.Line("func IsInf(f float32, sign int) bool {")
	w.Line("return math.IsInf(float64(f), sign)")
	w.Line("}")
	w.Line("")
	w.Line("func IsNaN(f float32) bool {")
	w.Line("return math.IsNaN(float64(f))")
	w.Line("}")
	w.Line("")
	return writeGoFile(dir, "functions_generated.go", &w)
}

func writeTopLevelFunction(w *Writer, fn parser.Function) {
	name := publicName(fn.Name)
	params, args := callParams(fn.Params, "")
	ret := stdType(fn.Return)
	if fn.Return.IsVoid() {
		w.Line("func %s(%s) {", name, strings.Join(params, ", "))
		w.Line("yoga.%s(%s)", name, strings.Join(args, ", "))
		w.Line("}")
		w.Line("")
		return
	}
	w.Line("func %s(%s) %s {", name, strings.Join(params, ", "), ret)
	w.Line("return %s", wrapReturn("yoga."+name+"("+strings.Join(args, ", ")+")", fn.Return))
	w.Line("}")
	w.Line("")
}

func isTopLevelFunction(fn parser.Function) bool {
	name := fn.Name
	if strings.HasSuffix(name, "ToString") {
		return false
	}
	return !strings.HasPrefix(name, "YGNode") && !strings.HasPrefix(name, "YGConfig")
}

func unsupportedFunction(fn parser.Function) bool {
	if unsupportedType(fn.Return) {
		return true
	}
	for _, param := range fn.Params {
		if unsupportedType(param.Type) {
			return true
		}
	}
	return false
}

func unsupportedType(t parser.Type) bool {
	if t.Name == "void" && t.Pointers > 0 {
		return true
	}
	if t.Name == "YGNodeRef" && t.Pointers > 0 {
		return true
	}
	return false
}

func callParams(params []parser.Param, receiver string) ([]string, []string) {
	out := make([]string, 0, len(params))
	args := make([]string, 0, len(params))
	for i, param := range params {
		name := safeIdent(param.Name)
		if name == "value" && param.Name == "" {
			name = fmt.Sprintf("arg%d", i)
		}
		if receiver != "" && i == 0 {
			args = append(args, receiver)
			continue
		}
		out = append(out, name+" "+stdType(param.Type))
		args = append(args, unwrapArg(name, param.Type))
	}
	return out, args
}

func unwrapArg(name string, t parser.Type) string {
	switch t.Name {
	case "YGNodeRef":
		return name + ".yogaRef()"
	case "YGNodeConstRef":
		return name + ".yogaConstRef()"
	case "YGConfigRef":
		return name + ".yogaRef()"
	case "YGConfigConstRef":
		return name + ".yogaConstRef()"
	}
	switch baseStdType(t) {
	case "Align", "BoxSizing", "Dimension", "Direction", "Display", "Edge", "Errata", "ExperimentalFeature",
		"FlexDirection", "GridTrackType", "Gutter", "Justify", "LogLevel", "MeasureMode", "NodeType",
		"Overflow", "PositionType", "Unit", "Wrap":
		return "yoga." + baseStdType(t) + "(" + name + ")"
	case "Value":
		return "yoga.Value(" + name + ")"
	default:
		return name
	}
}

func wrapReturn(expr string, t parser.Type) string {
	switch t.Name {
	case "YGNodeRef", "YGNodeConstRef":
		return "wrapNode(" + expr + ")"
	case "YGConfigRef":
		return "wrapConfig(" + expr + ")"
	case "YGConfigConstRef":
		return "wrapConstConfig(" + expr + ")"
	}
	switch baseStdType(t) {
	case "Align", "BoxSizing", "Dimension", "Direction", "Display", "Edge", "Errata", "ExperimentalFeature",
		"FlexDirection", "GridTrackType", "Gutter", "Justify", "LogLevel", "MeasureMode", "NodeType",
		"Overflow", "PositionType", "Unit", "Wrap", "Value":
		return baseStdType(t) + "(" + expr + ")"
	default:
		return expr
	}
}
