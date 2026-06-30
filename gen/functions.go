package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../yoga
#include <yoga/Yoga.h>
*/
import "C"
import "unsafe"

func AlignToString(value Align) string {
	return C.GoString(C.YGAlignToString(value))
}

func BoxSizingToString(value BoxSizing) string {
	return C.GoString(C.YGBoxSizingToString(value))
}

func ConfigFree(config ConfigRef) {
	C.YGConfigFree(config)
}

func ConfigGetContext(config ConfigConstRef) unsafe.Pointer {
	return C.YGConfigGetContext(config)
}

func ConfigGetDefault() ConfigConstRef {
	return C.YGConfigGetDefault()
}

func ConfigGetErrata(config ConfigConstRef) Errata {
	return C.YGConfigGetErrata(config)
}

func ConfigGetPointScaleFactor(config ConfigConstRef) float32 {
	return float32(C.YGConfigGetPointScaleFactor(config))
}

func ConfigGetUseWebDefaults(config ConfigConstRef) bool {
	return bool(C.YGConfigGetUseWebDefaults(config))
}

func ConfigIsExperimentalFeatureEnabled(config ConfigConstRef, feature ExperimentalFeature) bool {
	return bool(C.YGConfigIsExperimentalFeatureEnabled(config, feature))
}

func ConfigNew() ConfigRef {
	return C.YGConfigNew()
}

func ConfigSetCloneNodeFunc(config ConfigRef, callback CloneNodeFunc) {
	C.YGConfigSetCloneNodeFunc(config, callback)
}

func ConfigSetContext(config ConfigRef, context unsafe.Pointer) {
	C.YGConfigSetContext(config, context)
}

func ConfigSetErrata(config ConfigRef, errata Errata) {
	C.YGConfigSetErrata(config, errata)
}

func ConfigSetExperimentalFeatureEnabled(config ConfigRef, feature ExperimentalFeature, enabled bool) {
	C.YGConfigSetExperimentalFeatureEnabled(config, feature, C.bool(enabled))
}

func ConfigSetLogger(config ConfigRef, logger Logger) {
	C.YGConfigSetLogger(config, logger)
}

func ConfigSetPointScaleFactor(config ConfigRef, pixelsInPoint float32) {
	C.YGConfigSetPointScaleFactor(config, C.float(pixelsInPoint))
}

func ConfigSetUseWebDefaults(config ConfigRef, enabled bool) {
	C.YGConfigSetUseWebDefaults(config, C.bool(enabled))
}

func DimensionToString(value Dimension) string {
	return C.GoString(C.YGDimensionToString(value))
}

func DirectionToString(value Direction) string {
	return C.GoString(C.YGDirectionToString(value))
}

func DisplayToString(value Display) string {
	return C.GoString(C.YGDisplayToString(value))
}

func EdgeToString(value Edge) string {
	return C.GoString(C.YGEdgeToString(value))
}

func ErrataToString(value Errata) string {
	return C.GoString(C.YGErrataToString(value))
}

func ExperimentalFeatureToString(value ExperimentalFeature) string {
	return C.GoString(C.YGExperimentalFeatureToString(value))
}

func FlexDirectionToString(value FlexDirection) string {
	return C.GoString(C.YGFlexDirectionToString(value))
}

func FloatIsUndefined(value float32) bool {
	return bool(C.YGFloatIsUndefined(C.float(value)))
}

func GridTrackTypeToString(value GridTrackType) string {
	return C.GoString(C.YGGridTrackTypeToString(value))
}

func GutterToString(value Gutter) string {
	return C.GoString(C.YGGutterToString(value))
}

func JustifyToString(value Justify) string {
	return C.GoString(C.YGJustifyToString(value))
}

func LogLevelToString(value LogLevel) string {
	return C.GoString(C.YGLogLevelToString(value))
}

func MeasureModeToString(value MeasureMode) string {
	return C.GoString(C.YGMeasureModeToString(value))
}

func NodeCalculateLayout(node NodeRef, availableWidth float32, availableHeight float32, ownerDirection Direction) {
	C.YGNodeCalculateLayout(node, C.float(availableWidth), C.float(availableHeight), ownerDirection)
}

