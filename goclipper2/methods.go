package goclipper2

// Code was generated by generator.py. DO NOT EDIT.

// #include "../lib/clipper2c/clipper2c.h"
import "C"
import "unsafe"

func (subjects *Paths64) BooleanOp(clipType ClipperClipType, fillRule ClipperFillRule, clips *Paths64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

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

func (subjects *Paths64) Intersect(clips Paths64, fillrule ClipperFillRule) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_intersect(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *Paths64) Union(clips Paths64, fillrule ClipperFillRule) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_union(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *Paths64) Difference(clips Paths64, fillrule ClipperFillRule) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_difference(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *Paths64) Xor(clips Paths64, fillrule ClipperFillRule) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_xor(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *PathsD) BooleanOp(clipType ClipperClipType, fillRule ClipperFillRule, clips *PathsD, decimalPrec int) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

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

func (subjects *PathsD) Intersect(clips PathsD, fillrule ClipperFillRule, decimal_prec int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_intersect(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *PathsD) Union(clips PathsD, fillrule ClipperFillRule, decimal_prec int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_union(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *PathsD) Difference(clips PathsD, fillrule ClipperFillRule, decimal_prec int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_difference(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *PathsD) Xor(clips PathsD, fillrule ClipperFillRule, decimal_prec int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_xor(mem, subjects.p, clips.p, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (paths *Paths64) Inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_inflate(mem, paths.p, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit)),
	}

}

func (paths *PathsD) Inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_inflate(mem, paths.p, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit), C.int(precision)),
	}

}

func (paths *Paths64) Bounds() *Rect64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_rect64_size())

	return &Rect64{
		p: C.clipper_paths64_bounds(mem, paths.p),
	}

}

func (path *Path64) AddPoint(pt Point64) {

	C.clipper_path64_add_point(path.p, pt.p)
}

func (path *PathD) AddPoint(pt PointD) {

	C.clipper_pathd_add_point(path.p, pt.p)
}

func (paths *Paths64) AddPath(p Path64) {

	C.clipper_paths64_add_path(paths.p, p.p)
}

func (paths *PathsD) AddPath(p PathD) {

	C.clipper_pathsd_add_path(paths.p, p.p)
}

func (path *Path64) Length() int64 {

	return int64(C.clipper_path64_length(path.p))
}

func (path *PathD) Length() int64 {

	return int64(C.clipper_pathd_length(path.p))
}

func (path *Path64) GetPoint(idx int64) Point64 {

	return Point64{
		p: C.clipper_path64_get_point(path.p, C.int(idx)),
	}

}

func (path *PathD) GetPoint(idx int64) PointD {

	return PointD{
		p: C.clipper_pathd_get_point(path.p, C.int(idx)),
	}

}

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

func (paths *Paths64) Length() int64 {

	return int64(C.clipper_paths64_length(paths.p))
}

func (paths *PathsD) Length() int64 {

	return int64(C.clipper_pathsd_length(paths.p))
}

func (paths *Paths64) Lengths() []int64 {
	var lengths []int64

	l := paths.Length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.GetPath(i).Length())
	}

	return lengths
}

func (paths *PathsD) Lengths() []int64 {
	var lengths []int64

	l := paths.Length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.GetPath(i).Length())
	}

	return lengths
}

func (paths *Paths64) PathLength(idx int64) int64 {

	return int64(C.clipper_paths64_path_length(paths.p, C.int(idx)))
}

func (paths *PathsD) PathLength(idx int64) int64 {

	return int64(C.clipper_pathsd_path_length(paths.p, C.int(idx)))
}

func (paths *Paths64) GetPath(idx int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_paths64_get_path(mem, paths.p, C.int(idx)),
	}

}

func (paths *PathsD) GetPath(idx int64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathsd_get_path(mem, paths.p, C.int(idx)),
	}

}

func (paths *Paths64) GetPoint(path_idx int64, point_idx int64) Point64 {

	return Point64{
		p: C.clipper_paths64_get_point(paths.p, C.int(path_idx), C.int(point_idx)),
	}

}

func (paths *PathsD) GetPoint(path_idx int64, point_idx int64) PointD {

	return PointD{
		p: C.clipper_pathsd_get_point(paths.p, C.int(path_idx), C.int(point_idx)),
	}

}

func (p *Paths64) ToPoints() [][]Point64 {
	var result [][]Point64

	for i := 0; i < int(p.Length()); i++ {
		result = append(result, p.GetPath(int64(i)).ToPoints())
	}

	return result
}

