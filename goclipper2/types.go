package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"
import (
	"fmt"
	"log"
	"strings"
	"unsafe"
)

type Clipper64 struct {
	mem *unsafe.Pointer
	p   *C.ClipperClipper64
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

func (p *Path64) String() string {
	pts := p.ToPoints()

	pts_strings := []string{}
	for _, pt := range pts {
		pts_strings = append(pts_strings, pt.String())
	}

	return strings.Join(pts_strings, ", ")
}

func NewPath64OfString(str string) (*Path64, error) {
	pts_strings := strings.Split(str, ", ")

	result := NewPath64()
	for _, pt_str := range pts_strings {
		var x, y int64
		if _, err := fmt.Sscanf(pt_str, "%d,%d", &x, &y); err != nil {
			return nil, fmt.Errorf("failed to create new path64 of strings, error point: <%v>", pt_str)
		}

		result.AddPoint(*NewPoint64(x, y))
	}

	return result, nil
}

type PathD struct {
	p *C.ClipperPathD
}

type Paths64 struct {
	p *C.ClipperPaths64
}

func (p *Paths64) String() string {
	result := ""

	log.Println("path lengths", p.Length())
	for i := 0; i < int(p.Length()); i++ {
		result += p.GetPath(int64(i)).String()

		if i != int(p.Length())-1 {
			result += "\n"
		}
	}

	return result
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

func (p *Point64) String() string {
	return fmt.Sprintf("%d,%d", p.p.x, p.p.y)
}

func NewPoint64(x, y int64) *Point64 {
	return &Point64{
		p: C.ClipperPoint64{
			x: C.int64_t(x),
			y: C.int64_t(y),
		},
	}
}