func NodeCanUseCachedMeasurement(widthMode MeasureMode, availableWidth float32, heightMode MeasureMode, availableHeight float32, lastWidthMode MeasureMode, lastAvailableWidth float32, lastHeightMode MeasureMode, lastAvailableHeight float32, lastComputedWidth float32, lastComputedHeight float32, marginRow float32, marginColumn float32, config ConfigRef) bool {
	return bool(C.YGNodeCanUseCachedMeasurement(widthMode, C.float(availableWidth), heightMode, C.float(availableHeight), lastWidthMode, C.float(lastAvailableWidth), lastHeightMode, C.float(lastAvailableHeight), C.float(lastComputedWidth), C.float(lastComputedHeight), C.float(marginRow), C.float(marginColumn), config))
}

func NodeClone(node NodeConstRef) NodeRef {
	return C.YGNodeClone(node)
}

func NodeCopyStyle(dstNode NodeRef, srcNode NodeConstRef) {
	C.YGNodeCopyStyle(dstNode, srcNode)
}

func NodeFinalize(node NodeRef) {
	C.YGNodeFinalize(node)
}

func NodeFree(node NodeRef) {
	C.YGNodeFree(node)
}

func NodeFreeRecursive(node NodeRef) {
	C.YGNodeFreeRecursive(node)
}

func NodeGetAlwaysFormsContainingBlock(node NodeConstRef) bool {
	return bool(C.YGNodeGetAlwaysFormsContainingBlock(node))
}

func NodeGetChild(node NodeRef, index uint) NodeRef {
	return C.YGNodeGetChild(node, C.size_t(index))
}

func NodeGetChildCount(node NodeConstRef) uint {
	return uint(C.YGNodeGetChildCount(node))
}

func NodeGetConfig(node NodeRef) ConfigConstRef {
	return C.YGNodeGetConfig(node)
}

func NodeGetContext(node NodeConstRef) unsafe.Pointer {
	return C.YGNodeGetContext(node)
}

func NodeGetDirtiedFunc(node NodeConstRef) DirtiedFunc {
	return C.YGNodeGetDirtiedFunc(node)
}

func NodeGetHasNewLayout(node NodeConstRef) bool {
	return bool(C.YGNodeGetHasNewLayout(node))
}

func NodeGetMinContentHeight(node NodeConstRef) float32 {
	return float32(C.YGNodeGetMinContentHeight(node))
}

func NodeGetMinContentWidth(node NodeConstRef) float32 {
	return float32(C.YGNodeGetMinContentWidth(node))
}

func NodeGetNodeType(node NodeConstRef) NodeType {
	return C.YGNodeGetNodeType(node)
}

func NodeGetOwner(node NodeRef) NodeRef {
	return C.YGNodeGetOwner(node)
}

func NodeGetParent(node NodeRef) NodeRef {
	return C.YGNodeGetParent(node)
}

func NodeHasBaselineFunc(node NodeConstRef) bool {
	return bool(C.YGNodeHasBaselineFunc(node))
}

func NodeHasMeasureFunc(node NodeConstRef) bool {
	return bool(C.YGNodeHasMeasureFunc(node))
}

func NodeHasMinContentMeasureFunc(node NodeConstRef) bool {
	return bool(C.YGNodeHasMinContentMeasureFunc(node))
}

func NodeInsertChild(node NodeRef, child NodeRef, index uint) {
	C.YGNodeInsertChild(node, child, C.size_t(index))
}

func NodeIsDirty(node NodeConstRef) bool {
	return bool(C.YGNodeIsDirty(node))
}

func NodeIsReferenceBaseline(node NodeConstRef) bool {
	return bool(C.YGNodeIsReferenceBaseline(node))
}

func NodeLayoutGetBorder(node NodeConstRef, edge Edge) float32 {
	return float32(C.YGNodeLayoutGetBorder(node, edge))
}

func NodeLayoutGetBottom(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetBottom(node))
}

