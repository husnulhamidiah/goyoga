package goyoga

import yoga "github.com/husnulhamidiah/goyoga/gen"

func (n *Node) Clone() *Node {
	if n == nil {
		return nil
	}
	return wrapNode(yoga.NodeClone(n.yogaConstRef()))
}

func (n *Node) Free() {
	if n == nil || n.ref == nil {
		return
	}
	yoga.NodeFree(n.ref)
	n.ref = nil
}

func (n *Node) Destroy() {
	n.Free()
}

func (n *Node) FreeRecursive() {
	if n == nil || n.ref == nil {
		return
	}
	yoga.NodeFreeRecursive(n.ref)
	n.ref = nil
}

func (n *Node) Finalize() {
	if n == nil {
		return
	}
	yoga.NodeFinalize(n.yogaRef())
}

func (n *Node) Reset() {
	if n == nil {
		return
	}
	yoga.NodeReset(n.yogaRef())
}

func (n *Node) CalculateLayout(availableWidth float32, availableHeight float32, ownerDirection Direction) {
	if n == nil {
		return
	}
	yoga.NodeCalculateLayout(n.yogaRef(), availableWidth, availableHeight, yoga.Direction(ownerDirection))
}

func (n *Node) HasNewLayout() bool {
	if n == nil {
		return false
	}
	return yoga.NodeGetHasNewLayout(n.yogaConstRef())
}

func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	if n == nil {
		return
	}
	yoga.NodeSetHasNewLayout(n.yogaRef(), hasNewLayout)
}

func (n *Node) IsDirty() bool {
	if n == nil {
		return false
	}
	return yoga.NodeIsDirty(n.yogaConstRef())
}

func (n *Node) MarkDirty() {
	if n == nil {
		return
	}
	yoga.NodeMarkDirty(n.yogaRef())
}

func (n *Node) SetDirtiedFunc(dirtiedFunc DirtiedFunc) {
	if n == nil {
		return
	}
	yoga.NodeSetDirtiedFunc(n.yogaRef(), dirtiedFunc)
}

func (n *Node) GetDirtiedFunc() DirtiedFunc {
	if n == nil {
		return nil
	}
	return yoga.NodeGetDirtiedFunc(n.yogaConstRef())
}

func (n *Node) InsertChild(child *Node, index uint) {
	if n == nil {
		return
	}
	yoga.NodeInsertChild(n.yogaRef(), child.yogaRef(), index)
}

func (n *Node) SwapChild(child *Node, index uint) {
	if n == nil {
		return
	}
	yoga.NodeSwapChild(n.yogaRef(), child.yogaRef(), index)
}

func (n *Node) RemoveChild(child *Node) {
	if n == nil {
		return
	}
	yoga.NodeRemoveChild(n.yogaRef(), child.yogaRef())
}

func (n *Node) RemoveAllChildren() {
	if n == nil {
		return
	}
	yoga.NodeRemoveAllChildren(n.yogaRef())
}

func (n *Node) GetChild(index uint) *Node {
	if n == nil {
		return nil
	}
	return wrapNode(yoga.NodeGetChild(n.yogaRef(), index))
}

func (n *Node) GetChildCount() uint {
	if n == nil {
		return 0
	}
	return yoga.NodeGetChildCount(n.yogaConstRef())
}

func (n *Node) GetOwner() *Node {
	if n == nil {
		return nil
	}
	return wrapNode(yoga.NodeGetOwner(n.yogaRef()))
}

func (n *Node) GetParent() *Node {
	if n == nil {
		return nil
	}
	return wrapNode(yoga.NodeGetParent(n.yogaRef()))
}

func (n *Node) SetConfig(config *Config) {
	if n == nil {
		return
	}
	yoga.NodeSetConfig(n.yogaRef(), config.yogaRef())
}

func (n *Node) GetConfig() *Config {
	if n == nil {
		return nil
	}
	return wrapConstConfig(yoga.NodeGetConfig(n.yogaRef()))
}

func (n *Node) SetMeasureFunc(measureFunc MeasureFunc) {
	if n == nil {
		return
	}
	yoga.NodeSetMeasureFunc(n.yogaRef(), measureFunc)
}

