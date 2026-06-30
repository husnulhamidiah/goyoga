package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../yoga
#include <yoga/Yoga.h>
*/
import "C"

type Align = C.YGAlign
type BaselineFunc = C.YGBaselineFunc
type BoxSizing = C.YGBoxSizing
type CloneNodeFunc = C.YGCloneNodeFunc
type ConfigConstRef = C.YGConfigConstRef
type ConfigRef = C.YGConfigRef
type Dimension = C.YGDimension
type Direction = C.YGDirection
type DirtiedFunc = C.YGDirtiedFunc
type Display = C.YGDisplay
type Edge = C.YGEdge
type Errata = C.YGErrata
type ExperimentalFeature = C.YGExperimentalFeature
type FlexDirection = C.YGFlexDirection
type GridTrackType = C.YGGridTrackType
type Gutter = C.YGGutter
type Justify = C.YGJustify
type LogLevel = C.YGLogLevel
type Logger = C.YGLogger
type MeasureFunc = C.YGMeasureFunc
type MeasureMode = C.YGMeasureMode
type MinContentMeasureFunc = C.YGMinContentMeasureFunc
type NodeConstRef = C.YGNodeConstRef
type NodeRef = C.YGNodeRef
type NodeType = C.YGNodeType
type Overflow = C.YGOverflow
type PositionType = C.YGPositionType
type Size = C.YGSize
type Unit = C.YGUnit
type Value = C.YGValue
type Wrap = C.YGWrap