func (p *PathsD) ToPoints() [][]PointD {
	var result [][]PointD

	for i := 0; i < int(p.Length()); i++ {
		result = append(result, p.GetPath(int64(i)).ToPoints())
	}

	return result
}

func (path *Path64) Translate(dx int64, dy int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_path64_translate(mem, path.p, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (path *PathD) Translate(dx float64, dy float64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathd_translate(mem, path.p, C.double(dx), C.double(dy)),
	}

}

func (paths *Paths64) Translate(dx int64, dy int64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_translate(mem, paths.p, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (paths *PathsD) Translate(dx float64, dy float64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_translate(mem, paths.p, C.double(dx), C.double(dy)),
	}

}

func (path *Path64) Scale(sx float64, sy float64) (*Path64, int) {
	mem := C.malloc(C.clipper_path64_size())
	error_code := C.int(0)

	return &Path64{
		p: C.clipper_path64_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *PathD) Scale(sx float64, sy float64) (*PathD, int) {
	mem := C.malloc(C.clipper_pathd_size())
	error_code := C.int(0)

	return &PathD{
		p: C.clipper_pathd_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *Paths64) Scale(sx float64, sy float64) (*Paths64, int) {
	mem := C.malloc(C.clipper_paths64_size())
	error_code := C.int(0)

	return &Paths64{
		p: C.clipper_paths64_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *PathsD) Scale(sx float64, sy float64) (*PathsD, int) {
	mem := C.malloc(C.clipper_pathsd_size())
	error_code := C.int(0)

	return &PathsD{
		p: C.clipper_pathsd_scale(mem, path.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *Path64) TrimCollinear(is_open_path int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_path64_trim_collinear(mem, path.p, C.int(is_open_path)),
	}

}

func (path *PathD) TrimCollinear(precision int64, is_open_path int64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathd_trim_collinear(mem, path.p, C.int(precision), C.int(is_open_path)),
	}

}

func (path *Path64) Simplify(epsilon float64, is_open_path int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_path64_simplify(mem, path.p, C.double(epsilon), C.int(is_open_path)),
	}

}

func (path *PathD) Simplify(epsilon float64, is_open_path int64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathd_simplify(mem, path.p, C.double(epsilon), C.int(is_open_path)),
	}

}

func (paths *Paths64) Simplify(epsilon float64, is_open_paths int64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_simplify(mem, paths.p, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (paths *PathsD) Simplify(epsilon float64, is_open_paths int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_simplify(mem, paths.p, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (path *Path64) RamerDouglasPeucker(epsilon float64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_path64_ramer_douglas_peucker(mem, path.p, C.double(epsilon)),
	}

}

func (path *PathD) RamerDouglasPeucker(epsilon float64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathd_ramer_douglas_peucker(mem, path.p, C.double(epsilon)),
	}

}

func (paths *Paths64) RamerDouglasPeucker(epsilon float64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_ramer_douglas_peucker(mem, paths.p, C.double(epsilon)),
	}

}

func (paths *PathsD) RamerDouglasPeucker(epsilon float64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_ramer_douglas_peucker(mem, paths.p, C.double(epsilon)),
	}

}

func (path *Path64) StripNearEqual(max_dist_sqrd float64, is_closed_path int64) *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_path64_strip_near_equal(mem, path.p, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (path *PathD) StripNearEqual(max_dist_sqrd float64, is_closed_path int64) *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_pathd_strip_near_equal(mem, path.p, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (paths *Paths64) StripNearEqual(max_dist_sqrd float64, is_closed_paths int64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_paths64_strip_near_equal(mem, paths.p, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (paths *PathsD) StripNearEqual(max_dist_sqrd float64, is_closed_paths int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathsd_strip_near_equal(mem, paths.p, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (path *Path64) ToPathd() *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_path64_to_pathd(mem, path.p),
	}

}

func (path *PathD) ToPath64() *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_pathd_to_path64(mem, path.p),
	}

}

func (paths *Paths64) ToPathsd() *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_paths64_to_pathsd(mem, paths.p),
	}

}

func (paths *PathsD) ToPaths64() *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_pathsd_to_paths64(mem, paths.p),
	}

}

func (p *Path64) ScaleToPathD(sx, sy float64) (*PathD, int) {
	var (
		error_code                = C.int(0)
		mem        unsafe.Pointer = C.malloc(C.clipper_pathd_size())
	)

	return &PathD{
		p: C.clipper_scale_path64_to_pathd(mem, p.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (p *PathD) ScaleToPath64(sx, sy float64) (*Path64, int) {
	var (
		error_code                = C.int(0)
		mem        unsafe.Pointer = C.malloc(C.clipper_path64_size())
	)

	return &Path64{
		p: C.clipper_scale_pathd_to_path64(mem, p.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (p *Paths64) ScaleToPathsD(sx, sy float64) (*PathsD, int) {
	var (
		error_code                = C.int(0)
		mem        unsafe.Pointer = C.malloc(C.clipper_pathsd_size())
	)

	return &PathsD{
		p: C.clipper_scale_paths64_to_pathsd(mem, p.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (p *PathsD) ScaleToPaths64(sx, sy float64) (*Paths64, int) {
	var (
		error_code                = C.int(0)
		mem        unsafe.Pointer = C.malloc(C.clipper_paths64_size())
	)

	return &Paths64{
		p: C.clipper_scale_pathsd_to_paths64(mem, p.p, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (pattern *Path64) MinkowskiSum(path Path64, is_closed int64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_path64_minkowski_sum(mem, pattern.p, path.p, C.int(is_closed)),
	}

}

func (pattern *PathD) MinkowskiSum(path PathD, is_closed int64, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathd_minkowski_sum(mem, pattern.p, path.p, C.int(is_closed), C.int(precision)),
	}

}

func (pattern *Path64) MinkowskiDiff(path Path64, is_closed int64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_path64_minkowski_diff(mem, pattern.p, path.p, C.int(is_closed)),
	}

}

func (pattern *PathD) MinkowskiDiff(path PathD, is_closed int64, precision int64) *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_pathd_minkowski_diff(mem, pattern.p, path.p, C.int(is_closed), C.int(precision)),
	}

}

func (a *Point64) Distance(b Point64) float64 {

	return float64(C.clipper_point64_distance(a.p, b.p))
}

func (a *PointD) Distance(b PointD) float64 {

	return float64(C.clipper_pointd_distance(a.p, b.p))
}

func (a *Point64) NearCollinear(b Point64, c Point64, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_point64_near_collinear(a.p, b.p, c.p, C.double(sin_sqrd_min_angle_rads)))
}

func (a *PointD) NearCollinear(b PointD, c PointD, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_pointd_near_collinear(a.p, b.p, c.p, C.double(sin_sqrd_min_angle_rads)))
}

func (path *PathD) Area() float64 {

	return float64(C.clipper_pathd_area(path.p))
}

func (path *Path64) Area() float64 {

	return float64(C.clipper_path64_area(path.p))
}

func (paths *PathsD) Area() float64 {

	return float64(C.clipper_pathsd_area(paths.p))
}

func (paths *Paths64) Area() float64 {

	return float64(C.clipper_paths64_area(paths.p))
}

func (path *PathD) IsPositive() int64 {

	return int64(C.clipper_pathd_is_positive(path.p))
}

func (path *Path64) IsPositive() int64 {

	return int64(C.clipper_path64_is_positive(path.p))
}

func (p *Point64) InPath64(path *Path64) ClipperPointInPolygonResult {
	return ClipperPointInPolygonResult(
		C.clipper_point_in_path64(path.p, p.p),
	)
}

func (p *PointD) InPath64(path *PathD) ClipperPointInPolygonResult {
	return ClipperPointInPolygonResult(
		C.clipper_point_in_pathd(path.p, p.p),
	)
}

func (pt *PolyTree64) Parent() *PolyTree64 {

	return &PolyTree64{
		p: C.clipper_polytree64_parent(pt.p),
	}

}

func (pt *PolyTree64) GetChild(idx int64) *PolyTree64 {

	return &PolyTree64{
		p: C.clipper_polytree64_get_child(pt.p, C.size_t(idx)),
	}

}

func (pt *PolyTree64) AddChild(path Path64) *PolyTree64 {

	return &PolyTree64{
		p: C.clipper_polytree64_add_child(pt.p, path.p),
	}

}

func (pt *PolyTree64) Clear() {

	C.clipper_polytree64_clear(pt.p)
}

func (pt *PolyTree64) Count() int64 {

	return int64(C.clipper_polytree64_count(pt.p))
}

func (pt *PolyTree64) Level() int64 {

	return int64(C.clipper_polytree64_level(pt.p))
}

func (pt *PolyTree64) IsHole() int64 {

	return int64(C.clipper_polytree64_is_hole(pt.p))
}

func (pt *PolyTree64) Polygon() *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_polytree64_polygon(mem, pt.p),
	}

}

func (pt *PolyTree64) Area() float64 {

	return float64(C.clipper_polytree64_area(pt.p))
}

func (pt *PolyTree64) ToPaths() *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_polytree64_to_paths(mem, pt.p),
	}

}

func (pt *PolyTree64) FullyContainsChildren() int64 {

	return int64(C.clipper_polytree64_fully_contains_children(pt.p))
}

func (pt *PolyTreeD) Parent() *PolyTreeD {

	return &PolyTreeD{
		p: C.clipper_polytreed_parent(pt.p),
	}

}

func (pt *PolyTreeD) GetChild(idx int64) *PolyTreeD {

	return &PolyTreeD{
		p: C.clipper_polytreed_get_child(pt.p, C.size_t(idx)),
	}

}

func (pt *PolyTreeD) AddChild(path Path64) *PolyTreeD {

	return &PolyTreeD{
		p: C.clipper_polytreed_add_child(pt.p, path.p),
	}

}

func (pt *PolyTreeD) Clear() {

	C.clipper_polytreed_clear(pt.p)
}

func (pt *PolyTreeD) Count() int64 {

	return int64(C.clipper_polytreed_count(pt.p))
}

func (pt *PolyTreeD) Level() int64 {

	return int64(C.clipper_polytreed_level(pt.p))
}

func (pt *PolyTreeD) IsHole() int64 {

	return int64(C.clipper_polytreed_is_hole(pt.p))
}

func (pt *PolyTreeD) Polygon() *PathD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathd_size())

	return &PathD{
		p: C.clipper_polytreed_polygon(mem, pt.p),
	}

}

