package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"

type clipper64 struct {
	P *C.ClipperClipper64
}

type clipperD struct {
	P *C.ClipperClipperD
}

type clipperOffset struct {
	P *C.ClipperClipperOffset
}

type path64 struct {
	P *C.ClipperPath64
}

type pathD struct {
	P *C.ClipperPathD
}

type paths64 struct {
	P *C.ClipperPaths64
}

type pathsD struct {
	P *C.ClipperPathsD
}

type rect64 struct {
	P *C.ClipperRect64
}

type rectD struct {
	P *C.ClipperRectD
}

type polyTree64 struct {
	P *C.ClipperPolyTree64
}

type polyTreeD struct {
	P *C.ClipperPolyTreeD
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

type pointD struct {
	P C.ClipperPointD
}

func PointD(x, y float64) *pointD {
	return &pointD{
		P: C.ClipperPointD{
			x: C.double(x),
			y: C.double(y),
		},
	}
}

type point64 struct {
	P C.ClipperPoint64
}

func Point64(x, y int64) *point64 {
	return &point64{
		P: C.ClipperPoint64{
			x: C.int64_t(x),
			y: C.int64_t(y),
		},
	}
}