func (n *Node) HasMeasureFunc() bool {
	if n == nil {
		return false
	}
	return yoga.NodeHasMeasureFunc(n.yogaConstRef())
}

func (n *Node) SetMinContentMeasureFunc(minContentMeasureFunc MinContentMeasureFunc) {
	if n == nil {
		return
	}
	yoga.NodeSetMinContentMeasureFunc(n.yogaRef(), minContentMeasureFunc)
}

func (n *Node) HasMinContentMeasureFunc() bool {
	if n == nil {
		return false
	}
	return yoga.NodeHasMinContentMeasureFunc(n.yogaConstRef())
}

func (n *Node) SetMinContentWidth(minContentWidth float32) {
	if n == nil {
		return
	}
	yoga.NodeSetMinContentWidth(n.yogaRef(), minContentWidth)
}

func (n *Node) SetMinContentHeight(minContentHeight float32) {
	if n == nil {
		return
	}
	yoga.NodeSetMinContentHeight(n.yogaRef(), minContentHeight)
}

func (n *Node) GetMinContentWidth() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeGetMinContentWidth(n.yogaConstRef())
}

func (n *Node) GetMinContentHeight() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeGetMinContentHeight(n.yogaConstRef())
}

func (n *Node) SetBaselineFunc(baselineFunc BaselineFunc) {
	if n == nil {
		return
	}
	yoga.NodeSetBaselineFunc(n.yogaRef(), baselineFunc)
}

func (n *Node) HasBaselineFunc() bool {
	if n == nil {
		return false
	}
	return yoga.NodeHasBaselineFunc(n.yogaConstRef())
}

func (n *Node) SetIsReferenceBaseline(isReferenceBaseline bool) {
	if n == nil {
		return
	}
	yoga.NodeSetIsReferenceBaseline(n.yogaRef(), isReferenceBaseline)
}

func (n *Node) IsReferenceBaseline() bool {
	if n == nil {
		return false
	}
	return yoga.NodeIsReferenceBaseline(n.yogaConstRef())
}

func (n *Node) SetNodeType(nodeType NodeType) {
	if n == nil {
		return
	}
	yoga.NodeSetNodeType(n.yogaRef(), yoga.NodeType(nodeType))
}

func (n *Node) GetNodeType() NodeType {
	if n == nil {
		return NodeType(0)
	}
	return NodeType(yoga.NodeGetNodeType(n.yogaConstRef()))
}

func (n *Node) SetAlwaysFormsContainingBlock(alwaysFormsContainingBlock bool) {
	if n == nil {
		return
	}
	yoga.NodeSetAlwaysFormsContainingBlock(n.yogaRef(), alwaysFormsContainingBlock)
}

func (n *Node) GetAlwaysFormsContainingBlock() bool {
	if n == nil {
		return false
	}
	return yoga.NodeGetAlwaysFormsContainingBlock(n.yogaConstRef())
}

func (n *Node) GetComputedLeft() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetLeft(n.yogaConstRef())
}

func (n *Node) GetComputedTop() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetTop(n.yogaConstRef())
}

func (n *Node) GetComputedRight() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetRight(n.yogaConstRef())
}

func (n *Node) GetComputedBottom() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetBottom(n.yogaConstRef())
}

func (n *Node) GetComputedWidth() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetWidth(n.yogaConstRef())
}

func (n *Node) GetComputedHeight() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetHeight(n.yogaConstRef())
}

func (n *Node) GetLayoutDirection() Direction {
	if n == nil {
		return Direction(0)
	}
	return Direction(yoga.NodeLayoutGetDirection(n.yogaConstRef()))
}

func (n *Node) HadOverflow() bool {
	if n == nil {
		return false
	}
	return yoga.NodeLayoutGetHadOverflow(n.yogaConstRef())
}

func (n *Node) GetComputedMargin(edge Edge) float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetMargin(n.yogaConstRef(), yoga.Edge(edge))
}

func (n *Node) GetComputedBorder(edge Edge) float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetBorder(n.yogaConstRef(), yoga.Edge(edge))
}

func (n *Node) GetComputedPadding(edge Edge) float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetPadding(n.yogaConstRef(), yoga.Edge(edge))
}

