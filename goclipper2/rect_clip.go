package goclipper2

// Code was generated by generator.py. DO NOT EDIT.

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"
import "unsafe"

func (path *Path64) RectClip(rect Rect64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Paths64{
		p: C.clipper_path64_rect_clip(mem, rect.p, path.p),
	}

}

func (path *PathD) RectClip(rect RectD, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathsD{
		p: C.clipper_pathd_rect_clip(mem, rect.p, path.p, C.int(precision)),
	}

}

func (paths *Paths64) RectClip(rect Rect64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Paths64{
		p: C.clipper_paths64_rect_clip(mem, rect.p, paths.p),
	}

}

func (paths *PathsD) RectClip(rect RectD, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathsD{
		p: C.clipper_pathsd_rect_clip(mem, rect.p, paths.p, C.int(precision)),
	}

}

func (path *Path64) RectClipLine(rect Rect64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Paths64{
		p: C.clipper_path64_rect_clip_line(mem, rect.p, path.p),
	}

}

func (path *PathD) RectClipLine(rect RectD, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathsD{
		p: C.clipper_pathd_rect_clip_line(mem, rect.p, path.p, C.int(precision)),
	}

}

func (paths *Paths64) RectClipLines(rect Rect64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &Paths64{
		p: C.clipper_paths64_rect_clip_lines(mem, rect.p, paths.p),
	}

}

func (paths *PathsD) RectClipLines(rect RectD, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &PathsD{
		p: C.clipper_pathsd_rect_clip_lines(mem, rect.p, paths.p, C.int(precision)),
	}

}