func NodeLayoutGetDirection(node NodeConstRef) Direction {
	return C.YGNodeLayoutGetDirection(node)
}

func NodeLayoutGetHadOverflow(node NodeConstRef) bool {
	return bool(C.YGNodeLayoutGetHadOverflow(node))
}

func NodeLayoutGetHeight(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetHeight(node))
}

func NodeLayoutGetLeft(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetLeft(node))
}

func NodeLayoutGetMargin(node NodeConstRef, edge Edge) float32 {
	return float32(C.YGNodeLayoutGetMargin(node, edge))
}

func NodeLayoutGetPadding(node NodeConstRef, edge Edge) float32 {
	return float32(C.YGNodeLayoutGetPadding(node, edge))
}

func NodeLayoutGetRawHeight(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetRawHeight(node))
}

func NodeLayoutGetRawWidth(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetRawWidth(node))
}

func NodeLayoutGetRight(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetRight(node))
}

func NodeLayoutGetTop(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetTop(node))
}

func NodeLayoutGetWidth(node NodeConstRef) float32 {
	return float32(C.YGNodeLayoutGetWidth(node))
}

func NodeMarkDirty(node NodeRef) {
	C.YGNodeMarkDirty(node)
}

func NodeNew() NodeRef {
	return C.YGNodeNew()
}

func NodeNewWithConfig(config ConfigConstRef) NodeRef {
	return C.YGNodeNewWithConfig(config)
}

func NodeRemoveAllChildren(node NodeRef) {
	C.YGNodeRemoveAllChildren(node)
}

func NodeRemoveChild(node NodeRef, child NodeRef) {
	C.YGNodeRemoveChild(node, child)
}

func NodeReset(node NodeRef) {
	C.YGNodeReset(node)
}

func NodeSetAlwaysFormsContainingBlock(node NodeRef, alwaysFormsContainingBlock bool) {
	C.YGNodeSetAlwaysFormsContainingBlock(node, C.bool(alwaysFormsContainingBlock))
}

func NodeSetBaselineFunc(node NodeRef, baselineFunc BaselineFunc) {
	C.YGNodeSetBaselineFunc(node, baselineFunc)
}

func NodeSetChildren(owner NodeRef, children *NodeRef, count uint) {
	C.YGNodeSetChildren(owner, (*C.YGNodeRef)(children), C.size_t(count))
}

func NodeSetConfig(node NodeRef, config ConfigRef) {
	C.YGNodeSetConfig(node, config)
}

func NodeSetContext(node NodeRef, context unsafe.Pointer) {
	C.YGNodeSetContext(node, context)
}

func NodeSetDirtiedFunc(node NodeRef, dirtiedFunc DirtiedFunc) {
	C.YGNodeSetDirtiedFunc(node, dirtiedFunc)
}

func NodeSetHasNewLayout(node NodeRef, hasNewLayout bool) {
	C.YGNodeSetHasNewLayout(node, C.bool(hasNewLayout))
}

func NodeSetIsReferenceBaseline(node NodeRef, isReferenceBaseline bool) {
	C.YGNodeSetIsReferenceBaseline(node, C.bool(isReferenceBaseline))
}

func NodeSetMeasureFunc(node NodeRef, measureFunc MeasureFunc) {
	C.YGNodeSetMeasureFunc(node, measureFunc)
}

func NodeSetMinContentHeight(node NodeRef, minContentHeight float32) {
	C.YGNodeSetMinContentHeight(node, C.float(minContentHeight))
}

func NodeSetMinContentMeasureFunc(node NodeRef, minContentMeasureFunc MinContentMeasureFunc) {
	C.YGNodeSetMinContentMeasureFunc(node, minContentMeasureFunc)
}

func NodeSetMinContentWidth(node NodeRef, minContentWidth float32) {
	C.YGNodeSetMinContentWidth(node, C.float(minContentWidth))
}

func NodeSetNodeType(node NodeRef, nodeType NodeType) {
	C.YGNodeSetNodeType(node, nodeType)
}

func NodeStyleGetAlignContent(node NodeConstRef) Align {
	return C.YGNodeStyleGetAlignContent(node)
}

