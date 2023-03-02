# contains predefined functions

header = '''package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"'''

methods = {
    'clipper_path64_to_points': '''
    func (path *Path64) ToPoints() []Point64 {
		var res []Point64
		l := path.Length()
		var i int64

		for i = 0; i < l; i++ {
			n := path.GetPoint(i)
			res = append(res, n)
		}
		return res
	}
    ''',

    'clipper_pathd_to_points': '''
    func (path *PathD) ToPoints() []PointD {
		var res []PointD
		l := path.Length()
		var i int64

		for i = 0; i < l; i++ {
			n := path.GetPoint(i)
			res = append(res, n)
		}
		return res
	}
	''',

    'clipper_paths64_lengths': """
    func (paths *Paths64) Lengths() []int64 {
		var lengths []int64

		l := paths.Length()

		var i int64
		for i = 0; i < l; i++ {
			lengths = append(lengths, paths.GetPath(i).Length())
		}

		return lengths
	}
    """,

    'clipper_pathsd_lengths': """
    func (paths *PathsD) Lengths() []int64 {
		var lengths []int64

		l := paths.Length()

		var i int64
		for i = 0; i < l; i++ {
			lengths = append(lengths, paths.GetPath(i).Length())
		}

		return lengths
	}
	""",

    'clipper_rect64_to_struct': """
    func (rect *Rect64) ToStruct() Rect64 {
		return *rect
	}
	""",
    
	'clipper_rectd_to_struct': """
    func (rect *RectD) ToStruct() RectD {
		return *rect
	}
    """,

    'clipper_path64_scale': """
    func (path *Path64) Scale(sx float64, sy float64) (*Path64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &Path64{
			p: C.clipper_path64_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathd_scale': """
		func (path *PathD) Scale(sx float64, sy float64) (*PathD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &PathD{
			p: C.clipper_pathd_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_paths64_scale': """
    func (path *Paths64) Scale(sx float64, sy float64) (*Paths64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &Paths64{
			p: C.clipper_paths64_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathsd_scale': """
		func (path *PathsD) Scale(sx float64, sy float64) (*PathsD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &PathsD{
			p: C.clipper_pathsd_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",
}
