package gen

import (
	"sort"
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeTypedefs(dir string, api parser.API) error {
	var w Writer
	writePreamble(&w, false)

	seen := map[string]bool{}
	var names []string
	for _, enum := range api.Enums {
		if !seen[enum.Name] {
			seen[enum.Name] = true
			names = append(names, enum.Name)
		}
	}
	for _, td := range api.Typedefs {
		if !seen[td.Name] {
			seen[td.Name] = true
			names = append(names, td.Name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		w.Line("type %s = C.%s", publicName(name), name)
	}
	return writeGoFile(dir, "typedefs.go", &w)
}

func publicName(name string) string {
	return strings.TrimPrefix(name, "YG")
}
