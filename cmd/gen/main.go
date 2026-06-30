package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	bindgen "github.com/husnulhamidiah/goyoga/internal/gen"
	"github.com/husnulhamidiah/goyoga/internal/parser"
)

var includeRE = regexp.MustCompile(`^\s*#include\s+<yoga/([^>]+)>`)

func main() {
	header := flag.String("header", "yoga/yoga/Yoga.h", "Yoga umbrella header")
	out := flag.String("out", "gen", "output directory")
	flag.Parse()

	src, err := collectHeaders(*header)
	if err != nil {
		fatal(err)
	}
	api, err := parser.Parse(string(src))
	if err != nil {
		fatal(err)
	}
	if err := bindgen.Generate(api, *out); err != nil {
		fatal(err)
	}
	fmt.Printf("generated %d enums, %d typedefs, %d structs, %d variables, %d functions\n",
		len(api.Enums), len(api.Typedefs), len(api.Structs), len(api.Variables), len(api.Functions))
}

func collectHeaders(path string) ([]byte, error) {
	seen := map[string]bool{}
	var out bytes.Buffer
	if err := collectOne(filepath.Clean(path), seen, &out); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func collectOne(path string, seen map[string]bool, out *bytes.Buffer) error {
	if seen[path] {
		return nil
	}
	seen[path] = true

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if match := includeRE.FindStringSubmatch(line); match != nil {
			includePath := filepath.Join("yoga", "yoga", match[1])
			if err := collectOne(includePath, seen, out); err != nil {
				return err
			}
			continue
		}
		out.WriteString(line)
		out.WriteByte('\n')
	}
	return scanner.Err()
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