func NodeStyleGetAlignItems(node NodeConstRef) Align {
	return C.YGNodeStyleGetAlignItems(node)
}

func NodeStyleGetAlignSelf(node NodeConstRef) Align {
	return C.YGNodeStyleGetAlignSelf(node)
}

func NodeStyleGetAspectRatio(node NodeConstRef) float32 {
	return float32(C.YGNodeStyleGetAspectRatio(node))
}

func NodeStyleGetBorder(node NodeConstRef, edge Edge) float32 {
	return float32(C.YGNodeStyleGetBorder(node, edge))
}

func NodeStyleGetBoxSizing(node NodeConstRef) BoxSizing {
	return C.YGNodeStyleGetBoxSizing(node)
}

func NodeStyleGetDirection(node NodeConstRef) Direction {
	return C.YGNodeStyleGetDirection(node)
}

func NodeStyleGetDisplay(node NodeConstRef) Display {
	return C.YGNodeStyleGetDisplay(node)
}

func NodeStyleGetFlex(node NodeConstRef) float32 {
	return float32(C.YGNodeStyleGetFlex(node))
}

func NodeStyleGetFlexBasis(node NodeConstRef) Value {
	return C.YGNodeStyleGetFlexBasis(node)
}

func NodeStyleGetFlexDirection(node NodeConstRef) FlexDirection {
	return C.YGNodeStyleGetFlexDirection(node)
}

func NodeStyleGetFlexGrow(node NodeConstRef) float32 {
	return float32(C.YGNodeStyleGetFlexGrow(node))
}

func NodeStyleGetFlexShrink(node NodeConstRef) float32 {
	return float32(C.YGNodeStyleGetFlexShrink(node))
}

func NodeStyleGetFlexWrap(node NodeConstRef) Wrap {
	return C.YGNodeStyleGetFlexWrap(node)
}

func NodeStyleGetGap(node NodeConstRef, gutter Gutter) Value {
	return C.YGNodeStyleGetGap(node, gutter)
}

func NodeStyleGetGridColumnEnd(node NodeConstRef) int {
	return int(C.YGNodeStyleGetGridColumnEnd(node))
}

func NodeStyleGetGridColumnStart(node NodeConstRef) int {
	return int(C.YGNodeStyleGetGridColumnStart(node))
}

func NodeStyleGetGridRowEnd(node NodeConstRef) int {
	return int(C.YGNodeStyleGetGridRowEnd(node))
}

func NodeStyleGetGridRowStart(node NodeConstRef) int {
	return int(C.YGNodeStyleGetGridRowStart(node))
}

func NodeStyleGetHeight(node NodeConstRef) Value {
	return C.YGNodeStyleGetHeight(node)
}

func NodeStyleGetJustifyContent(node NodeConstRef) Justify {
	return C.YGNodeStyleGetJustifyContent(node)
}

func NodeStyleGetJustifyItems(node NodeConstRef) Justify {
	return C.YGNodeStyleGetJustifyItems(node)
}

func NodeStyleGetJustifySelf(node NodeConstRef) Justify {
	return C.YGNodeStyleGetJustifySelf(node)
}

func NodeStyleGetMargin(node NodeConstRef, edge Edge) Value {
	return C.YGNodeStyleGetMargin(node, edge)
}

func NodeStyleGetMaxHeight(node NodeConstRef) Value {
	return C.YGNodeStyleGetMaxHeight(node)
}

func NodeStyleGetMaxWidth(node NodeConstRef) Value {
	return C.YGNodeStyleGetMaxWidth(node)
}

func NodeStyleGetMinHeight(node NodeConstRef) Value {
	return C.YGNodeStyleGetMinHeight(node)
}

func NodeStyleGetMinWidth(node NodeConstRef) Value {
	return C.YGNodeStyleGetMinWidth(node)
}

func NodeStyleGetOverflow(node NodeConstRef) Overflow {
	return C.YGNodeStyleGetOverflow(node)
}

func NodeStyleGetPadding(node NodeConstRef, edge Edge) Value {
	return C.YGNodeStyleGetPadding(node, edge)
}