func (n *Node) GetRawHeight() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetRawHeight(n.yogaConstRef())
}

func (n *Node) GetRawWidth() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeLayoutGetRawWidth(n.yogaConstRef())
}

func (n *Node) CopyStyle(srcNode *Node) {
	if n == nil {
		return
	}
	yoga.NodeCopyStyle(n.yogaRef(), srcNode.yogaConstRef())
}

func (n *Node) SetDirection(direction Direction) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetDirection(n.yogaRef(), yoga.Direction(direction))
}

func (n *Node) GetDirection() Direction {
	if n == nil {
		return Direction(0)
	}
	return Direction(yoga.NodeStyleGetDirection(n.yogaConstRef()))
}

func (n *Node) SetFlexDirection(flexDirection FlexDirection) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexDirection(n.yogaRef(), yoga.FlexDirection(flexDirection))
}

func (n *Node) GetFlexDirection() FlexDirection {
	if n == nil {
		return FlexDirection(0)
	}
	return FlexDirection(yoga.NodeStyleGetFlexDirection(n.yogaConstRef()))
}

func (n *Node) SetJustifyContent(justifyContent Justify) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetJustifyContent(n.yogaRef(), yoga.Justify(justifyContent))
}

func (n *Node) GetJustifyContent() Justify {
	if n == nil {
		return Justify(0)
	}
	return Justify(yoga.NodeStyleGetJustifyContent(n.yogaConstRef()))
}

func (n *Node) SetJustifyItems(justifyItems Justify) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetJustifyItems(n.yogaRef(), yoga.Justify(justifyItems))
}

func (n *Node) GetJustifyItems() Justify {
	if n == nil {
		return Justify(0)
	}
	return Justify(yoga.NodeStyleGetJustifyItems(n.yogaConstRef()))
}

func (n *Node) SetJustifySelf(justifySelf Justify) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetJustifySelf(n.yogaRef(), yoga.Justify(justifySelf))
}

func (n *Node) GetJustifySelf() Justify {
	if n == nil {
		return Justify(0)
	}
	return Justify(yoga.NodeStyleGetJustifySelf(n.yogaConstRef()))
}

func (n *Node) SetAlignContent(alignContent Align) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetAlignContent(n.yogaRef(), yoga.Align(alignContent))
}

func (n *Node) GetAlignContent() Align {
	if n == nil {
		return Align(0)
	}
	return Align(yoga.NodeStyleGetAlignContent(n.yogaConstRef()))
}

func (n *Node) SetAlignItems(alignItems Align) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetAlignItems(n.yogaRef(), yoga.Align(alignItems))
}

func (n *Node) GetAlignItems() Align {
	if n == nil {
		return Align(0)
	}
	return Align(yoga.NodeStyleGetAlignItems(n.yogaConstRef()))
}

func (n *Node) SetAlignSelf(alignSelf Align) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetAlignSelf(n.yogaRef(), yoga.Align(alignSelf))
}

func (n *Node) GetAlignSelf() Align {
	if n == nil {
		return Align(0)
	}
	return Align(yoga.NodeStyleGetAlignSelf(n.yogaConstRef()))
}

func (n *Node) SetPositionType(positionType PositionType) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPositionType(n.yogaRef(), yoga.PositionType(positionType))
}

func (n *Node) GetPositionType() PositionType {
	if n == nil {
		return PositionType(0)
	}
	return PositionType(yoga.NodeStyleGetPositionType(n.yogaConstRef()))
}

func (n *Node) SetFlexWrap(flexWrap Wrap) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexWrap(n.yogaRef(), yoga.Wrap(flexWrap))
}

func (n *Node) GetFlexWrap() Wrap {
	if n == nil {
		return Wrap(0)
	}
	return Wrap(yoga.NodeStyleGetFlexWrap(n.yogaConstRef()))
}

func (n *Node) SetOverflow(overflow Overflow) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetOverflow(n.yogaRef(), yoga.Overflow(overflow))
}

func (n *Node) GetOverflow() Overflow {
	if n == nil {
		return Overflow(0)
	}
	return Overflow(yoga.NodeStyleGetOverflow(n.yogaConstRef()))
}

