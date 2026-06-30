package goyoga

import (
	yoga "github.com/husnulhamidiah/goyoga/gen"
	"math"
)

func FloatIsUndefined(value float32) bool {
	return yoga.FloatIsUndefined(value)
}

func RoundValueToPixelGrid(value float64, pointScaleFactor float64, forceCeil bool, forceFloor bool) float32 {
	return yoga.RoundValueToPixelGrid(value, pointScaleFactor, forceCeil, forceFloor)
}

func If[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

func IsInf(f float32, sign int) bool {
	return math.IsInf(float64(f), sign)
}

func IsNaN(f float32) bool {
	return math.IsNaN(float64(f))
}