func NodeStyleGetPosition(node NodeConstRef, edge Edge) Value {
	return C.YGNodeStyleGetPosition(node, edge)
}

func NodeStyleGetPositionType(node NodeConstRef) PositionType {
	return C.YGNodeStyleGetPositionType(node)
}

func NodeStyleGetWidth(node NodeConstRef) Value {
	return C.YGNodeStyleGetWidth(node)
}

func NodeStyleSetAlignContent(node NodeRef, alignContent Align) {
	C.YGNodeStyleSetAlignContent(node, alignContent)
}

func NodeStyleSetAlignItems(node NodeRef, alignItems Align) {
	C.YGNodeStyleSetAlignItems(node, alignItems)
}

func NodeStyleSetAlignSelf(node NodeRef, alignSelf Align) {
	C.YGNodeStyleSetAlignSelf(node, alignSelf)
}

func NodeStyleSetAspectRatio(node NodeRef, aspectRatio float32) {
	C.YGNodeStyleSetAspectRatio(node, C.float(aspectRatio))
}

func NodeStyleSetBorder(node NodeRef, edge Edge, border float32) {
	C.YGNodeStyleSetBorder(node, edge, C.float(border))
}

func NodeStyleSetBoxSizing(node NodeRef, boxSizing BoxSizing) {
	C.YGNodeStyleSetBoxSizing(node, boxSizing)
}

func NodeStyleSetDirection(node NodeRef, direction Direction) {
	C.YGNodeStyleSetDirection(node, direction)
}

func NodeStyleSetDisplay(node NodeRef, display Display) {
	C.YGNodeStyleSetDisplay(node, display)
}

func NodeStyleSetFlex(node NodeRef, flex float32) {
	C.YGNodeStyleSetFlex(node, C.float(flex))
}

func NodeStyleSetFlexBasis(node NodeRef, flexBasis float32) {
	C.YGNodeStyleSetFlexBasis(node, C.float(flexBasis))
}

func NodeStyleSetFlexBasisAuto(node NodeRef) {
	C.YGNodeStyleSetFlexBasisAuto(node)
}

func NodeStyleSetFlexBasisFitContent(node NodeRef) {
	C.YGNodeStyleSetFlexBasisFitContent(node)
}

func NodeStyleSetFlexBasisMaxContent(node NodeRef) {
	C.YGNodeStyleSetFlexBasisMaxContent(node)
}

func NodeStyleSetFlexBasisPercent(node NodeRef, flexBasis float32) {
	C.YGNodeStyleSetFlexBasisPercent(node, C.float(flexBasis))
}

func NodeStyleSetFlexBasisStretch(node NodeRef) {
	C.YGNodeStyleSetFlexBasisStretch(node)
}

func NodeStyleSetFlexDirection(node NodeRef, flexDirection FlexDirection) {
	C.YGNodeStyleSetFlexDirection(node, flexDirection)
}

func NodeStyleSetFlexGrow(node NodeRef, flexGrow float32) {
	C.YGNodeStyleSetFlexGrow(node, C.float(flexGrow))
}

func NodeStyleSetFlexShrink(node NodeRef, flexShrink float32) {
	C.YGNodeStyleSetFlexShrink(node, C.float(flexShrink))
}

func NodeStyleSetFlexWrap(node NodeRef, flexWrap Wrap) {
	C.YGNodeStyleSetFlexWrap(node, flexWrap)
}

func NodeStyleSetGap(node NodeRef, gutter Gutter, gapLength float32) {
	C.YGNodeStyleSetGap(node, gutter, C.float(gapLength))
}

func NodeStyleSetGapPercent(node NodeRef, gutter Gutter, gapLength float32) {
	C.YGNodeStyleSetGapPercent(node, gutter, C.float(gapLength))
}

func NodeStyleSetGridAutoColumn(node NodeRef, index uint, typeValue GridTrackType, value float32) {
	C.YGNodeStyleSetGridAutoColumn(node, C.size_t(index), typeValue, C.float(value))
}

