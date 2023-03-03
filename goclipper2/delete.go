package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"

func (p *Path64) Delete() {
	C.clipper_delete_path64(p.p)
}

func (p *PathD) Delete() {
	C.clipper_delete_pathd(p.p)
}

func (p *Paths64) Delete() {
	C.clipper_delete_paths64(p.p)
}

func (p *PathsD) Delete() {
	C.clipper_delete_pathsd(p.p)
}

func (p *Rect64) Delete() {
	C.clipper_delete_rect64(p.p)
}

func (p *RectD) Delete() {
	C.clipper_delete_rectd(p.p)
}

func (p *PolyTree64) Delete() {
	C.clipper_delete_polytree64(p.p)
}

func (p *PolyTreeD) Delete() {
	C.clipper_delete_polytreed(p.p)
}

func (p *Clipper64) Delete() {
	C.clipper_delete_clipper64(p.p)
}

func (p *ClipperD) Delete() {
	C.clipper_delete_clipperd(p.p)
}

func (p *ClipperOffset) Delete() {
	C.clipper_delete_clipperoffset(p.p)
}
