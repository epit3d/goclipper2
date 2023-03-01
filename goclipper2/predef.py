# contains predefined functions

header = '''package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"'''

methods = {
    'clipper_path64_to_points': """\
    func (path *ClipperPath64) Clipper_path64_to_points() []ClipperPoint64 {\n\
	\tvar res []ClipperPoint64\n\
	\tl := path.Clipper_path64_length()\n\
	\tvar i int64\n\
    \t\n\
	\tfor i = 0; i < l; i++ {\n\
	\t\tn := path.Clipper_path64_get_point(i)\n\
	\t\tres = append(res, n)\n\
	\t}\n\
	\treturn res\n\
    }\n\
    """,

    'clipper_pathd_to_points': """\
    func (path *ClipperPathD) Clipper_pathd_to_points() []ClipperPointD {\n\
	\tvar res []ClipperPointD\n\
	\tl := path.Clipper_pathd_length()\n\
	\tvar i int64\n\
    \t\n\
	\tfor i = 0; i < l; i++ {\n\
	\t\tn := path.Clipper_pathd_get_point(i)\n\
	\t\tres = append(res, n)\n\
	\t}\n\
	\treturn res\n\
    }\n\
    """,

    'clipper_paths64_lengths': """
    func (paths *ClipperPaths64) Clipper_paths64_lengths() []int64 {
		var lengths []int64

		l := paths.Clipper_paths64_length()

		var i int64
		for i = 0; i < l; i++ {
			lengths = append(lengths, paths.Clipper_paths64_get_path(i).Clipper_path64_length())
		}

		return lengths
	}""",

    'clipper_pathsd_lengths': """
    func (paths *ClipperPathsD) Clipper_pathsd_lengths() []int64 {
		var lengths []int64

		l := paths.Clipper_pathsd_length()

		var i int64
		for i = 0; i < l; i++ {
			lengths = append(lengths, paths.Clipper_pathsd_get_path(i).Clipper_pathd_length())
		}

		return lengths
	}""",

    'clipper_rect64_to_struct': """
    func (rect *ClipperRect64) Clipper_rect64_to_struct() ClipperRect64 {
		return *rect
	}""",

    'clipper_rectd_to_struct': """
    func (rect *ClipperRectD) Clipper_rectd_to_struct() ClipperRectD {
		return *rect
	}
    """,

    'clipper_path64_scale': """
    func (path *ClipperPath64) Clipper_path64_scale(sx float64, sy float64) (*ClipperPath64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &ClipperPath64{
			P: C.clipper_path64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathd_scale': """
		func (path *ClipperPathD) Clipper_pathd_scale(sx float64, sy float64) (*ClipperPathD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &ClipperPathD{
			P: C.clipper_pathd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_paths64_scale': """
    func (path *ClipperPaths64) Clipper_paths64_scale(sx float64, sy float64) (*ClipperPaths64, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &ClipperPaths64{
			P: C.clipper_paths64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",

    'clipper_pathsd_scale': """
		func (path *ClipperPathsD) Clipper_pathsd_scale(sx float64, sy float64) (*ClipperPathsD, int) {
		mem := C.malloc(0)
		error_code := C.int(0)

		return &ClipperPathsD{
			P: C.clipper_pathsd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}""",
}
