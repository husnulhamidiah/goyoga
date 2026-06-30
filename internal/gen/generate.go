package gen

import (
	"fmt"
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
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		if ext != ".go" && !(ext == ".cpp" && strings.HasPrefix(entry.Name(), "yoga_source_")) {
			continue
		}
		if err := os.Remove(filepath.Join(dir, entry.Name())); err != nil {
			return err
		}
	}
	if err := writeTypedefs(dir, api); err != nil {
		return err
	}
	if err := writeEnums(dir, api); err != nil {
		return err
	}
	if err := writeVariables(dir, api); err != nil {
		return err
	}
	if err := writeFunctions(dir, api); err != nil {
		return err
	}
	return writeYogaSources(dir)
}

func writeYogaSources(dir string) error {
	for i, source := range yogaSources {
		var w Writer
		w.Line(`#include "%s"`, source)
		name := fmt.Sprintf("yoga_source_%02d.cpp", i)
		if err := os.WriteFile(filepath.Join(dir, name), w.Bytes(), 0o644); err != nil {
			return err
		}
	}
	return nil
}

var yogaSources = []string{
	"../yoga/yoga/YGConfig.cpp",
	"../yoga/yoga/YGEnums.cpp",
	"../yoga/yoga/YGNode.cpp",
	"../yoga/yoga/YGNodeLayout.cpp",
	"../yoga/yoga/YGNodeStyle.cpp",
	"../yoga/yoga/YGPixelGrid.cpp",
	"../yoga/yoga/YGValue.cpp",
	"../yoga/yoga/algorithm/AbsoluteLayout.cpp",
	"../yoga/yoga/algorithm/Baseline.cpp",
	"../yoga/yoga/algorithm/Cache.cpp",
	"../yoga/yoga/algorithm/CalculateLayout.cpp",
	"../yoga/yoga/algorithm/FlexLine.cpp",
	"../yoga/yoga/algorithm/PixelGrid.cpp",
	"../yoga/yoga/config/Config.cpp",
	"../yoga/yoga/debug/AssertFatal.cpp",
	"../yoga/yoga/debug/Log.cpp",
	"../yoga/yoga/event/event.cpp",
	"../yoga/yoga/node/LayoutResults.cpp",
	"../yoga/yoga/node/Node.cpp",
}