func (n *Node) SetDisplay(display Display) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetDisplay(n.yogaRef(), yoga.Display(display))
}

func (n *Node) GetDisplay() Display {
	if n == nil {
		return Display(0)
	}
	return Display(yoga.NodeStyleGetDisplay(n.yogaConstRef()))
}

func (n *Node) SetFlex(flex float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlex(n.yogaRef(), flex)
}

func (n *Node) GetFlex() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetFlex(n.yogaConstRef())
}

func (n *Node) SetFlexGrow(flexGrow float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexGrow(n.yogaRef(), flexGrow)
}

func (n *Node) GetFlexGrow() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetFlexGrow(n.yogaConstRef())
}

func (n *Node) SetFlexShrink(flexShrink float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexShrink(n.yogaRef(), flexShrink)
}

func (n *Node) GetFlexShrink() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetFlexShrink(n.yogaConstRef())
}

func (n *Node) SetFlexBasis(flexBasis float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasis(n.yogaRef(), flexBasis)
}

func (n *Node) SetFlexBasisPercent(flexBasis float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasisPercent(n.yogaRef(), flexBasis)
}

func (n *Node) SetFlexBasisAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasisAuto(n.yogaRef())
}

func (n *Node) SetFlexBasisMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasisMaxContent(n.yogaRef())
}

func (n *Node) SetFlexBasisFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasisFitContent(n.yogaRef())
}

func (n *Node) SetFlexBasisStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetFlexBasisStretch(n.yogaRef())
}

func (n *Node) GetFlexBasis() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetFlexBasis(n.yogaConstRef()))
}

func (n *Node) SetPosition(edge Edge, position float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPosition(n.yogaRef(), yoga.Edge(edge), position)
}

func (n *Node) SetPositionPercent(edge Edge, position float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPositionPercent(n.yogaRef(), yoga.Edge(edge), position)
}

func (n *Node) GetPosition(edge Edge) Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetPosition(n.yogaConstRef(), yoga.Edge(edge)))
}

func (n *Node) SetPositionAuto(edge Edge) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPositionAuto(n.yogaRef(), yoga.Edge(edge))
}

func (n *Node) SetMargin(edge Edge, margin float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMargin(n.yogaRef(), yoga.Edge(edge), margin)
}

func (n *Node) SetMarginPercent(edge Edge, margin float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMarginPercent(n.yogaRef(), yoga.Edge(edge), margin)
}

func (n *Node) SetMarginAuto(edge Edge) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMarginAuto(n.yogaRef(), yoga.Edge(edge))
}

func (n *Node) GetMargin(edge Edge) Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetMargin(n.yogaConstRef(), yoga.Edge(edge)))
}

func (n *Node) SetPadding(edge Edge, padding float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPadding(n.yogaRef(), yoga.Edge(edge), padding)
}

func (n *Node) SetPaddingPercent(edge Edge, padding float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetPaddingPercent(n.yogaRef(), yoga.Edge(edge), padding)
}

func (n *Node) GetPadding(edge Edge) Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetPadding(n.yogaConstRef(), yoga.Edge(edge)))
}

func (n *Node) SetBorder(edge Edge, border float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetBorder(n.yogaRef(), yoga.Edge(edge), border)
}

func (n *Node) GetBorder(edge Edge) float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetBorder(n.yogaConstRef(), yoga.Edge(edge))
}

func (n *Node) SetGap(gutter Gutter, gapLength float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGap(n.yogaRef(), yoga.Gutter(gutter), gapLength)
}

func (n *Node) SetGapPercent(gutter Gutter, gapLength float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGapPercent(n.yogaRef(), yoga.Gutter(gutter), gapLength)
}

func (n *Node) GetGap(gutter Gutter) Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetGap(n.yogaConstRef(), yoga.Gutter(gutter)))
}

func (n *Node) SetBoxSizing(boxSizing BoxSizing) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetBoxSizing(n.yogaRef(), yoga.BoxSizing(boxSizing))
}

func (n *Node) GetBoxSizing() BoxSizing {
	if n == nil {
		return BoxSizing(0)
	}
	return BoxSizing(yoga.NodeStyleGetBoxSizing(n.yogaConstRef()))
}

func (n *Node) SetWidth(width float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidth(n.yogaRef(), width)
}

