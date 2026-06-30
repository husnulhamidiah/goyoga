package stdgen

import (
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeNode(dir string, api parser.API) error {
	var w Writer
	writeHeader(&w)
	w.Line(`import yoga "github.com/husnulhamidiah/goyoga/gen"`)
	w.Line("")
	for _, fn := range api.Functions {
		if !strings.HasPrefix(fn.Name, "YGNode") || unsupportedFunction(fn) {
			continue
		}
		if len(fn.Params) == 0 || (fn.Params[0].Type.Name != "YGNodeRef" && fn.Params[0].Type.Name != "YGNodeConstRef") {
			continue
		}
		if strings.HasSuffix(fn.Name, "ToString") {
			continue
		}
		switch fn.Name {
		case "YGNodeNew":
			w.Line("func NewNode() *Node {")
			w.Line("return wrapNode(yoga.NodeNew())")
			w.Line("}")
			w.Line("")
		case "YGNodeNewWithConfig":
			w.Line("func NewNodeWithConfig(config *Config) *Node {")
			w.Line("return wrapNode(yoga.NodeNewWithConfig(config.yogaConstRef()))")
			w.Line("}")
			w.Line("")
		case "YGNodeFree":
			writeNodeFree(&w, "Free", "NodeFree")
			w.Line("func (n *Node) Destroy() {")
			w.Line("n.Free()")
			w.Line("}")
			w.Line("")
		case "YGNodeFreeRecursive":
			writeNodeFree(&w, "FreeRecursive", "NodeFreeRecursive")
		default:
			writeNodeMethod(&w, fn)
		}
	}
	writeComputedLayout(&w)
	return writeGoFile(dir, "node_generated.go", &w)
}

func writeNodeMethod(w *Writer, fn parser.Function) {
	if len(fn.Params) == 0 {
		return
	}
	receiver := "n.yogaRef()"
	if fn.Params[0].Type.Name == "YGNodeConstRef" {
		receiver = "n.yogaConstRef()"
	}
	method := nodeMethodName(fn.Name)
	params, args := callParams(fn.Params, receiver)
	ret := stdType(fn.Return)
	if fn.Return.IsVoid() {
		w.Line("func (n *Node) %s(%s) {", method, strings.Join(params, ", "))
		w.Line("if n == nil { return }")
		w.Line("yoga.%s(%s)", publicName(fn.Name), strings.Join(args, ", "))
		w.Line("}")
		w.Line("")
		return
	}
	w.Line("func (n *Node) %s(%s) %s {", method, strings.Join(params, ", "), ret)
	w.Line("if n == nil { return %s }", zeroValue(fn.Return))
	w.Line("return %s", wrapReturn("yoga."+publicName(fn.Name)+"("+strings.Join(args, ", ")+")", fn.Return))
	w.Line("}")
	w.Line("")
}

func writeNodeFree(w *Writer, method, lowLevel string) {
	w.Line("func (n *Node) %s() {", method)
	w.Line("if n == nil || n.ref == nil { return }")
	w.Line("yoga.%s(n.ref)", lowLevel)
	w.Line("n.ref = nil")
	w.Line("}")
	w.Line("")
}

func nodeMethodName(name string) string {
	name = publicName(name)
	switch {
	case strings.HasPrefix(name, "NodeStyle"):
		return strings.TrimPrefix(name, "NodeStyle")
	case strings.HasPrefix(name, "NodeLayoutGetRaw"):
		return "GetRaw" + strings.TrimPrefix(name, "NodeLayoutGetRaw")
	case strings.HasPrefix(name, "NodeLayoutGetHadOverflow"):
		return "HadOverflow"
	case strings.HasPrefix(name, "NodeLayoutGetDirection"):
		return "GetLayoutDirection"
	case strings.HasPrefix(name, "NodeLayoutGet"):
		return "GetComputed" + strings.TrimPrefix(name, "NodeLayoutGet")
	case strings.HasPrefix(name, "NodeGetHasNewLayout"):
		return "HasNewLayout"
	case strings.HasPrefix(name, "NodeGet"):
		return "Get" + strings.TrimPrefix(name, "NodeGet")
	default:
		return strings.TrimPrefix(name, "Node")
	}
}

func writeComputedLayout(w *Writer) {
	w.Line("func (n *Node) GetComputedLayout() Layout {")
	w.Line("return Layout{")
	w.Line("Left: n.GetComputedLeft(),")
	w.Line("Top: n.GetComputedTop(),")
	w.Line("Right: n.GetComputedRight(),")
	w.Line("Bottom: n.GetComputedBottom(),")
	w.Line("Width: n.GetComputedWidth(),")
	w.Line("Height: n.GetComputedHeight(),")
	w.Line("Direction: n.GetLayoutDirection(),")
	w.Line("HadOverflow: n.HadOverflow(),")
	w.Line("}")
	w.Line("}")
	w.Line("")
}
