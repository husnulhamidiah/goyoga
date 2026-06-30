package goyoga

import (
	"fmt"
	yoga "github.com/husnulhamidiah/goyoga/gen"
	"strings"
)

type Align yoga.Align

const (
	AlignAuto         Align = Align(yoga.AlignAuto)
	AlignFlexStart    Align = Align(yoga.AlignFlexStart)
	AlignCenter       Align = Align(yoga.AlignCenter)
	AlignFlexEnd      Align = Align(yoga.AlignFlexEnd)
	AlignStretch      Align = Align(yoga.AlignStretch)
	AlignBaseline     Align = Align(yoga.AlignBaseline)
	AlignSpaceBetween Align = Align(yoga.AlignSpaceBetween)
	AlignSpaceAround  Align = Align(yoga.AlignSpaceAround)
	AlignSpaceEvenly  Align = Align(yoga.AlignSpaceEvenly)
	AlignStart        Align = Align(yoga.AlignStart)
	AlignEnd          Align = Align(yoga.AlignEnd)
)

func (e Align) String() string {
	return yoga.AlignToString(yoga.Align(e))
}

func AlignFromString(s string) (Align, error) {
	if strings.EqualFold(s, AlignAuto.String()) || strings.EqualFold(s, "AlignAuto") {
		return AlignAuto, nil
	}
	if strings.EqualFold(s, AlignFlexStart.String()) || strings.EqualFold(s, "AlignFlexStart") {
		return AlignFlexStart, nil
	}
	if strings.EqualFold(s, AlignCenter.String()) || strings.EqualFold(s, "AlignCenter") {
		return AlignCenter, nil
	}
	if strings.EqualFold(s, AlignFlexEnd.String()) || strings.EqualFold(s, "AlignFlexEnd") {
		return AlignFlexEnd, nil
	}
	if strings.EqualFold(s, AlignStretch.String()) || strings.EqualFold(s, "AlignStretch") {
		return AlignStretch, nil
	}
	if strings.EqualFold(s, AlignBaseline.String()) || strings.EqualFold(s, "AlignBaseline") {
		return AlignBaseline, nil
	}
	if strings.EqualFold(s, AlignSpaceBetween.String()) || strings.EqualFold(s, "AlignSpaceBetween") {
		return AlignSpaceBetween, nil
	}
	if strings.EqualFold(s, AlignSpaceAround.String()) || strings.EqualFold(s, "AlignSpaceAround") {
		return AlignSpaceAround, nil
	}
	if strings.EqualFold(s, AlignSpaceEvenly.String()) || strings.EqualFold(s, "AlignSpaceEvenly") {
		return AlignSpaceEvenly, nil
	}
	if strings.EqualFold(s, AlignStart.String()) || strings.EqualFold(s, "AlignStart") {
		return AlignStart, nil
	}
	if strings.EqualFold(s, AlignEnd.String()) || strings.EqualFold(s, "AlignEnd") {
		return AlignEnd, nil
	}
	return Align(0), fmt.Errorf("unknown Align: %q", s)
}

type BoxSizing yoga.BoxSizing

const (
	BoxSizingBorderBox  BoxSizing = BoxSizing(yoga.BoxSizingBorderBox)
	BoxSizingContentBox BoxSizing = BoxSizing(yoga.BoxSizingContentBox)
)

func (e BoxSizing) String() string {
	return yoga.BoxSizingToString(yoga.BoxSizing(e))
}

func BoxSizingFromString(s string) (BoxSizing, error) {
	if strings.EqualFold(s, BoxSizingBorderBox.String()) || strings.EqualFold(s, "BoxSizingBorderBox") {
		return BoxSizingBorderBox, nil
	}
	if strings.EqualFold(s, BoxSizingContentBox.String()) || strings.EqualFold(s, "BoxSizingContentBox") {
		return BoxSizingContentBox, nil
	}
	return BoxSizing(0), fmt.Errorf("unknown BoxSizing: %q", s)
}

type Dimension yoga.Dimension

const (
	DimensionWidth  Dimension = Dimension(yoga.DimensionWidth)
	DimensionHeight Dimension = Dimension(yoga.DimensionHeight)
)