func NodeStyleSetGridAutoColumnMinMax(node NodeRef, index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	C.YGNodeStyleSetGridAutoColumnMinMax(node, C.size_t(index), minType, C.float(minValue), maxType, C.float(maxValue))
}

func NodeStyleSetGridAutoColumnsCount(node NodeRef, count uint) {
	C.YGNodeStyleSetGridAutoColumnsCount(node, C.size_t(count))
}

func NodeStyleSetGridAutoRow(node NodeRef, index uint, typeValue GridTrackType, value float32) {
	C.YGNodeStyleSetGridAutoRow(node, C.size_t(index), typeValue, C.float(value))
}

func NodeStyleSetGridAutoRowMinMax(node NodeRef, index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	C.YGNodeStyleSetGridAutoRowMinMax(node, C.size_t(index), minType, C.float(minValue), maxType, C.float(maxValue))
}

func NodeStyleSetGridAutoRowsCount(node NodeRef, count uint) {
	C.YGNodeStyleSetGridAutoRowsCount(node, C.size_t(count))
}

func NodeStyleSetGridColumnEnd(node NodeRef, gridColumnEnd int) {
	C.YGNodeStyleSetGridColumnEnd(node, C.int(gridColumnEnd))
}

func NodeStyleSetGridColumnEndAuto(node NodeRef) {
	C.YGNodeStyleSetGridColumnEndAuto(node)
}

func NodeStyleSetGridColumnEndSpan(node NodeRef, span int) {
	C.YGNodeStyleSetGridColumnEndSpan(node, C.int(span))
}

func NodeStyleSetGridColumnStart(node NodeRef, gridColumnStart int) {
	C.YGNodeStyleSetGridColumnStart(node, C.int(gridColumnStart))
}

func NodeStyleSetGridColumnStartAuto(node NodeRef) {
	C.YGNodeStyleSetGridColumnStartAuto(node)
}

func NodeStyleSetGridColumnStartSpan(node NodeRef, span int) {
	C.YGNodeStyleSetGridColumnStartSpan(node, C.int(span))
}

func NodeStyleSetGridRowEnd(node NodeRef, gridRowEnd int) {
	C.YGNodeStyleSetGridRowEnd(node, C.int(gridRowEnd))
}

func NodeStyleSetGridRowEndAuto(node NodeRef) {
	C.YGNodeStyleSetGridRowEndAuto(node)
}

func NodeStyleSetGridRowEndSpan(node NodeRef, span int) {
	C.YGNodeStyleSetGridRowEndSpan(node, C.int(span))
}

func NodeStyleSetGridRowStart(node NodeRef, gridRowStart int) {
	C.YGNodeStyleSetGridRowStart(node, C.int(gridRowStart))
}

func NodeStyleSetGridRowStartAuto(node NodeRef) {
	C.YGNodeStyleSetGridRowStartAuto(node)
}

func NodeStyleSetGridRowStartSpan(node NodeRef, span int) {
	C.YGNodeStyleSetGridRowStartSpan(node, C.int(span))
}

func NodeStyleSetGridTemplateColumn(node NodeRef, index uint, typeValue GridTrackType, value float32) {
	C.YGNodeStyleSetGridTemplateColumn(node, C.size_t(index), typeValue, C.float(value))
}

func NodeStyleSetGridTemplateColumnMinMax(node NodeRef, index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	C.YGNodeStyleSetGridTemplateColumnMinMax(node, C.size_t(index), minType, C.float(minValue), maxType, C.float(maxValue))
}

func NodeStyleSetGridTemplateColumnsCount(node NodeRef, count uint) {
	C.YGNodeStyleSetGridTemplateColumnsCount(node, C.size_t(count))
}

func NodeStyleSetGridTemplateRow(node NodeRef, index uint, typeValue GridTrackType, value float32) {
	C.YGNodeStyleSetGridTemplateRow(node, C.size_t(index), typeValue, C.float(value))
}

func NodeStyleSetGridTemplateRowMinMax(node NodeRef, index uint, minType GridTrackType, minValue float32, maxType GridTrackType, maxValue float32) {
	C.YGNodeStyleSetGridTemplateRowMinMax(node, C.size_t(index), minType, C.float(minValue), maxType, C.float(maxValue))
}

