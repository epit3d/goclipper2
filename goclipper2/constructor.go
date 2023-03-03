package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"

func NewPath64() *Path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Path64{
		p: C.clipper_path64(mem),
	}
}

func NewPathd() *PathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathD{
		p: C.clipper_pathd(mem),
	}
}

func NewPath64OfString(str string) *Path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Path64{
		p: C.clipper_path64_of_string(mem, C.CString(str)),
	}
}

func NewPathdOfString(str string) *PathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathD{
		p: C.clipper_pathd_of_string(mem, C.CString(str)),
	}
}

func NewPath64OfPoints(pts []Point64) *Path64 {
	p := NewPath64()

	for _, pt := range pts {
		p.AddPoint(pt)
	}
	return p
}

func NewPathDOfPoints(pts []PointD) *PathD {
	p := NewPathd()

	for _, pt := range pts {
		p.AddPoint(pt)
	}

	return p
}

func NewPath64Ellipse(center Point64, radius_x float64, radius_y float64, steps int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Path64{
		p: C.clipper_path64_ellipse(mem, center.p, C.double(radius_x), C.double(radius_y), C.int(steps)),
	}
}

func NewPathdEllipse(center PointD, radius_x float64, radius_y float64, steps int64) *PathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathD{
		p: C.clipper_pathd_ellipse(mem, center.p, C.double(radius_x), C.double(radius_y), C.int(steps)),
	}
}

func NewPaths64() *Paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Paths64{
		p: C.clipper_paths64(mem),
	}
}

func NewPathsd() *PathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathsD{
		p: C.clipper_pathsd(mem),
	}
}

func NewPaths64OfPaths(pths []Path64) *Paths64 {
	p := NewPaths64()

	for _, pth := range pths {
		p.AddPath(pth)
	}

	return p
}

func NewPathsdOfPaths(pths []PathD) *PathsD {
	p := NewPathsd()

	for _, pth := range pths {
		p.AddPath(pth)
	}

	return p
}

func NewPolytree64(parent PolyTree64) *PolyTree64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &PolyTree64{
		p: C.clipper_polytree64(mem, parent.p),
	}
}

func NewPolytreed(parent PolyTreeD) *PolyTreeD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PolyTreeD{
		p: C.clipper_polytreed(mem, parent.p),
	}
}

func NewRect64(left int64, top int64, right int64, bottom int64) *Rect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Rect64{
		p: C.clipper_rect64(mem, C.int64_t(left), C.int64_t(top), C.int64_t(right), C.int64_t(bottom)),
	}
}

func NewRectd(left float64, top float64, right float64, bottom float64) *RectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &RectD{
		p: C.clipper_rectd(mem, C.double(left), C.double(top), C.double(right), C.double(bottom)),
	}
}

func NewClipper64() *Clipper64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Clipper64{
		p: C.clipper_clipper64(mem),
	}
}

func NewClipperd(precision int64) *ClipperD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperD{
		p: C.clipper_clipperd(mem, C.int(precision)),
	}
}

func NewClipperoffset(miter_limit float64, arc_tolerance float64, preserve_collinear int64, reverse_solution int64) *ClipperOffset {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperOffset{
		p: C.clipper_clipperoffset(mem, C.double(miter_limit), C.double(arc_tolerance), C.int(preserve_collinear), C.int(reverse_solution)),
	}
}