func (e Dimension) String() string {
	return yoga.DimensionToString(yoga.Dimension(e))
}

func DimensionFromString(s string) (Dimension, error) {
	if strings.EqualFold(s, DimensionWidth.String()) || strings.EqualFold(s, "DimensionWidth") {
		return DimensionWidth, nil
	}
	if strings.EqualFold(s, DimensionHeight.String()) || strings.EqualFold(s, "DimensionHeight") {
		return DimensionHeight, nil
	}
	return Dimension(0), fmt.Errorf("unknown Dimension: %q", s)
}

type Direction yoga.Direction

const (
	DirectionInherit Direction = Direction(yoga.DirectionInherit)
	DirectionLTR     Direction = Direction(yoga.DirectionLTR)
	DirectionRTL     Direction = Direction(yoga.DirectionRTL)
)

func (e Direction) String() string {
	return yoga.DirectionToString(yoga.Direction(e))
}

func DirectionFromString(s string) (Direction, error) {
	if strings.EqualFold(s, DirectionInherit.String()) || strings.EqualFold(s, "DirectionInherit") {
		return DirectionInherit, nil
	}
	if strings.EqualFold(s, DirectionLTR.String()) || strings.EqualFold(s, "DirectionLTR") {
		return DirectionLTR, nil
	}
	if strings.EqualFold(s, DirectionRTL.String()) || strings.EqualFold(s, "DirectionRTL") {
		return DirectionRTL, nil
	}
	return Direction(0), fmt.Errorf("unknown Direction: %q", s)
}

type Display yoga.Display

const (
	DisplayFlex     Display = Display(yoga.DisplayFlex)
	DisplayNone     Display = Display(yoga.DisplayNone)
	DisplayContents Display = Display(yoga.DisplayContents)
	DisplayGrid     Display = Display(yoga.DisplayGrid)
)

func (e Display) String() string {
	return yoga.DisplayToString(yoga.Display(e))
}

func DisplayFromString(s string) (Display, error) {
	if strings.EqualFold(s, DisplayFlex.String()) || strings.EqualFold(s, "DisplayFlex") {
		return DisplayFlex, nil
	}
	if strings.EqualFold(s, DisplayNone.String()) || strings.EqualFold(s, "DisplayNone") {
		return DisplayNone, nil
	}
	if strings.EqualFold(s, DisplayContents.String()) || strings.EqualFold(s, "DisplayContents") {
		return DisplayContents, nil
	}
	if strings.EqualFold(s, DisplayGrid.String()) || strings.EqualFold(s, "DisplayGrid") {
		return DisplayGrid, nil
	}
	return Display(0), fmt.Errorf("unknown Display: %q", s)
}

type Edge yoga.Edge

const (
	EdgeLeft       Edge = Edge(yoga.EdgeLeft)
	EdgeTop        Edge = Edge(yoga.EdgeTop)
	EdgeRight      Edge = Edge(yoga.EdgeRight)
	EdgeBottom     Edge = Edge(yoga.EdgeBottom)
	EdgeStart      Edge = Edge(yoga.EdgeStart)
	EdgeEnd        Edge = Edge(yoga.EdgeEnd)
	EdgeHorizontal Edge = Edge(yoga.EdgeHorizontal)
	EdgeVertical   Edge = Edge(yoga.EdgeVertical)
	EdgeAll        Edge = Edge(yoga.EdgeAll)
)

func (e Edge) String() string {
	return yoga.EdgeToString(yoga.Edge(e))
}

