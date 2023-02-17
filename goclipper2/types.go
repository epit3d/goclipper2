package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import (
	"log"
	"unsafe"
)

type CPathsD struct {
	Path C.ClipperPathsD
}

type CPointD struct {
	Point C.ClipperPointD
}

type CPaths64 struct {
	Path *C.ClipperPath64
}

type CPath64 struct {
	P *C.ClipperPath64
}

func NewCPath64() *CPath64 {
	var mem unsafe.Pointer = C.malloc(500)

	res := C.clipper_path64(mem)

	return &CPath64{
		P: res,
	}
}

func (p *CPath64) AddPoint(pt CPoint64) {
	sizeBefore := C.clipper_path64_length(p.P)

	C.clipper_path64_add_point(p.P, pt.Point)

	sizeAfter := C.clipper_path64_length(p.P)
	log.Println("sizes: ", sizeBefore, sizeAfter)
}

func (subjects *CPaths64) BooleanOp(clips *CPaths64, clipType CClipType, fillRule CFillRule) *CPaths64 {
	// mem := make(unsafe.Pointer, 500)
	// var mem unsafe.Pointer
	// C.clipper_paths64_boolean_op(mem, 0, 0, subjects.Path, clips.Path)

	return nil
}

type CPoint64 struct {
	Point C.ClipperPoint64
}

func NewCPoint64(x, y int64) CPoint64 {
	return CPoint64{
		Point: C.ClipperPoint64{
			x: C.long(x),
			y: C.long(y),
		},
	}
}

type CRect64 struct {
	Rect C.ClipperRect64
}

type CRectD struct {
	Rect C.ClipperRectD
}

type CFillRule string

const (
	EvenOdd  CFillRule = "EvenOdd"
	NonZero  CFillRule = "NonZero"
	Positive CFillRule = "Positive"
	Negative CFillRule = "Negative"
)

type CClipType string

const (
	None         CClipType = "None"
	Intersection CClipType = "Intersection"
	Union        CClipType = "Union"
	Difference   CClipType = "Difference"
	XOR          CClipType = "XOR"
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
