package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"

type Clipper64 struct {
	p *C.ClipperClipper64
}

type ClipperD struct {
	p *C.ClipperClipperD
}

type ClipperOffset struct {
	p *C.ClipperClipperOffset
}

type Path64 struct {
	p *C.ClipperPath64
}

type PathD struct {
	p *C.ClipperPathD
}

type Paths64 struct {
	p *C.ClipperPaths64
}

type PathsD struct {
	p *C.ClipperPathsD
}

type Rect64 struct {
	p *C.ClipperRect64
}

type RectD struct {
	p *C.ClipperRectD
}

type PolyTree64 struct {
	p *C.ClipperPolyTree64
}

type PolyTreeD struct {
	p *C.ClipperPolyTreeD
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

type PointD struct {
	p C.ClipperPointD
}

func NewPointD(x, y float64) *PointD {
	return &PointD{
		p: C.ClipperPointD{
			x: C.double(x),
			y: C.double(y),
		},
	}
}

type Point64 struct {
	p C.ClipperPoint64
}

func NewPoint64(x, y int64) *Point64 {
	return &Point64{
		p: C.ClipperPoint64{
			x: C.int64_t(x),
			y: C.int64_t(y),
		},
	}
}
