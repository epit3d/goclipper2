# contains predefined functions

header = '''package goclipper2

// Code was generated by generator.py. DO NOT EDIT.

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"
import "unsafe"'''

header_no_unsafe = '''package goclipper2

// Code was generated by generator.py. DO NOT EDIT.

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "/usr/local/include/clipper2c/clipper2c.h"
import "C"
'''

constructors = {
    'clipper_path64_of_points': '''
    func NewPath64OfPoints(pts []Point64) *Path64 {
		p := NewPath64()

		for _, pt := range pts {
			p.AddPoint(pt)
		}
		return p
	}
    ''',

    'clipper_pathd_of_points': '''
    func NewPathDOfPoints(pts []PointD) *PathD {
		p := NewPathd()

		for _, pt := range pts {
			p.AddPoint(pt)
		}

		return p
	}
    ''',

    'clipper_paths64_of_paths': '''
    func NewPaths64OfPaths(pths []Path64) *Paths64 {
		p := NewPaths64()

		for _, pth := range pths {
			p.AddPath(pth)
		}

		return p
	}
    ''',

    'clipper_pathsd_of_paths': '''
    func NewPathsdOfPaths(pths []PathD) *PathsD {
		p := NewPathsd()

		for _, pth := range pths {
			p.AddPath(pth)
		}

		return p
	}
	''',


}

methods = {
    'clipper_paths64_to_points': '''
    func (p *Paths64) ToPoints() [][]Point64 {
		var result [][]Point64

		for i := 0; i < int(p.Length()); i++ {
			result = append(result, p.GetPath(int64(i)).ToPoints())
		}

		return result
	}
    ''',

    'clipper_pathsd_to_points': '''
	func (p *PathsD) ToPoints() [][]PointD {
		var result [][]PointD

		for i := 0; i < int(p.Length()); i++ {
			result = append(result, p.GetPath(int64(i)).ToPoints())
		}

		return result
	}
    ''',

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

    'clipper_paths64_boolean_op': '''
    func (subjects *Paths64) BooleanOp(clipType ClipperClipType, fillRule ClipperFillRule, clips *Paths64) *Paths64 {
		var mem unsafe.Pointer = C.malloc(0)

		return &Paths64{
			p: C.clipper_paths64_boolean_op(
				mem,
				C.ClipperClipType(clipType),
				C.ClipperFillRule(fillRule),
				subjects.p,
				clips.p,
			),
		}
	}
    ''',

    'clipper_paths64_boolean_op_tree': '''
    func (subjects *Paths64) BooleanOpTree(clipType ClipperClipType, fillRule ClipperFillRule, clips *Paths64) *PolyTree64 {
		result := NewPolytree64(PolyTree64{})

		C.clipper_paths64_boolean_op_tree(
			C.ClipperClipType(clipType),
			C.ClipperFillRule(fillRule),
			subjects.p,
			clips.p,
			result.p,
		)

		return result
	}
    ''',

    'clipper_pathsd_boolean_op': '''
    func (subjects *PathsD) BooleanOp(clipType ClipperClipType, fillRule ClipperFillRule, clips *PathsD, decimalPrec int) *PathsD {
		var mem unsafe.Pointer = C.malloc(0)

		return &PathsD{
			p: C.clipper_pathsd_boolean_op(
				mem,
				C.ClipperClipType(clipType),
				C.ClipperFillRule(fillRule),
				subjects.p,
				clips.p,
				C.int(decimalPrec),
			),
		}
	}
    ''',

    'clipper_pathsd_boolean_op_tree': '''
    func (subjects *PathsD) BooleanOpTree(clipType ClipperClipType, fillRule ClipperFillRule, clips *PathsD, decimalPrec int) *PolyTreeD {
		result := NewPolytreed(PolyTreeD{})

		C.clipper_pathsd_boolean_op_tree(
			C.ClipperClipType(clipType),
			C.ClipperFillRule(fillRule),
			subjects.p,
			clips.p,
			result.p,
			C.int(decimalPrec),
		)

		return result
	}
    ''',

    'clipper_scale_pathd_to_path64': '''
    func (p *PathD) ScaleToPath64(sx, sy float64) (*Path64, int) {
		var (
			error_code                = C.int(0)
			mem        unsafe.Pointer = C.malloc(0)
		)

		return &Path64{
			p: C.clipper_scale_pathd_to_path64(mem, p.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}
    ''',

    'clipper_scale_path64_to_pathd': '''
    func (p *Path64) ScaleToPathD(sx, sy float64) (*PathD, int) {
		var (
			error_code                = C.int(0)
			mem        unsafe.Pointer = C.malloc(0)
		)

		return &PathD{
			p: C.clipper_scale_path64_to_pathd(mem, p.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}
	''',

    'clipper_scale_paths64_to_pathsd': '''
    func (p *Paths64) ScaleToPathsD(sx, sy float64) (*PathsD, int) {
		var (
			error_code                = C.int(0)
			mem        unsafe.Pointer = C.malloc(0)
		)

		return &PathsD{
			p: C.clipper_scale_paths64_to_pathsd(mem, p.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}
	''',

    'clipper_scale_pathsd_to_paths64': '''
    func (p *PathsD) ScaleToPaths64(sx, sy float64) (*Paths64, int) {
		var (
			error_code                = C.int(0)
			mem        unsafe.Pointer = C.malloc(0)
		)

		return &Paths64{
			p: C.clipper_scale_pathsd_to_paths64(mem, p.p, C.double(sx), C.double(sy), &error_code),
		}, int(error_code)
	}
    ''',
    
	'clipper_point_in_path64': '''   
	func (p *Point64) InPath64(path *Path64) ClipperPointInPolygonResult {
		return ClipperPointInPolygonResult(
			C.clipper_point_in_path64(path.p, p.p),
		)
	}
    ''',
    
	'clipper_point_in_pathd': '''
    func (p *PointD) InPath64(path *PathD) ClipperPointInPolygonResult {
		return ClipperPointInPolygonResult(
			C.clipper_point_in_pathd(path.p, p.p),
		)
	}
    ''',
}