func NodeStyleSetGridTemplateRowsCount(node NodeRef, count uint) {
	C.YGNodeStyleSetGridTemplateRowsCount(node, C.size_t(count))
}

func NodeStyleSetHeight(node NodeRef, height float32) {
	C.YGNodeStyleSetHeight(node, C.float(height))
}

func NodeStyleSetHeightAuto(node NodeRef) {
	C.YGNodeStyleSetHeightAuto(node)
}

func NodeStyleSetHeightFitContent(node NodeRef) {
	C.YGNodeStyleSetHeightFitContent(node)
}

func NodeStyleSetHeightMaxContent(node NodeRef) {
	C.YGNodeStyleSetHeightMaxContent(node)
}

func NodeStyleSetHeightPercent(node NodeRef, height float32) {
	C.YGNodeStyleSetHeightPercent(node, C.float(height))
}

func NodeStyleSetHeightStretch(node NodeRef) {
	C.YGNodeStyleSetHeightStretch(node)
}

func NodeStyleSetJustifyContent(node NodeRef, justifyContent Justify) {
	C.YGNodeStyleSetJustifyContent(node, justifyContent)
}

func NodeStyleSetJustifyItems(node NodeRef, justifyItems Justify) {
	C.YGNodeStyleSetJustifyItems(node, justifyItems)
}

func NodeStyleSetJustifySelf(node NodeRef, justifySelf Justify) {
	C.YGNodeStyleSetJustifySelf(node, justifySelf)
}

func NodeStyleSetMargin(node NodeRef, edge Edge, margin float32) {
	C.YGNodeStyleSetMargin(node, edge, C.float(margin))
}

func NodeStyleSetMarginAuto(node NodeRef, edge Edge) {
	C.YGNodeStyleSetMarginAuto(node, edge)
}

func NodeStyleSetMarginPercent(node NodeRef, edge Edge, margin float32) {
	C.YGNodeStyleSetMarginPercent(node, edge, C.float(margin))
}

func NodeStyleSetMaxHeight(node NodeRef, maxHeight float32) {
	C.YGNodeStyleSetMaxHeight(node, C.float(maxHeight))
}

func NodeStyleSetMaxHeightFitContent(node NodeRef) {
	C.YGNodeStyleSetMaxHeightFitContent(node)
}

func NodeStyleSetMaxHeightMaxContent(node NodeRef) {
	C.YGNodeStyleSetMaxHeightMaxContent(node)
}

func NodeStyleSetMaxHeightPercent(node NodeRef, maxHeight float32) {
	C.YGNodeStyleSetMaxHeightPercent(node, C.float(maxHeight))
}

func NodeStyleSetMaxHeightStretch(node NodeRef) {
	C.YGNodeStyleSetMaxHeightStretch(node)
}

func NodeStyleSetMaxWidth(node NodeRef, maxWidth float32) {
	C.YGNodeStyleSetMaxWidth(node, C.float(maxWidth))
}

func NodeStyleSetMaxWidthFitContent(node NodeRef) {
	C.YGNodeStyleSetMaxWidthFitContent(node)
}

func NodeStyleSetMaxWidthMaxContent(node NodeRef) {
	C.YGNodeStyleSetMaxWidthMaxContent(node)
}

func NodeStyleSetMaxWidthPercent(node NodeRef, maxWidth float32) {
	C.YGNodeStyleSetMaxWidthPercent(node, C.float(maxWidth))
}

func NodeStyleSetMaxWidthStretch(node NodeRef) {
	C.YGNodeStyleSetMaxWidthStretch(node)
}

func NodeStyleSetMinHeight(node NodeRef, minHeight float32) {
	C.YGNodeStyleSetMinHeight(node, C.float(minHeight))
}

func NodeStyleSetMinHeightFitContent(node NodeRef) {
	C.YGNodeStyleSetMinHeightFitContent(node)
}