func (pt *PolyTreeD) Area() float64 {

	return float64(C.clipper_polytreed_area(pt.p))
}

func (pt *PolyTreeD) ToPaths() *PathsD {
	var mem unsafe.Pointer = C.malloc(C.clipper_pathsd_size())

	return &PathsD{
		p: C.clipper_polytreed_to_paths(mem, pt.p),
	}

}

func (r *Rect64) Width() int64 {

	return int64(C.clipper_rect64_width(r.p))
}

func (r *Rect64) Height() int64 {

	return int64(C.clipper_rect64_height(r.p))
}

func (r *Rect64) Midpoint() Point64 {

	return Point64{
		p: C.clipper_rect64_midpoint(r.p),
	}

}

func (r *Rect64) AsPath() *Path64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_path64_size())

	return &Path64{
		p: C.clipper_rect64_as_path(mem, r.p),
	}

}

func (r *Rect64) ContainsPt(pt Point64) int64 {

	return int64(C.clipper_rect64_contains_pt(r.p, pt.p))
}

func (a *Rect64) ContainsRect(b Rect64) int64 {

	return int64(C.clipper_rect64_contains_rect(a.p, b.p))
}

func (r *Rect64) ScaleMut(scale float64) {

	C.clipper_rect64_scale_mut(r.p, C.double(scale))
}

