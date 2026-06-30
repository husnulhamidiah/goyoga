package yoga

/*
#cgo CFLAGS: -I${SRCDIR}/../yoga
#include <yoga/Yoga.h>
*/
import "C"

var ValueAuto Value = C.YGValueAuto
var ValueUndefined Value = C.YGValueUndefined
var ValueZero Value = C.YGValueZero
