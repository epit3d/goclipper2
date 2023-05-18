package goclipper2

// #cgo CFLAGS: -I${SRCDIR}/lib
// #cgo LDFLAGS: -L${SRCDIR}/../lib -Wl,-rpath=\$ORIGIN/lib -lclipper2c
// #include "../lib/clipper2c/clipper2c.h"
import "C"
import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type Clipper64 struct {
	p *C.ClipperClipper64
}

type ClipperD struct {
	p *C.ClipperClipperD
}

type ClipperOffset struct {
	p *C.ClipperClipperOffset
}

func (c *ClipperOffset) SetZCallback(f func(e1bot, e1top, e2bot, e2top, pt *Point64)) {
	// void (*ClipperZCallback64)(const ClipperPoint64 *e1bot,
	// const ClipperPoint64 *e1top,
	// const ClipperPoint64 *e2bot,
	// const ClipperPoint64 *e2top, ClipperPoint64 *pt);

	cbk := C.ClipperZCallback64(unsafe.Pointer(&f))

	// C.clipper_offset_set_z_callback(c.p, C.ClipperZCallback64(C.ClipperZCallback64Func(func(e1bot, e1top, e2bot, e2top, pt *C.ClipperPoint64) {
	// 	e1bot_go := &Point64{p: e1bot}
	// 	e1top_go := &Point64{p: e1top}
	// 	e2bot_go := &Point64{p: e2bot}
	// 	e2top_go := &Point64{p: e2top}
	// 	pt_go := &Point64{p: pt}

	// 	c.p.z_callback(e1bot_go, e1top_go, e2bot_go, e2top_go, pt_go)
	// })))

	C.clipper_clipperoffset_set_zcallback(c.p, cbk)
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

func NewPath64MustOfString(str string) *Path64 {
	p, _ := NewPath64OfString(str)
	return p
}

func NewPath64OfString(str string) (*Path64, error) {
	re := regexp.MustCompile("[0-9-]+")
	int_strings := re.FindAllString(str, -1)
	if len(int_strings)%2 == 1 {
		log.Println(str)
		return nil, fmt.Errorf("cannot parse path with non even # of coordinates")
	}

	result := NewPath64()
	for i := 0; i < len(int_strings); i = i + 2 {
		x, err := strconv.Atoi(int_strings[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse X from '%s'", int_strings[0])
		}

		y, err := strconv.Atoi(int_strings[i+1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse Y from '%s'", int_strings[1])
		}

		result.AddPoint(*NewPoint64(int64(x), int64(y), 0))
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

func (p *PointD) String() string {
	return fmt.Sprintf("%f,%f,%f", p.p.x, p.p.y, p.p.z)
}

func (p *PointD) X() float64 {
	return float64(p.p.x)
}

func (p *PointD) Y() float64 {
	return float64(p.p.y)
}

func (p *PointD) Z() float64 {
	return float64(p.p.z)
}

func NewPointD(x, y, z float64) *PointD {
	return &PointD{
		p: C.ClipperPointD{
			x: C.double(x),
			y: C.double(y),
			z: C.double(z),
		},
	}
}

type Point64 struct {
	p C.ClipperPoint64
}

func (p *Point64) String() string {
	return fmt.Sprintf("%d,%d,%d", p.p.x, p.p.y, p.p.z)
}

func (p *Point64) X() int64 {
	return int64(p.p.x)
}

func (p *Point64) Y() int64 {
	return int64(p.p.y)
}

func (p *Point64) Z() int64 {
	return int64(p.p.z)
}

func NewPoint64(x, y, z int64) *Point64 {
	return &Point64{
		p: C.ClipperPoint64{
			x: C.int64_t(x),
			y: C.int64_t(y),
			z: C.int64_t(z),
		},
	}
}