func (r *Rect64) Scale(scale float64) *Rect64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_rect64_size())

	return &Rect64{
		p: C.clipper_rect64_scale(mem, r.p, C.double(scale)),
	}

}

func (r *Rect64) IsEmpty() int64 {

	return int64(C.clipper_rect64_is_empty(r.p))
}

func (a *Rect64) Intersects(b Rect64) int64 {

	return int64(C.clipper_rect64_intersects(a.p, b.p))
}

func (rect *Rect64) ToStruct() Rect64 {
	return *rect
}

func (c *Clipper64) SetPreserveCollinear(t int64) {

	C.clipper_clipper64_set_preserve_collinear(c.p, C.int(t))
}

func (c *Clipper64) SetReverseSolution(t int64) {

	C.clipper_clipper64_set_reverse_solution(c.p, C.int(t))
}

func (c *Clipper64) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipper64_get_preserve_collinear(c.p))
}

func (c *Clipper64) GetReverseSolution() int64 {

	return int64(C.clipper_clipper64_get_reverse_solution(c.p))
}

func (c *Clipper64) Clear() {

	C.clipper_clipper64_clear(c.p)
}

func (c *ClipperD) SetPreserveCollinear(t int64) {

	C.clipper_clipperd_set_preserve_collinear(c.p, C.int(t))
}

func (c *ClipperD) SetReverseSolution(t int64) {

	C.clipper_clipperd_set_reverse_solution(c.p, C.int(t))
}

