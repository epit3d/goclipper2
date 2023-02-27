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

type ClipperPointD struct {
	P C.ClipperPointD
}

type ClipperPoint64 struct {
	P C.ClipperPoint64
}

type ClipperFillRule int

const (
	EvenOdd  ClipperFillRule = 0
	NonZero  ClipperFillRule = 1
	Positive ClipperFillRule = 2
	Negative ClipperFillRule = 3
)

type ClipperClipType int

const (
	None         ClipperClipType = 0
	Intersection ClipperClipType = 1
	Union        ClipperClipType = 2
	Difference   ClipperClipType = 3
	XOR          ClipperClipType = 4
)

type ClipperPathType int

const (
	Subject ClipperPathType = 0
	Clip    ClipperPathType = 1
)

type ClipperJoinType int

const (
	SquareJoin ClipperJoinType = 0
	RoundJoin  ClipperJoinType = 1
	MiterJoin  ClipperJoinType = 2
)

type ClipperEndType int

const (
	PolygonEnd ClipperEndType = 0
	JoinedEnd  ClipperEndType = 1
	ButtEnd    ClipperEndType = 2
	SquareEnd  ClipperEndType = 3
	RoundEnd   ClipperEndType = 4
)

type ClipperPointInPolygonResult int

const (
	IsOn      ClipperPointInPolygonResult = 0
	IsInside  ClipperPointInPolygonResult = 1
	IsOutside ClipperPointInPolygonResult = 2
)
