package gen

import (
	"sort"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func writeEnums(dir string, api parser.API) error {
	var w Writer
	writePreamble(&w, false)

	enums := append([]parser.Enum(nil), api.Enums...)
	sort.Slice(enums, func(i, j int) bool { return enums[i].Name < enums[j].Name })
	for _, enum := range enums {
		if len(enum.Values) == 0 {
			continue
		}
		w.Line("const (")
		for _, value := range enum.Values {
			w.Line("%s %s = C.%s", publicName(value.Name), publicName(enum.Name), value.Name)
		}
		w.Line(")")
		w.Line("")
	}
	return writeGoFile(dir, "enums.go", &w)
}