func (c *ClipperD) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipperd_get_preserve_collinear(c.p))
}

func (c *ClipperD) GetReverseSolution() int64 {

	return int64(C.clipper_clipperd_get_reverse_solution(c.p))
}

func (c *ClipperD) Clear() {

	C.clipper_clipperd_clear(c.p)
}

func (c *Clipper64) AddSubject(subjects Paths64) {

	C.clipper_clipper64_add_subject(c.p, subjects.p)
}

func (c *Clipper64) AddOpenSubject(open_subjects Paths64) {

	C.clipper_clipper64_add_open_subject(c.p, open_subjects.p)
}

func (c *Clipper64) AddClip(clips Paths64) {

	C.clipper_clipper64_add_clip(c.p, clips.p)
}

func (c64 *Clipper64) Execute(ct ClipperClipType, fr ClipperFillRule, closed Paths64, open Paths64) int64 {

	return int64(C.clipper_clipper64_execute(c64.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.p, open.p))
}

func (c64 *Clipper64) ExecuteTree(ct ClipperClipType, fr ClipperFillRule, tree PolyTree64) int64 {

	return int64(C.clipper_clipper64_execute_tree(c64.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.p))
}

func (c64 *Clipper64) ExecuteTreeWithOpen(ct ClipperClipType, fr ClipperFillRule, tree PolyTree64, open Paths64) int64 {

	return int64(C.clipper_clipper64_execute_tree_with_open(c64.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.p, open.p))
}

func (c *ClipperD) AddSubject(subjects PathsD) {

	C.clipper_clipperd_add_subject(c.p, subjects.p)
}

func (c *ClipperD) AddOpenSubject(open_subjects PathsD) {

	C.clipper_clipperd_add_open_subject(c.p, open_subjects.p)
}

func (c *ClipperD) AddClip(clips PathsD) {

	C.clipper_clipperd_add_clip(c.p, clips.p)
}

func (cD *ClipperD) Execute(ct ClipperClipType, fr ClipperFillRule, closed PathsD, open PathsD) int64 {

	return int64(C.clipper_clipperd_execute(cD.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.p, open.p))
}

func (cD *ClipperD) ExecuteTree(ct ClipperClipType, fr ClipperFillRule, tree PolyTreeD) int64 {

	return int64(C.clipper_clipperd_execute_tree(cD.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.p))
}

func (cD *ClipperD) ExecuteTreeWithOpen(ct ClipperClipType, fr ClipperFillRule, tree PolyTreeD, open PathsD) int64 {

	return int64(C.clipper_clipperd_execute_tree_with_open(cD.p, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.p, open.p))
}

func (c *ClipperOffset) SetMiterLimit(l float64) {

	C.clipper_clipperoffset_set_miter_limit(c.p, C.double(l))
}

func (c *ClipperOffset) SetArcTolerance(t float64) {

	C.clipper_clipperoffset_set_arc_tolerance(c.p, C.double(t))
}

func (c *ClipperOffset) SetPreserveCollinear(t int64) {

	C.clipper_clipperoffset_set_preserve_collinear(c.p, C.int(t))
}

func (c *ClipperOffset) SetReverseSolution(t int64) {

	C.clipper_clipperoffset_set_reverse_solution(c.p, C.int(t))
}

func (c *ClipperOffset) GetMiterLimit() float64 {

	return float64(C.clipper_clipperoffset_get_miter_limit(c.p))
}

func (c *ClipperOffset) GetArcTolerance() float64 {

	return float64(C.clipper_clipperoffset_get_arc_tolerance(c.p))
}

func (c *ClipperOffset) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipperoffset_get_preserve_collinear(c.p))
}

func (c *ClipperOffset) GetReverseSolution() int64 {

	return int64(C.clipper_clipperoffset_get_reverse_solution(c.p))
}

func (c *ClipperOffset) ErrorCode() int64 {

	return int64(C.clipper_clipperoffset_error_code(c.p))
}

func (c *ClipperOffset) Clear() {

	C.clipper_clipperoffset_clear(c.p)
}

func (c *ClipperOffset) AddPath64(p Path64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_path64(c.p, p.p, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperOffset) AddPaths64(p Paths64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_paths64(c.p, p.p, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperOffset) Execute(delta float64) *Paths64 {
	var mem unsafe.Pointer = C.malloc(C.clipper_paths64_size())

	return &Paths64{
		p: C.clipper_clipperoffset_execute(mem, c.p, C.double(delta)),
	}

}
