package gen

import (
	"sort"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeVariables(dir string, api parser.API) error {
	var w Writer
	writePreamble(&w, false)

	vars := append([]parser.Variable(nil), api.Variables...)
	sort.Slice(vars, func(i, j int) bool { return vars[i].Name < vars[j].Name })
	for _, variable := range vars {
		w.Line("var %s %s = C.%s", publicName(variable.Name), goType(variable.Type), variable.Name)
	}
	return writeGoFile(dir, "variables.go", &w)
}
