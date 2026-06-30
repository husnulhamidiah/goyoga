package stdgen

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

type Writer struct {
	bytes.Buffer
}

func (w *Writer) Line(format string, args ...any) {
	if len(args) == 0 {
		w.WriteString(format)
	} else {
		fmt.Fprintf(w, format, args...)
	}
	w.WriteByte('\n')
}

func writeGoFile(dir, name string, w *Writer) error {
	src, err := format.Source(w.Bytes())
	if err != nil {
		src = w.Bytes()
	}
	return os.WriteFile(filepath.Join(dir, name), src, 0o644)
}

func writeHeader(w *Writer) {
	w.Line("package goyoga")
	w.Line("")
}
