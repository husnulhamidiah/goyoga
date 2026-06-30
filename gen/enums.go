package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../yoga
#cgo CXXFLAGS: -I${SRCDIR}/../yoga -std=c++20
#cgo darwin LDFLAGS: -lc++
#cgo linux LDFLAGS: -lstdc++
#include <yoga/Yoga.h>
*/
import "C"

const (
	AlignAuto         Align = C.YGAlignAuto
	AlignFlexStart    Align = C.YGAlignFlexStart
	AlignCenter       Align = C.YGAlignCenter
	AlignFlexEnd      Align = C.YGAlignFlexEnd
	AlignStretch      Align = C.YGAlignStretch
	AlignBaseline     Align = C.YGAlignBaseline
	AlignSpaceBetween Align = C.YGAlignSpaceBetween
	AlignSpaceAround  Align = C.YGAlignSpaceAround
	AlignSpaceEvenly  Align = C.YGAlignSpaceEvenly
	AlignStart        Align = C.YGAlignStart
	AlignEnd          Align = C.YGAlignEnd
)

const (
	BoxSizingBorderBox  BoxSizing = C.YGBoxSizingBorderBox
	BoxSizingContentBox BoxSizing = C.YGBoxSizingContentBox
)

const (
	DimensionWidth  Dimension = C.YGDimensionWidth
	DimensionHeight Dimension = C.YGDimensionHeight
)

const (
	DirectionInherit Direction = C.YGDirectionInherit
	DirectionLTR     Direction = C.YGDirectionLTR
	DirectionRTL     Direction = C.YGDirectionRTL
)

const (
	DisplayFlex     Display = C.YGDisplayFlex
	DisplayNone     Display = C.YGDisplayNone
	DisplayContents Display = C.YGDisplayContents
	DisplayGrid     Display = C.YGDisplayGrid
)

const (
	EdgeLeft       Edge = C.YGEdgeLeft
	EdgeTop        Edge = C.YGEdgeTop
	EdgeRight      Edge = C.YGEdgeRight
	EdgeBottom     Edge = C.YGEdgeBottom
	EdgeStart      Edge = C.YGEdgeStart
	EdgeEnd        Edge = C.YGEdgeEnd
	EdgeHorizontal Edge = C.YGEdgeHorizontal
	EdgeVertical   Edge = C.YGEdgeVertical
	EdgeAll        Edge = C.YGEdgeAll
)

const (
	ErrataNone                                         Errata = C.YGErrataNone
	ErrataStretchFlexBasis                             Errata = C.YGErrataStretchFlexBasis
	ErrataAbsolutePositionWithoutInsetsExcludesPadding Errata = C.YGErrataAbsolutePositionWithoutInsetsExcludesPadding
	ErrataAbsolutePercentAgainstInnerSize              Errata = C.YGErrataAbsolutePercentAgainstInnerSize
	ErrataMinSizeUndefinedInsteadOfAuto                Errata = C.YGErrataMinSizeUndefinedInsteadOfAuto
	ErrataAll                                          Errata = C.YGErrataAll
	ErrataClassic                                      Errata = C.YGErrataClassic
)

const (
	ExperimentalFeatureWebFlexBasis           ExperimentalFeature = C.YGExperimentalFeatureWebFlexBasis
	ExperimentalFeatureFixFlexBasisFitContent ExperimentalFeature = C.YGExperimentalFeatureFixFlexBasisFitContent
)

const (
	FlexDirectionColumn        FlexDirection = C.YGFlexDirectionColumn
	FlexDirectionColumnReverse FlexDirection = C.YGFlexDirectionColumnReverse
	FlexDirectionRow           FlexDirection = C.YGFlexDirectionRow
	FlexDirectionRowReverse    FlexDirection = C.YGFlexDirectionRowReverse
)

const (
	GridTrackTypeAuto    GridTrackType = C.YGGridTrackTypeAuto
	GridTrackTypePoints  GridTrackType = C.YGGridTrackTypePoints
	GridTrackTypePercent GridTrackType = C.YGGridTrackTypePercent
	GridTrackTypeFr      GridTrackType = C.YGGridTrackTypeFr
	GridTrackTypeMinmax  GridTrackType = C.YGGridTrackTypeMinmax
)

const (
	GutterColumn Gutter = C.YGGutterColumn
	GutterRow    Gutter = C.YGGutterRow
	GutterAll    Gutter = C.YGGutterAll
)

const (
	JustifyAuto         Justify = C.YGJustifyAuto
	JustifyFlexStart    Justify = C.YGJustifyFlexStart
	JustifyCenter       Justify = C.YGJustifyCenter
	JustifyFlexEnd      Justify = C.YGJustifyFlexEnd
	JustifySpaceBetween Justify = C.YGJustifySpaceBetween
	JustifySpaceAround  Justify = C.YGJustifySpaceAround
	JustifySpaceEvenly  Justify = C.YGJustifySpaceEvenly
	JustifyStretch      Justify = C.YGJustifyStretch
	JustifyStart        Justify = C.YGJustifyStart
	JustifyEnd          Justify = C.YGJustifyEnd
)

const (
	LogLevelError   LogLevel = C.YGLogLevelError
	LogLevelWarn    LogLevel = C.YGLogLevelWarn
	LogLevelInfo    LogLevel = C.YGLogLevelInfo
	LogLevelDebug   LogLevel = C.YGLogLevelDebug
	LogLevelVerbose LogLevel = C.YGLogLevelVerbose
	LogLevelFatal   LogLevel = C.YGLogLevelFatal
)

const (
	MeasureModeUndefined MeasureMode = C.YGMeasureModeUndefined
	MeasureModeExactly   MeasureMode = C.YGMeasureModeExactly
	MeasureModeAtMost    MeasureMode = C.YGMeasureModeAtMost
)

const (
	NodeTypeDefault NodeType = C.YGNodeTypeDefault
	NodeTypeText    NodeType = C.YGNodeTypeText
)

const (
	OverflowVisible Overflow = C.YGOverflowVisible
	OverflowHidden  Overflow = C.YGOverflowHidden
	OverflowScroll  Overflow = C.YGOverflowScroll
)

const (
	PositionTypeStatic   PositionType = C.YGPositionTypeStatic
	PositionTypeRelative PositionType = C.YGPositionTypeRelative
	PositionTypeAbsolute PositionType = C.YGPositionTypeAbsolute
)

const (
	UnitUndefined  Unit = C.YGUnitUndefined
	UnitPoint      Unit = C.YGUnitPoint
	UnitPercent    Unit = C.YGUnitPercent
	UnitAuto       Unit = C.YGUnitAuto
	UnitMaxContent Unit = C.YGUnitMaxContent
	UnitFitContent Unit = C.YGUnitFitContent
	UnitStretch    Unit = C.YGUnitStretch
)

const (
	WrapNoWrap      Wrap = C.YGWrapNoWrap
	WrapWrap        Wrap = C.YGWrapWrap
	WrapWrapReverse Wrap = C.YGWrapWrapReverse
)
