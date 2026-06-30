# goyoga

Go bindings for Facebook Yoga, generated from Yoga's public C headers.

This project uses a tiny handwritten parser for the narrow declaration shapes
Yoga exposes. It does not use libclang, SWIG, c-for-go, parser generators, or a
general C parser.

## Requirements

- Go 1.25+
- A C/C++ toolchain usable by cgo
- The `yoga/` submodule checked out

If the submodule is missing:

```sh
git submodule update --init --recursive
```

## Generate Bindings

Generated files are written to `gen/`.

```sh
go run ./cmd/gen
```

The generator reads `yoga/yoga/Yoga.h`, follows Yoga public includes, parses the
supported public C API declarations, and rewrites the generated Go files.

## Build And Test

```sh
go test ./...
```

Yoga may emit deprecation warnings from public APIs that are still generated.
Those warnings are harmless as long as the build succeeds.

## Use

Import the generated package:

```go
package main

import yoga "github.com/husnulhamidiah/goyoga/gen"

func main() {
	node := yoga.NodeNew()
	defer yoga.NodeFree(node)

	yoga.NodeStyleSetWidth(node, 100)
	yoga.NodeStyleSetHeight(node, 50)
	yoga.NodeCalculateLayout(node, 300, 300, yoga.DirectionLTR)

	println(yoga.NodeLayoutGetWidth(node))
	println(yoga.NodeLayoutGetHeight(node))
}
```

The Go names mostly mirror Yoga C names with the `YG` prefix removed, for
example `YGNodeNew` becomes `NodeNew`, and `YGDirectionLTR` becomes
`DirectionLTR`.

## Project Layout

- `cmd/gen`: generator command
- `internal/parser`: Yoga-specific tokenizer and parser
- `internal/gen`: Go binding writer
- `gen`: generated cgo bindings
- `yoga`: Facebook Yoga submodule
