package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"

type ClipperClipper64 struct {
	P *C.ClipperClipper64
}

type ClipperClipperD struct {
	P *C.ClipperClipperD
}

type ClipperClipperOffset struct {
	P *C.ClipperClipperOffset
}

type ClipperPath64 struct {
	P *C.ClipperPath64
}

type ClipperPathD struct {
	P *C.ClipperPathD
}

type ClipperPaths64 struct {
	P *C.ClipperPaths64
}

type ClipperPathsD struct {
	P *C.ClipperPathsD
}

type ClipperRect64 struct {
	P *C.ClipperRect64
}

type ClipperRectD struct {
	P *C.ClipperRectD
}

type ClipperPolyTree64 struct {
	P *C.ClipperPolyTree64
}

type ClipperPolyTreeD struct {
	P *C.ClipperPolyTreeD
}

type CFillRule int

const (
	EvenOdd  CFillRule = 0
	NonZero  CFillRule = 1
	Positive CFillRule = 2
	Negative CFillRule = 3
)

type CClipType int

const (
	None         CClipType = 0
	Intersection CClipType = 1
	Union        CClipType = 2
	Difference   CClipType = 3
	XOR          CClipType = 4
)

type CPathType string

const (
	Subject CPathType = "Subject"
	Clip    CPathType = "Clip"
)

type CJoinType string

const (
	SquareJoin CJoinType = "SquareJoin"
	RoundJoin  CJoinType = "RoundJoin"
	MiterJoin  CJoinType = "MiterJoin"
)

type CEndType string

const (
	PolygonEnd CEndType = "PolygonEnd"
	JoinedEnd  CEndType = "JoinedEnd"
	ButtEnd    CEndType = "ButtEnd"
	SquareEnd  CEndType = "SquareEnd"
	RoundEnd   CEndType = "RoundEnd"
)

type CPointInPolygonResult string

const (
	IsOn      CPointInPolygonResult = "IsOn"
	IsInside  CPointInPolygonResult = "IsInside"
	IsOutside CPointInPolygonResult = "IsOutside"
)
