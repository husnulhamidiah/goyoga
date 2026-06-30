package stdgen

import "github.com/husnulhamidiah/goyoga/internal/parser"

func writeEnums(dir string, api parser.API) error {
	var w Writer
	writeHeader(&w)
	w.Line("import (")
	w.Line(`"fmt"`)
	w.Line(`"strings"`)
	w.Line(`yoga "github.com/husnulhamidiah/goyoga/gen"`)
	w.Line(")")
	w.Line("")
	for _, enum := range api.Enums {
		name := publicName(enum.Name)
		w.Line("type %s yoga.%s", name, name)
		w.Line("")
		if len(enum.Values) > 0 {
			w.Line("const (")
			for _, value := range enum.Values {
				w.Line("%s %s = %s(yoga.%s)", enumValueName(value.Name), name, name, enumValueName(value.Name))
			}
			w.Line(")")
			w.Line("")
		}
		w.Line("func (e %s) String() string {", name)
		w.Line("return yoga.%sToString(yoga.%s(e))", name, name)
		w.Line("}")
		w.Line("")
		w.Line("func %sFromString(s string) (%s, error) {", name, name)
		for _, value := range enum.Values {
			valueName := enumValueName(value.Name)
			w.Line("if strings.EqualFold(s, %s.String()) || strings.EqualFold(s, %q) {", valueName, valueName)
			w.Line("return %s, nil", valueName)
			w.Line("}")
		}
		w.Line("return %s(0), fmt.Errorf(\"unknown %s: %%q\", s)", name, name)
		w.Line("}")
		w.Line("")
	}
	return writeGoFile(dir, "enums_generated.go", &w)
}
