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
	"../native/yoga/yoga/YGConfig.cpp",
	"../native/yoga/yoga/YGEnums.cpp",
	"../native/yoga/yoga/YGNode.cpp",
	"../native/yoga/yoga/YGNodeLayout.cpp",
	"../native/yoga/yoga/YGNodeStyle.cpp",
	"../native/yoga/yoga/YGPixelGrid.cpp",
	"../native/yoga/yoga/YGValue.cpp",
	"../native/yoga/yoga/algorithm/AbsoluteLayout.cpp",
	"../native/yoga/yoga/algorithm/Baseline.cpp",
	"../native/yoga/yoga/algorithm/Cache.cpp",
	"../native/yoga/yoga/algorithm/CalculateLayout.cpp",
	"../native/yoga/yoga/algorithm/FlexLine.cpp",
	"../native/yoga/yoga/algorithm/PixelGrid.cpp",
	"../native/yoga/yoga/config/Config.cpp",
	"../native/yoga/yoga/debug/AssertFatal.cpp",
	"../native/yoga/yoga/debug/Log.cpp",
	"../native/yoga/yoga/event/event.cpp",
	"../native/yoga/yoga/node/LayoutResults.cpp",
	"../native/yoga/yoga/node/Node.cpp",
}