func EdgeFromString(s string) (Edge, error) {
	if strings.EqualFold(s, EdgeLeft.String()) || strings.EqualFold(s, "EdgeLeft") {
		return EdgeLeft, nil
	}
	if strings.EqualFold(s, EdgeTop.String()) || strings.EqualFold(s, "EdgeTop") {
		return EdgeTop, nil
	}
	if strings.EqualFold(s, EdgeRight.String()) || strings.EqualFold(s, "EdgeRight") {
		return EdgeRight, nil
	}
	if strings.EqualFold(s, EdgeBottom.String()) || strings.EqualFold(s, "EdgeBottom") {
		return EdgeBottom, nil
	}
	if strings.EqualFold(s, EdgeStart.String()) || strings.EqualFold(s, "EdgeStart") {
		return EdgeStart, nil
	}
	if strings.EqualFold(s, EdgeEnd.String()) || strings.EqualFold(s, "EdgeEnd") {
		return EdgeEnd, nil
	}
	if strings.EqualFold(s, EdgeHorizontal.String()) || strings.EqualFold(s, "EdgeHorizontal") {
		return EdgeHorizontal, nil
	}
	if strings.EqualFold(s, EdgeVertical.String()) || strings.EqualFold(s, "EdgeVertical") {
		return EdgeVertical, nil
	}
	if strings.EqualFold(s, EdgeAll.String()) || strings.EqualFold(s, "EdgeAll") {
		return EdgeAll, nil
	}
	return Edge(0), fmt.Errorf("unknown Edge: %q", s)
}

type Errata yoga.Errata

const (
	ErrataNone                                         Errata = Errata(yoga.ErrataNone)
	ErrataStretchFlexBasis                             Errata = Errata(yoga.ErrataStretchFlexBasis)
	ErrataAbsolutePositionWithoutInsetsExcludesPadding Errata = Errata(yoga.ErrataAbsolutePositionWithoutInsetsExcludesPadding)
	ErrataAbsolutePercentAgainstInnerSize              Errata = Errata(yoga.ErrataAbsolutePercentAgainstInnerSize)
	ErrataMinSizeUndefinedInsteadOfAuto                Errata = Errata(yoga.ErrataMinSizeUndefinedInsteadOfAuto)
	ErrataAll                                          Errata = Errata(yoga.ErrataAll)
	ErrataClassic                                      Errata = Errata(yoga.ErrataClassic)
)

func (e Errata) String() string {
	return yoga.ErrataToString(yoga.Errata(e))
}

func ErrataFromString(s string) (Errata, error) {
	if strings.EqualFold(s, ErrataNone.String()) || strings.EqualFold(s, "ErrataNone") {
		return ErrataNone, nil
	}
	if strings.EqualFold(s, ErrataStretchFlexBasis.String()) || strings.EqualFold(s, "ErrataStretchFlexBasis") {
		return ErrataStretchFlexBasis, nil
	}
	if strings.EqualFold(s, ErrataAbsolutePositionWithoutInsetsExcludesPadding.String()) || strings.EqualFold(s, "ErrataAbsolutePositionWithoutInsetsExcludesPadding") {
		return ErrataAbsolutePositionWithoutInsetsExcludesPadding, nil
	}
	if strings.EqualFold(s, ErrataAbsolutePercentAgainstInnerSize.String()) || strings.EqualFold(s, "ErrataAbsolutePercentAgainstInnerSize") {
		return ErrataAbsolutePercentAgainstInnerSize, nil
	}
	if strings.EqualFold(s, ErrataMinSizeUndefinedInsteadOfAuto.String()) || strings.EqualFold(s, "ErrataMinSizeUndefinedInsteadOfAuto") {
		return ErrataMinSizeUndefinedInsteadOfAuto, nil
	}
	if strings.EqualFold(s, ErrataAll.String()) || strings.EqualFold(s, "ErrataAll") {
		return ErrataAll, nil
	}
	if strings.EqualFold(s, ErrataClassic.String()) || strings.EqualFold(s, "ErrataClassic") {
		return ErrataClassic, nil
	}
	return Errata(0), fmt.Errorf("unknown Errata: %q", s)
}

type ExperimentalFeature yoga.ExperimentalFeature

const (
	ExperimentalFeatureWebFlexBasis           ExperimentalFeature = ExperimentalFeature(yoga.ExperimentalFeatureWebFlexBasis)
	ExperimentalFeatureFixFlexBasisFitContent ExperimentalFeature = ExperimentalFeature(yoga.ExperimentalFeatureFixFlexBasisFitContent)
)

func (e ExperimentalFeature) String() string {
	return yoga.ExperimentalFeatureToString(yoga.ExperimentalFeature(e))
}

