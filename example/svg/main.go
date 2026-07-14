package main

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	goyoga "github.com/husnulhamidiah/goyoga/std"
)

type box struct {
	node     *goyoga.Node
	label    string
	children []*box
}

func main() {
	config := goyoga.NewConfig()
	defer config.Free()
	config.SetUseWebDefaults(true)

	root := newBox(config, "root")
	defer root.node.FreeRecursive()
	root.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	root.node.SetWidth(300)
	root.node.SetHeight(600)

	header := newBox(config, "header")
	header.node.SetPadding(goyoga.EdgeAll, 15)
	header.node.SetJustifyContent(goyoga.JustifySpaceBetween)
	header.node.SetAlignItems(goyoga.AlignCenter)
	for i := 1; i <= 3; i++ {
		header.addChild(fixedBox(config, fmt.Sprintf("top-%d", i), 50, 25))
	}

	content := newBox(config, "content")
	content.node.SetFlexGrow(1)
	content.node.SetPadding(goyoga.EdgeAll, 20)
	content.node.SetGap(goyoga.GutterAll, 5)
	content.node.SetJustifyContent(goyoga.JustifyCenter)
	content.node.SetFlexWrap(goyoga.WrapWrap)
	content.node.SetAlignContent(goyoga.AlignFlexStart)
	for i := 1; i <= 16; i++ {
		content.addChild(fixedBox(config, fmt.Sprintf("item-%02d", i), 50, 25))
	}

	footer := newBox(config, "footer")
	footer.node.SetPadding(goyoga.EdgeAll, 15)
	footer.node.SetJustifyContent(goyoga.JustifyCenter)
	footer.node.SetAlignItems(goyoga.AlignCenter)
	footer.node.SetFlexDirection(goyoga.FlexDirectionColumn)
	footer.addChild(fixedBox(config, "dot", 20, 20))
	footer.addChild(fixedBox(config, "bar", 100, 10))

	root.addChild(header)
	root.addChild(content)
	root.addChild(footer)
	root.node.CalculateLayout(300, 600, goyoga.DirectionLTR)

	out := filepath.Join("example", "svg", "layout.svg")
	if err := os.WriteFile(out, []byte(renderSVG(root)), 0o644); err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func newBox(config *goyoga.Config, label string) *box {
	return &box{node: goyoga.NewNodeWithConfig(config), label: label}
}

func fixedBox(config *goyoga.Config, label string, width, height float32) *box {
	b := newBox(config, label)
	b.node.SetWidth(width)
	b.node.SetHeight(height)
	return b
}

func (b *box) addChild(child *box) {
	b.node.InsertChild(child.node, uint(len(b.children)))
	b.children = append(b.children, child)
}

func renderSVG(root *box) string {
	width := root.node.GetComputedWidth()
	height := root.node.GetComputedHeight()

	var out strings.Builder
	out.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%.0f" height="%.0f" viewBox="0 0 %.0f %.0f">`, width, height, width, height))
	out.WriteByte('\n')
	out.WriteString(`<rect width="100%" height="100%" fill="#f8fafc"/>`)
	out.WriteByte('\n')
	renderBox(&out, root, 0, 0, 0)
	out.WriteString("</svg>\n")
	return out.String()
}

func renderBox(out *strings.Builder, b *box, parentX, parentY float32, depth int) {
	x := parentX + b.node.GetComputedLeft()
	y := parentY + b.node.GetComputedTop()
	w := b.node.GetComputedWidth()
	h := b.node.GetComputedHeight()
	fill := []string{"#e0f2fe", "#dcfce7", "#fef3c7", "#fce7f3"}[depth%4]

	out.WriteString(fmt.Sprintf(`<rect x="%.1f" y="%.1f" width="%.1f" height="%.1f" rx="4" fill="%s" stroke="#334155" stroke-width="1"/>`, x, y, w, h, fill))
	out.WriteByte('\n')
	if w >= 36 && h >= 16 {
		out.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="ui-sans-serif, system-ui, sans-serif" font-size="9" fill="#0f172a">%s</text>`, x+6, y+13, html.EscapeString(b.label)))
		out.WriteByte('\n')
	}
	for _, child := range b.children {
		renderBox(out, child, x, y, depth+1)
	}
}
