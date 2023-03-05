package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"

func (p *Path64) Destruct() {
	C.clipper_destruct_path64(p.p)
}

func (p *PathD) Destruct() {
	C.clipper_destruct_pathd(p.p)
}

func (p *Paths64) Destruct() {
	C.clipper_destruct_paths64(p.p)
}

func (p *PathsD) Destruct() {
	C.clipper_destruct_pathsd(p.p)
}

func (p *Rect64) Destruct() {
	C.clipper_destruct_rect64(p.p)
}

func (p *RectD) Destruct() {
	C.clipper_destruct_rectd(p.p)
}

func (p *PolyTree64) Destruct() {
	C.clipper_destruct_polytree64(p.p)
}

func (p *PolyTreeD) Destruct() {
	C.clipper_destruct_polytreed(p.p)
}

func (p *Clipper64) Destruct() {
	C.clipper_destruct_clipper64(p.p)
}

func (p *ClipperD) Destruct() {
	C.clipper_destruct_clipperd(p.p)
}

func (p *ClipperOffset) Destruct() {
	C.clipper_destruct_clipperoffset(p.p)
}