func ExperimentalFeatureFromString(s string) (ExperimentalFeature, error) {
	if strings.EqualFold(s, ExperimentalFeatureWebFlexBasis.String()) || strings.EqualFold(s, "ExperimentalFeatureWebFlexBasis") {
		return ExperimentalFeatureWebFlexBasis, nil
	}
	if strings.EqualFold(s, ExperimentalFeatureFixFlexBasisFitContent.String()) || strings.EqualFold(s, "ExperimentalFeatureFixFlexBasisFitContent") {
		return ExperimentalFeatureFixFlexBasisFitContent, nil
	}
	return ExperimentalFeature(0), fmt.Errorf("unknown ExperimentalFeature: %q", s)
}

type FlexDirection yoga.FlexDirection

const (
	FlexDirectionColumn        FlexDirection = FlexDirection(yoga.FlexDirectionColumn)
	FlexDirectionColumnReverse FlexDirection = FlexDirection(yoga.FlexDirectionColumnReverse)
	FlexDirectionRow           FlexDirection = FlexDirection(yoga.FlexDirectionRow)
	FlexDirectionRowReverse    FlexDirection = FlexDirection(yoga.FlexDirectionRowReverse)
)

func (e FlexDirection) String() string {
	return yoga.FlexDirectionToString(yoga.FlexDirection(e))
}

func FlexDirectionFromString(s string) (FlexDirection, error) {
	if strings.EqualFold(s, FlexDirectionColumn.String()) || strings.EqualFold(s, "FlexDirectionColumn") {
		return FlexDirectionColumn, nil
	}
	if strings.EqualFold(s, FlexDirectionColumnReverse.String()) || strings.EqualFold(s, "FlexDirectionColumnReverse") {
		return FlexDirectionColumnReverse, nil
	}
	if strings.EqualFold(s, FlexDirectionRow.String()) || strings.EqualFold(s, "FlexDirectionRow") {
		return FlexDirectionRow, nil
	}
	if strings.EqualFold(s, FlexDirectionRowReverse.String()) || strings.EqualFold(s, "FlexDirectionRowReverse") {
		return FlexDirectionRowReverse, nil
	}
	return FlexDirection(0), fmt.Errorf("unknown FlexDirection: %q", s)
}

type GridTrackType yoga.GridTrackType

const (
	GridTrackTypeAuto    GridTrackType = GridTrackType(yoga.GridTrackTypeAuto)
	GridTrackTypePoints  GridTrackType = GridTrackType(yoga.GridTrackTypePoints)
	GridTrackTypePercent GridTrackType = GridTrackType(yoga.GridTrackTypePercent)
	GridTrackTypeFr      GridTrackType = GridTrackType(yoga.GridTrackTypeFr)
	GridTrackTypeMinmax  GridTrackType = GridTrackType(yoga.GridTrackTypeMinmax)
)

func (e GridTrackType) String() string {
	return yoga.GridTrackTypeToString(yoga.GridTrackType(e))
}

func GridTrackTypeFromString(s string) (GridTrackType, error) {
	if strings.EqualFold(s, GridTrackTypeAuto.String()) || strings.EqualFold(s, "GridTrackTypeAuto") {
		return GridTrackTypeAuto, nil
	}
	if strings.EqualFold(s, GridTrackTypePoints.String()) || strings.EqualFold(s, "GridTrackTypePoints") {
		return GridTrackTypePoints, nil
	}
	if strings.EqualFold(s, GridTrackTypePercent.String()) || strings.EqualFold(s, "GridTrackTypePercent") {
		return GridTrackTypePercent, nil
	}
	if strings.EqualFold(s, GridTrackTypeFr.String()) || strings.EqualFold(s, "GridTrackTypeFr") {
		return GridTrackTypeFr, nil
	}
	if strings.EqualFold(s, GridTrackTypeMinmax.String()) || strings.EqualFold(s, "GridTrackTypeMinmax") {
		return GridTrackTypeMinmax, nil
	}
	return GridTrackType(0), fmt.Errorf("unknown GridTrackType: %q", s)
}

type Gutter yoga.Gutter

