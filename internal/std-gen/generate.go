package stdgen

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/husnulhamidiah/goyoga/internal/parser"
)

func Generate(api parser.API, dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), "_generated.go") {
			continue
		}
		if err := os.Remove(filepath.Join(dir, entry.Name())); err != nil {
			return err
		}
	}
	if err := writeTypes(dir, api); err != nil {
		return err
	}
	if err := writeEnums(dir, api); err != nil {
		return err
	}
	if err := writeConfig(dir, api); err != nil {
		return err
	}
	if err := writeNode(dir, api); err != nil {
		return err
	}
	return writeFunctions(dir, api)
}
