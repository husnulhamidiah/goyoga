package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../yoga
#cgo CXXFLAGS: -I${SRCDIR}/../yoga -std=c++20
#cgo darwin LDFLAGS: -lc++
#cgo linux LDFLAGS: -lstdc++
#include <yoga/Yoga.h>
*/
import "C"

var ValueAuto Value = C.YGValueAuto
var ValueUndefined Value = C.YGValueUndefined
var ValueZero Value = C.YGValueZero