const (
	GutterColumn Gutter = Gutter(yoga.GutterColumn)
	GutterRow    Gutter = Gutter(yoga.GutterRow)
	GutterAll    Gutter = Gutter(yoga.GutterAll)
)

func (e Gutter) String() string {
	return yoga.GutterToString(yoga.Gutter(e))
}

func GutterFromString(s string) (Gutter, error) {
	if strings.EqualFold(s, GutterColumn.String()) || strings.EqualFold(s, "GutterColumn") {
		return GutterColumn, nil
	}
	if strings.EqualFold(s, GutterRow.String()) || strings.EqualFold(s, "GutterRow") {
		return GutterRow, nil
	}
	if strings.EqualFold(s, GutterAll.String()) || strings.EqualFold(s, "GutterAll") {
		return GutterAll, nil
	}
	return Gutter(0), fmt.Errorf("unknown Gutter: %q", s)
}

type Justify yoga.Justify

const (
	JustifyAuto         Justify = Justify(yoga.JustifyAuto)
	JustifyFlexStart    Justify = Justify(yoga.JustifyFlexStart)
	JustifyCenter       Justify = Justify(yoga.JustifyCenter)
	JustifyFlexEnd      Justify = Justify(yoga.JustifyFlexEnd)
	JustifySpaceBetween Justify = Justify(yoga.JustifySpaceBetween)
	JustifySpaceAround  Justify = Justify(yoga.JustifySpaceAround)
	JustifySpaceEvenly  Justify = Justify(yoga.JustifySpaceEvenly)
	JustifyStretch      Justify = Justify(yoga.JustifyStretch)
	JustifyStart        Justify = Justify(yoga.JustifyStart)
	JustifyEnd          Justify = Justify(yoga.JustifyEnd)
)

func (e Justify) String() string {
	return yoga.JustifyToString(yoga.Justify(e))
}

func JustifyFromString(s string) (Justify, error) {
	if strings.EqualFold(s, JustifyAuto.String()) || strings.EqualFold(s, "JustifyAuto") {
		return JustifyAuto, nil
	}
	if strings.EqualFold(s, JustifyFlexStart.String()) || strings.EqualFold(s, "JustifyFlexStart") {
		return JustifyFlexStart, nil
	}
	if strings.EqualFold(s, JustifyCenter.String()) || strings.EqualFold(s, "JustifyCenter") {
		return JustifyCenter, nil
	}
	if strings.EqualFold(s, JustifyFlexEnd.String()) || strings.EqualFold(s, "JustifyFlexEnd") {
		return JustifyFlexEnd, nil
	}
	if strings.EqualFold(s, JustifySpaceBetween.String()) || strings.EqualFold(s, "JustifySpaceBetween") {
		return JustifySpaceBetween, nil
	}
	if strings.EqualFold(s, JustifySpaceAround.String()) || strings.EqualFold(s, "JustifySpaceAround") {
		return JustifySpaceAround, nil
	}
	if strings.EqualFold(s, JustifySpaceEvenly.String()) || strings.EqualFold(s, "JustifySpaceEvenly") {
		return JustifySpaceEvenly, nil
	}
	if strings.EqualFold(s, JustifyStretch.String()) || strings.EqualFold(s, "JustifyStretch") {
		return JustifyStretch, nil
	}
	if strings.EqualFold(s, JustifyStart.String()) || strings.EqualFold(s, "JustifyStart") {
		return JustifyStart, nil
	}
	if strings.EqualFold(s, JustifyEnd.String()) || strings.EqualFold(s, "JustifyEnd") {
		return JustifyEnd, nil
	}
	return Justify(0), fmt.Errorf("unknown Justify: %q", s)
}

type LogLevel yoga.LogLevel

const (
	LogLevelError   LogLevel = LogLevel(yoga.LogLevelError)
	LogLevelWarn    LogLevel = LogLevel(yoga.LogLevelWarn)
	LogLevelInfo    LogLevel = LogLevel(yoga.LogLevelInfo)
	LogLevelDebug   LogLevel = LogLevel(yoga.LogLevelDebug)
	LogLevelVerbose LogLevel = LogLevel(yoga.LogLevelVerbose)
	LogLevelFatal   LogLevel = LogLevel(yoga.LogLevelFatal)
)