func (n *Node) SetWidthPercent(width float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidthPercent(n.yogaRef(), width)
}

func (n *Node) SetWidthAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidthAuto(n.yogaRef())
}

func (n *Node) SetWidthMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidthMaxContent(n.yogaRef())
}

func (n *Node) SetWidthFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidthFitContent(n.yogaRef())
}

func (n *Node) SetWidthStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetWidthStretch(n.yogaRef())
}

func (n *Node) GetWidth() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetWidth(n.yogaConstRef()))
}

func (n *Node) SetHeight(height float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeight(n.yogaRef(), height)
}

func (n *Node) SetHeightPercent(height float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeightPercent(n.yogaRef(), height)
}

func (n *Node) SetHeightAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeightAuto(n.yogaRef())
}

func (n *Node) SetHeightMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeightMaxContent(n.yogaRef())
}

func (n *Node) SetHeightFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeightFitContent(n.yogaRef())
}

func (n *Node) SetHeightStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetHeightStretch(n.yogaRef())
}

func (n *Node) GetHeight() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetHeight(n.yogaConstRef()))
}

func (n *Node) SetMinWidth(minWidth float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinWidth(n.yogaRef(), minWidth)
}

func (n *Node) SetMinWidthPercent(minWidth float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinWidthPercent(n.yogaRef(), minWidth)
}

func (n *Node) SetMinWidthMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinWidthMaxContent(n.yogaRef())
}

func (n *Node) SetMinWidthFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinWidthFitContent(n.yogaRef())
}

func (n *Node) SetMinWidthStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinWidthStretch(n.yogaRef())
}

func (n *Node) GetMinWidth() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetMinWidth(n.yogaConstRef()))
}

func (n *Node) SetMinHeight(minHeight float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinHeight(n.yogaRef(), minHeight)
}

func (n *Node) SetMinHeightPercent(minHeight float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinHeightPercent(n.yogaRef(), minHeight)
}

func (n *Node) SetMinHeightMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinHeightMaxContent(n.yogaRef())
}

func (n *Node) SetMinHeightFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinHeightFitContent(n.yogaRef())
}

func (n *Node) SetMinHeightStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMinHeightStretch(n.yogaRef())
}

func (n *Node) GetMinHeight() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetMinHeight(n.yogaConstRef()))
}

func (n *Node) SetMaxWidth(maxWidth float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxWidth(n.yogaRef(), maxWidth)
}

func (n *Node) SetMaxWidthPercent(maxWidth float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxWidthPercent(n.yogaRef(), maxWidth)
}

func (n *Node) SetMaxWidthMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxWidthMaxContent(n.yogaRef())
}

func (n *Node) SetMaxWidthFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxWidthFitContent(n.yogaRef())
}

func (n *Node) SetMaxWidthStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxWidthStretch(n.yogaRef())
}

func (n *Node) GetMaxWidth() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetMaxWidth(n.yogaConstRef()))
}

func (n *Node) SetMaxHeight(maxHeight float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxHeight(n.yogaRef(), maxHeight)
}

func (n *Node) SetMaxHeightPercent(maxHeight float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxHeightPercent(n.yogaRef(), maxHeight)
}

func (n *Node) SetMaxHeightMaxContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxHeightMaxContent(n.yogaRef())
}

func (n *Node) SetMaxHeightFitContent() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxHeightFitContent(n.yogaRef())
}

func (n *Node) SetMaxHeightStretch() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetMaxHeightStretch(n.yogaRef())
}

func (n *Node) GetMaxHeight() Value {
	if n == nil {
		return Value{}
	}
	return Value(yoga.NodeStyleGetMaxHeight(n.yogaConstRef()))
}

func (n *Node) SetAspectRatio(aspectRatio float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetAspectRatio(n.yogaRef(), aspectRatio)
}

func (n *Node) GetAspectRatio() float32 {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetAspectRatio(n.yogaConstRef())
}

func (n *Node) SetGridColumnStart(gridColumnStart int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnStart(n.yogaRef(), gridColumnStart)
}

func (n *Node) SetGridColumnStartAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnStartAuto(n.yogaRef())
}

