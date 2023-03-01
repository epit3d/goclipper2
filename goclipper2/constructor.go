package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"
    func Path64() *path64 {
        var mem unsafe.Pointer = C.malloc(0)

        return &path64{
            P: C.clipper_path64(mem, ),
        }
    }
    
    func Pathd() *pathD {
        var mem unsafe.Pointer = C.malloc(0)

        return &pathD{
            P: C.clipper_pathd(mem, ),
        }
    }
    
    func Paths64() *paths64 {
        var mem unsafe.Pointer = C.malloc(0)

        return &paths64{
            P: C.clipper_paths64(mem, ),
        }
    }
    
    func Pathsd() *pathsD {
        var mem unsafe.Pointer = C.malloc(0)

        return &pathsD{
            P: C.clipper_pathsd(mem, ),
        }
    }
    
    func Polytree64(parent polyTree64) *polyTree64 {
        var mem unsafe.Pointer = C.malloc(0)

        return &polyTree64{
            P: C.clipper_polytree64(mem, parent.P),
        }
    }
    
    func Polytreed(parent polyTreeD) *polyTreeD {
        var mem unsafe.Pointer = C.malloc(0)

        return &polyTreeD{
            P: C.clipper_polytreed(mem, parent.P),
        }
    }
    
    func Rect64(left int64, top int64, right int64, bottom int64) *rect64 {
        var mem unsafe.Pointer = C.malloc(0)

        return &rect64{
            P: C.clipper_rect64(mem, C.int64_t(left), C.int64_t(top), C.int64_t(right), C.int64_t(bottom)),
        }
    }
    
    func Rectd(left float64, top float64, right float64, bottom float64) *rectD {
        var mem unsafe.Pointer = C.malloc(0)

        return &rectD{
            P: C.clipper_rectd(mem, C.double(left), C.double(top), C.double(right), C.double(bottom)),
        }
    }
    
    func Clipper64() *clipper64 {
        var mem unsafe.Pointer = C.malloc(0)

        return &clipper64{
            P: C.clipper_clipper64(mem, ),
        }
    }
    
    func Clipperd(precision int64) *clipperD {
        var mem unsafe.Pointer = C.malloc(0)

        return &clipperD{
            P: C.clipper_clipperd(mem, C.int(precision)),
        }
    }
    
    func Clipperoffset(miter_limit float64, arc_tolerance float64, preserve_collinear int64, reverse_solution int64) *clipperOffset {
        var mem unsafe.Pointer = C.malloc(0)

        return &clipperOffset{
            P: C.clipper_clipperoffset(mem, C.double(miter_limit), C.double(arc_tolerance), C.int(preserve_collinear), C.int(reverse_solution)),
        }
    }
    