func (e LogLevel) String() string {
	return yoga.LogLevelToString(yoga.LogLevel(e))
}

func LogLevelFromString(s string) (LogLevel, error) {
	if strings.EqualFold(s, LogLevelError.String()) || strings.EqualFold(s, "LogLevelError") {
		return LogLevelError, nil
	}
	if strings.EqualFold(s, LogLevelWarn.String()) || strings.EqualFold(s, "LogLevelWarn") {
		return LogLevelWarn, nil
	}
	if strings.EqualFold(s, LogLevelInfo.String()) || strings.EqualFold(s, "LogLevelInfo") {
		return LogLevelInfo, nil
	}
	if strings.EqualFold(s, LogLevelDebug.String()) || strings.EqualFold(s, "LogLevelDebug") {
		return LogLevelDebug, nil
	}
	if strings.EqualFold(s, LogLevelVerbose.String()) || strings.EqualFold(s, "LogLevelVerbose") {
		return LogLevelVerbose, nil
	}
	if strings.EqualFold(s, LogLevelFatal.String()) || strings.EqualFold(s, "LogLevelFatal") {
		return LogLevelFatal, nil
	}
	return LogLevel(0), fmt.Errorf("unknown LogLevel: %q", s)
}

type MeasureMode yoga.MeasureMode

const (
	MeasureModeUndefined MeasureMode = MeasureMode(yoga.MeasureModeUndefined)
	MeasureModeExactly   MeasureMode = MeasureMode(yoga.MeasureModeExactly)
	MeasureModeAtMost    MeasureMode = MeasureMode(yoga.MeasureModeAtMost)
)

func (e MeasureMode) String() string {
	return yoga.MeasureModeToString(yoga.MeasureMode(e))
}

func MeasureModeFromString(s string) (MeasureMode, error) {
	if strings.EqualFold(s, MeasureModeUndefined.String()) || strings.EqualFold(s, "MeasureModeUndefined") {
		return MeasureModeUndefined, nil
	}
	if strings.EqualFold(s, MeasureModeExactly.String()) || strings.EqualFold(s, "MeasureModeExactly") {
		return MeasureModeExactly, nil
	}
	if strings.EqualFold(s, MeasureModeAtMost.String()) || strings.EqualFold(s, "MeasureModeAtMost") {
		return MeasureModeAtMost, nil
	}
	return MeasureMode(0), fmt.Errorf("unknown MeasureMode: %q", s)
}

type NodeType yoga.NodeType

const (
	NodeTypeDefault NodeType = NodeType(yoga.NodeTypeDefault)
	NodeTypeText    NodeType = NodeType(yoga.NodeTypeText)
)

func (e NodeType) String() string {
	return yoga.NodeTypeToString(yoga.NodeType(e))
}

func NodeTypeFromString(s string) (NodeType, error) {
	if strings.EqualFold(s, NodeTypeDefault.String()) || strings.EqualFold(s, "NodeTypeDefault") {
		return NodeTypeDefault, nil
	}
	if strings.EqualFold(s, NodeTypeText.String()) || strings.EqualFold(s, "NodeTypeText") {
		return NodeTypeText, nil
	}
	return NodeType(0), fmt.Errorf("unknown NodeType: %q", s)
}

type Overflow yoga.Overflow

const (
	OverflowVisible Overflow = Overflow(yoga.OverflowVisible)
	OverflowHidden  Overflow = Overflow(yoga.OverflowHidden)
	OverflowScroll  Overflow = Overflow(yoga.OverflowScroll)
)

func (e Overflow) String() string {
	return yoga.OverflowToString(yoga.Overflow(e))
}

func OverflowFromString(s string) (Overflow, error) {
	if strings.EqualFold(s, OverflowVisible.String()) || strings.EqualFold(s, "OverflowVisible") {
		return OverflowVisible, nil
	}
	if strings.EqualFold(s, OverflowHidden.String()) || strings.EqualFold(s, "OverflowHidden") {
		return OverflowHidden, nil
	}
	if strings.EqualFold(s, OverflowScroll.String()) || strings.EqualFold(s, "OverflowScroll") {
		return OverflowScroll, nil
	}
	return Overflow(0), fmt.Errorf("unknown Overflow: %q", s)
}

