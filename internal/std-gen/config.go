package stdgen

import (
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeConfig(dir string, api parser.API) error {
	var w Writer
	writeHeader(&w)
	w.Line(`import yoga "github.com/husnulhamidiah/goyoga/gen"`)
	w.Line("")
	for _, fn := range api.Functions {
		if !strings.HasPrefix(fn.Name, "YGConfig") || unsupportedFunction(fn) {
			continue
		}
		switch fn.Name {
		case "YGConfigNew":
			w.Line("func NewConfig() *Config {")
			w.Line("return wrapConfig(yoga.ConfigNew())")
			w.Line("}")
			w.Line("")
		case "YGConfigFree":
			w.Line("func (c *Config) Free() {")
			w.Line("if c == nil || c.ref == nil { return }")
			w.Line("yoga.ConfigFree(c.ref)")
			w.Line("c.ref = nil")
			w.Line("c.constRef = nil")
			w.Line("}")
			w.Line("")
			w.Line("func (c *Config) Destroy() {")
			w.Line("c.Free()")
			w.Line("}")
			w.Line("")
		case "YGConfigGetDefault":
			w.Line("func DefaultConfig() *Config {")
			w.Line("return wrapConstConfig(yoga.ConfigGetDefault())")
			w.Line("}")
			w.Line("")
		default:
			writeConfigMethod(&w, fn)
		}
	}
	return writeGoFile(dir, "config_generated.go", &w)
}

func writeConfigMethod(w *Writer, fn parser.Function) {
	if len(fn.Params) == 0 {
		return
	}
	receiver := "c.yogaRef()"
	if fn.Params[0].Type.Name == "YGConfigConstRef" {
		receiver = "c.yogaConstRef()"
	}
	method := strings.TrimPrefix(publicName(fn.Name), "Config")
	if strings.HasPrefix(method, "Get") {
		method = strings.TrimPrefix(method, "Get")
	}
	params, args := callParams(fn.Params, receiver)
	ret := stdType(fn.Return)
	if fn.Return.IsVoid() {
		w.Line("func (c *Config) %s(%s) {", method, strings.Join(params, ", "))
		w.Line("if c == nil { return }")
		w.Line("yoga.%s(%s)", publicName(fn.Name), strings.Join(args, ", "))
		w.Line("}")
		w.Line("")
		return
	}
	w.Line("func (c *Config) %s(%s) %s {", method, strings.Join(params, ", "), ret)
	w.Line("if c == nil { return %s }", zeroValue(fn.Return))
	w.Line("return %s", wrapReturn("yoga."+publicName(fn.Name)+"("+strings.Join(args, ", ")+")", fn.Return))
	w.Line("}")
	w.Line("")
}
