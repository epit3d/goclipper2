package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"

func Clipper_path64() *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64(mem),
	}
}

func Clipper_pathd() *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd(mem),
	}
}

func Clipper_paths64() *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64(mem),
	}
}

func Clipper_pathsd() *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd(mem),
	}
}

func Clipper_polytree64(parent ClipperPolyTree64) *ClipperPolyTree64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPolyTree64{
		P: C.clipper_polytree64(mem, parent.P),
	}
}

func Clipper_polytreed(parent ClipperPolyTreeD) *ClipperPolyTreeD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPolyTreeD{
		P: C.clipper_polytreed(mem, parent.P),
	}
}

func Clipper_rect64(left int64, top int64, right int64, bottom int64) *ClipperRect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRect64{
		P: C.clipper_rect64(mem, C.int64_t(left), C.int64_t(top), C.int64_t(right), C.int64_t(bottom)),
	}
}

func Clipper_rectd(left float64, top float64, right float64, bottom float64) *ClipperRectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRectD{
		P: C.clipper_rectd(mem, C.double(left), C.double(top), C.double(right), C.double(bottom)),
	}
}

func Clipper_clipper64() *ClipperClipper64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperClipper64{
		P: C.clipper_clipper64(mem),
	}
}

func Clipper_clipperd(precision int64) *ClipperClipperD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperClipperD{
		P: C.clipper_clipperd(mem, C.int(precision)),
	}
}

func Clipper_clipperoffset(miter_limit float64, arc_tolerance float64, preserve_collinear int64, reverse_solution int64) *ClipperClipperOffset {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperClipperOffset{
		P: C.clipper_clipperoffset(mem, C.double(miter_limit), C.double(arc_tolerance), C.int(preserve_collinear), C.int(reverse_solution)),
	}
}