type PositionType yoga.PositionType

const (
	PositionTypeStatic   PositionType = PositionType(yoga.PositionTypeStatic)
	PositionTypeRelative PositionType = PositionType(yoga.PositionTypeRelative)
	PositionTypeAbsolute PositionType = PositionType(yoga.PositionTypeAbsolute)
)

func (e PositionType) String() string {
	return yoga.PositionTypeToString(yoga.PositionType(e))
}

func PositionTypeFromString(s string) (PositionType, error) {
	if strings.EqualFold(s, PositionTypeStatic.String()) || strings.EqualFold(s, "PositionTypeStatic") {
		return PositionTypeStatic, nil
	}
	if strings.EqualFold(s, PositionTypeRelative.String()) || strings.EqualFold(s, "PositionTypeRelative") {
		return PositionTypeRelative, nil
	}
	if strings.EqualFold(s, PositionTypeAbsolute.String()) || strings.EqualFold(s, "PositionTypeAbsolute") {
		return PositionTypeAbsolute, nil
	}
	return PositionType(0), fmt.Errorf("unknown PositionType: %q", s)
}

type Unit yoga.Unit

const (
	UnitUndefined  Unit = Unit(yoga.UnitUndefined)
	UnitPoint      Unit = Unit(yoga.UnitPoint)
	UnitPercent    Unit = Unit(yoga.UnitPercent)
	UnitAuto       Unit = Unit(yoga.UnitAuto)
	UnitMaxContent Unit = Unit(yoga.UnitMaxContent)
	UnitFitContent Unit = Unit(yoga.UnitFitContent)
	UnitStretch    Unit = Unit(yoga.UnitStretch)
)

func (e Unit) String() string {
	return yoga.UnitToString(yoga.Unit(e))
}

func UnitFromString(s string) (Unit, error) {
	if strings.EqualFold(s, UnitUndefined.String()) || strings.EqualFold(s, "UnitUndefined") {
		return UnitUndefined, nil
	}
	if strings.EqualFold(s, UnitPoint.String()) || strings.EqualFold(s, "UnitPoint") {
		return UnitPoint, nil
	}
	if strings.EqualFold(s, UnitPercent.String()) || strings.EqualFold(s, "UnitPercent") {
		return UnitPercent, nil
	}
	if strings.EqualFold(s, UnitAuto.String()) || strings.EqualFold(s, "UnitAuto") {
		return UnitAuto, nil
	}
	if strings.EqualFold(s, UnitMaxContent.String()) || strings.EqualFold(s, "UnitMaxContent") {
		return UnitMaxContent, nil
	}
	if strings.EqualFold(s, UnitFitContent.String()) || strings.EqualFold(s, "UnitFitContent") {
		return UnitFitContent, nil
	}
	if strings.EqualFold(s, UnitStretch.String()) || strings.EqualFold(s, "UnitStretch") {
		return UnitStretch, nil
	}
	return Unit(0), fmt.Errorf("unknown Unit: %q", s)
}

type Wrap yoga.Wrap

const (
	WrapNoWrap      Wrap = Wrap(yoga.WrapNoWrap)
	WrapWrap        Wrap = Wrap(yoga.WrapWrap)
	WrapWrapReverse Wrap = Wrap(yoga.WrapWrapReverse)
)

func (e Wrap) String() string {
	return yoga.WrapToString(yoga.Wrap(e))
}

func WrapFromString(s string) (Wrap, error) {
	if strings.EqualFold(s, WrapNoWrap.String()) || strings.EqualFold(s, "WrapNoWrap") {
		return WrapNoWrap, nil
	}
	if strings.EqualFold(s, WrapWrap.String()) || strings.EqualFold(s, "WrapWrap") {
		return WrapWrap, nil
	}
	if strings.EqualFold(s, WrapWrapReverse.String()) || strings.EqualFold(s, "WrapWrapReverse") {
		return WrapWrapReverse, nil
	}
	return Wrap(0), fmt.Errorf("unknown Wrap: %q", s)
}
