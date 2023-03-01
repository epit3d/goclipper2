# contains predefined functions

header = '''package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"'''

methods = {
    'clipper_path64_to_points': '''
    func (path *path64) ToPoints() []point64 {
		var res []point64
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
    func (path *pathD) ToPoints() []pointD {
		var res []pointD
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
    func (paths *paths64) Lengths() []int64 {
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
    func (paths *pathsD) Lengths() []int64 {
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
    func (rect *rect64) ToStruct() rect64 {
		return *rect
	}
	""",
    
	'clipper_rectd_to_struct': """
    func (rect *rectD) ToStruct() rectD {
		return *rect
	}
    """,

    'clipper_path64_scale': """
    func (path *path64) Scale(sx float64, sy float64) (*path64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &path64{
			P: C.clipper_path64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathd_scale': """
		func (path *pathD) Scale(sx float64, sy float64) (*pathD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &pathD{
			P: C.clipper_pathd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_paths64_scale': """
    func (path *paths64) Scale(sx float64, sy float64) (*paths64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &paths64{
			P: C.clipper_paths64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathsd_scale': """
		func (path *pathsD) Scale(sx float64, sy float64) (*pathsD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &pathsD{
			P: C.clipper_pathsd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",
}