func (n *Node) SetGridColumnStartSpan(span int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnStartSpan(n.yogaRef(), span)
}

func (n *Node) GetGridColumnStart() int {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetGridColumnStart(n.yogaConstRef())
}

func (n *Node) SetGridColumnEnd(gridColumnEnd int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnEnd(n.yogaRef(), gridColumnEnd)
}

func (n *Node) SetGridColumnEndAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnEndAuto(n.yogaRef())
}

func (n *Node) SetGridColumnEndSpan(span int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridColumnEndSpan(n.yogaRef(), span)
}

func (n *Node) GetGridColumnEnd() int {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetGridColumnEnd(n.yogaConstRef())
}

func (n *Node) SetGridRowStart(gridRowStart int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowStart(n.yogaRef(), gridRowStart)
}

func (n *Node) SetGridRowStartAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowStartAuto(n.yogaRef())
}

func (n *Node) SetGridRowStartSpan(span int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowStartSpan(n.yogaRef(), span)
}

func (n *Node) GetGridRowStart() int {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetGridRowStart(n.yogaConstRef())
}

func (n *Node) SetGridRowEnd(gridRowEnd int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowEnd(n.yogaRef(), gridRowEnd)
}

func (n *Node) SetGridRowEndAuto() {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowEndAuto(n.yogaRef())
}

func (n *Node) SetGridRowEndSpan(span int) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridRowEndSpan(n.yogaRef(), span)
}

func (n *Node) GetGridRowEnd() int {
	if n == nil {
		return 0
	}
	return yoga.NodeStyleGetGridRowEnd(n.yogaConstRef())
}

func (n *Node) SetGridTemplateColumnsCount(count uint) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateColumnsCount(n.yogaRef(), count)
}

func (n *Node) SetGridTemplateColumn(index uint, typeValue GridTrackType, value float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateColumn(n.yogaRef(), index, yoga.GridTrackType(typeValue), value)
}

func (n *Node) SetGridTemplateColumnMinMax(index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateColumnMinMax(n.yogaRef(), index, yoga.GridTrackType(minType), minValue, yoga.GridTrackType(maxType), maxValue)
}

func (n *Node) SetGridTemplateRowsCount(count uint) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateRowsCount(n.yogaRef(), count)
}

func (n *Node) SetGridTemplateRow(index uint, typeValue GridTrackType, value float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateRow(n.yogaRef(), index, yoga.GridTrackType(typeValue), value)
}

func (n *Node) SetGridTemplateRowMinMax(index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridTemplateRowMinMax(n.yogaRef(), index, yoga.GridTrackType(minType), minValue, yoga.GridTrackType(maxType), maxValue)
}

func (n *Node) SetGridAutoColumnsCount(count uint) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoColumnsCount(n.yogaRef(), count)
}

func (n *Node) SetGridAutoColumn(index uint, typeValue GridTrackType, value float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoColumn(n.yogaRef(), index, yoga.GridTrackType(typeValue), value)
}

func (n *Node) SetGridAutoColumnMinMax(index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoColumnMinMax(n.yogaRef(), index, yoga.GridTrackType(minType), minValue, yoga.GridTrackType(maxType), maxValue)
}

func (n *Node) SetGridAutoRowsCount(count uint) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoRowsCount(n.yogaRef(), count)
}

func (n *Node) SetGridAutoRow(index uint, typeValue GridTrackType, value float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoRow(n.yogaRef(), index, yoga.GridTrackType(typeValue), value)
}

func (n *Node) SetGridAutoRowMinMax(index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	if n == nil {
		return
	}
	yoga.NodeStyleSetGridAutoRowMinMax(n.yogaRef(), index, yoga.GridTrackType(minType), minValue, yoga.GridTrackType(maxType), maxValue)
}

func (n *Node) GetComputedLayout() Layout {
	return Layout{
		Left:        n.GetComputedLeft(),
		Top:         n.GetComputedTop(),
		Right:       n.GetComputedRight(),
		Bottom:      n.GetComputedBottom(),
		Width:       n.GetComputedWidth(),
		Height:      n.GetComputedHeight(),
		Direction:   n.GetLayoutDirection(),
		HadOverflow: n.HadOverflow(),
	}
}