func NodeStyleSetMinHeightMaxContent(node NodeRef) {
	C.YGNodeStyleSetMinHeightMaxContent(node)
}

func NodeStyleSetMinHeightPercent(node NodeRef, minHeight float32) {
	C.YGNodeStyleSetMinHeightPercent(node, C.float(minHeight))
}

func NodeStyleSetMinHeightStretch(node NodeRef) {
	C.YGNodeStyleSetMinHeightStretch(node)
}

func NodeStyleSetMinWidth(node NodeRef, minWidth float32) {
	C.YGNodeStyleSetMinWidth(node, C.float(minWidth))
}

func NodeStyleSetMinWidthFitContent(node NodeRef) {
	C.YGNodeStyleSetMinWidthFitContent(node)
}

func NodeStyleSetMinWidthMaxContent(node NodeRef) {
	C.YGNodeStyleSetMinWidthMaxContent(node)
}

func NodeStyleSetMinWidthPercent(node NodeRef, minWidth float32) {
	C.YGNodeStyleSetMinWidthPercent(node, C.float(minWidth))
}

func NodeStyleSetMinWidthStretch(node NodeRef) {
	C.YGNodeStyleSetMinWidthStretch(node)
}

func NodeStyleSetOverflow(node NodeRef, overflow Overflow) {
	C.YGNodeStyleSetOverflow(node, overflow)
}

func NodeStyleSetPadding(node NodeRef, edge Edge, padding float32) {
	C.YGNodeStyleSetPadding(node, edge, C.float(padding))
}

func NodeStyleSetPaddingPercent(node NodeRef, edge Edge, padding float32) {
	C.YGNodeStyleSetPaddingPercent(node, edge, C.float(padding))
}

func NodeStyleSetPosition(node NodeRef, edge Edge, position float32) {
	C.YGNodeStyleSetPosition(node, edge, C.float(position))
}

func NodeStyleSetPositionAuto(node NodeRef, edge Edge) {
	C.YGNodeStyleSetPositionAuto(node, edge)
}

func NodeStyleSetPositionPercent(node NodeRef, edge Edge, position float32) {
	C.YGNodeStyleSetPositionPercent(node, edge, C.float(position))
}

func NodeStyleSetPositionType(node NodeRef, positionType PositionType) {
	C.YGNodeStyleSetPositionType(node, positionType)
}

func NodeStyleSetWidth(node NodeRef, width float32) {
	C.YGNodeStyleSetWidth(node, C.float(width))
}

func NodeStyleSetWidthAuto(node NodeRef) {
	C.YGNodeStyleSetWidthAuto(node)
}

func NodeStyleSetWidthFitContent(node NodeRef) {
	C.YGNodeStyleSetWidthFitContent(node)
}

func NodeStyleSetWidthMaxContent(node NodeRef) {
	C.YGNodeStyleSetWidthMaxContent(node)
}

func NodeStyleSetWidthPercent(node NodeRef, width float32) {
	C.YGNodeStyleSetWidthPercent(node, C.float(width))
}

func NodeStyleSetWidthStretch(node NodeRef) {
	C.YGNodeStyleSetWidthStretch(node)
}

func NodeSwapChild(node NodeRef, child NodeRef, index uint) {
	C.YGNodeSwapChild(node, child, C.size_t(index))
}

func NodeTypeToString(value NodeType) string {
	return C.GoString(C.YGNodeTypeToString(value))
}

func OverflowToString(value Overflow) string {
	return C.GoString(C.YGOverflowToString(value))
}

func PositionTypeToString(value PositionType) string {
	return C.GoString(C.YGPositionTypeToString(value))
}

func RoundValueToPixelGrid(value float64, pointScaleFactor float64, forceCeil bool, forceFloor bool) float32 {
	return float32(C.YGRoundValueToPixelGrid(C.double(value), C.double(pointScaleFactor), C.bool(forceCeil), C.bool(forceFloor)))
}

func UnitToString(value Unit) string {
	return C.GoString(C.YGUnitToString(value))
}

func WrapToString(value Wrap) string {
	return C.GoString(C.YGWrapToString(value))
}
