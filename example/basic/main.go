package main

import (
	"fmt"

	goyoga "github.com/husnulhamidiah/goyoga/std"
)

func main() {
	root := goyoga.NewNode()
	defer root.FreeRecursive()

	root.SetWidth(300)
	root.SetHeight(200)
	root.SetFlexDirection(goyoga.FlexDirectionRow)
	root.SetAlignItems(goyoga.AlignCenter)

	left := goyoga.NewNode()
	left.SetWidth(100)
	left.SetHeight(50)

	right := goyoga.NewNode()
	right.SetWidth(75)
	right.SetHeight(80)

	root.InsertChild(left, 0)
	root.InsertChild(right, 1)
	root.CalculateLayout(300, 200, goyoga.DirectionLTR)

	fmt.Printf("root:  %.0fx%.0f\n", root.GetComputedWidth(), root.GetComputedHeight())
	fmt.Printf("left:  x=%.0f y=%.0f %.0fx%.0f\n",
		left.GetComputedLeft(),
		left.GetComputedTop(),
		left.GetComputedWidth(),
		left.GetComputedHeight(),
	)
	fmt.Printf("right: x=%.0f y=%.0f %.0fx%.0f\n",
		right.GetComputedLeft(),
		right.GetComputedTop(),
		right.GetComputedWidth(),
		right.GetComputedHeight(),
	)
}
