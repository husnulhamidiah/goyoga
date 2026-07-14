package gen

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

func writePreamble(w *Writer, needsUnsafe bool) {
	w.Line("package yoga")
	w.Line("")
	w.Line("/*")
	w.Line("#cgo CFLAGS: -I${SRCDIR}/../native/yoga")
	w.Line("#cgo CXXFLAGS: -I${SRCDIR}/../native/yoga -std=c++20")
	w.Line("#cgo darwin LDFLAGS: -lc++")
	w.Line("#cgo linux LDFLAGS: -lstdc++")
	w.Line("#include <yoga/Yoga.h>")
	w.Line("*/")
	w.Line(`import "C"`)
	if needsUnsafe {
		w.Line(`import "unsafe"`)
	}
	w.Line("")
}
