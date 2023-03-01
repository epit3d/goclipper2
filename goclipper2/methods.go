package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"

func (subjects *paths64) Intersect(clips paths64, fillrule ClipperFillRule) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_intersect(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *paths64) Union(clips paths64, fillrule ClipperFillRule) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_union(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *paths64) Difference(clips paths64, fillrule ClipperFillRule) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_difference(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *paths64) Xor(clips paths64, fillrule ClipperFillRule) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_xor(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *pathsD) Intersect(clips pathsD, fillrule ClipperFillRule, decimal_prec int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_intersect(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *pathsD) Union(clips pathsD, fillrule ClipperFillRule, decimal_prec int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_union(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *pathsD) Difference(clips pathsD, fillrule ClipperFillRule, decimal_prec int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_difference(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *pathsD) Xor(clips pathsD, fillrule ClipperFillRule, decimal_prec int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_xor(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (paths *paths64) Inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_inflate(mem, paths.P, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit)),
	}

}

func (paths *pathsD) Inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64, precision int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_inflate(mem, paths.P, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit), C.int(precision)),
	}

}

func (path *path64) Bounds() *rect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &rect64{
		P: C.clipper_path64_bounds(mem, path.P),
	}

}

func (path *pathD) Bounds() *rectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &rectD{
		P: C.clipper_pathd_bounds(mem, path.P),
	}

}

func (paths *paths64) Bounds() *rect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &rect64{
		P: C.clipper_paths64_bounds(mem, paths.P),
	}

}

func (paths *pathsD) Bounds() *rectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &rectD{
		P: C.clipper_pathsd_bounds(mem, paths.P),
	}

}

func (path *path64) AddPoint(pt point64) {

	C.clipper_path64_add_point(path.P, pt.P)
}

func (path *pathD) AddPoint(pt pointD) {

	C.clipper_pathd_add_point(path.P, pt.P)
}

func (paths *paths64) AddPath(p path64) {

	C.clipper_paths64_add_path(paths.P, p.P)
}

func (paths *pathsD) AddPath(p pathD) {

	C.clipper_pathsd_add_path(paths.P, p.P)
}

func (path *path64) Length() int64 {

	return int64(C.clipper_path64_length(path.P))
}

func (path *pathD) Length() int64 {

	return int64(C.clipper_pathd_length(path.P))
}

func (path *path64) GetPoint(idx int64) point64 {

	return point64{
		P: C.clipper_path64_get_point(path.P, C.int(idx)),
	}

}

func (path *pathD) GetPoint(idx int64) pointD {

	return pointD{
		P: C.clipper_pathd_get_point(path.P, C.int(idx)),
	}

}

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

func (paths *paths64) Length() int64 {

	return int64(C.clipper_paths64_length(paths.P))
}

func (paths *pathsD) Length() int64 {

	return int64(C.clipper_pathsd_length(paths.P))
}

func (paths *paths64) Lengths() []int64 {
	var lengths []int64

	l := paths.Length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.GetPath(i).Length())
	}

	return lengths
}

func (paths *pathsD) Lengths() []int64 {
	var lengths []int64

	l := paths.Length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.GetPath(i).Length())
	}

	return lengths
}

func (paths *paths64) PathLength(idx int64) int64 {

	return int64(C.clipper_paths64_path_length(paths.P, C.int(idx)))
}

func (paths *pathsD) PathLength(idx int64) int64 {

	return int64(C.clipper_pathsd_path_length(paths.P, C.int(idx)))
}

func (paths *paths64) GetPath(idx int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_paths64_get_path(mem, paths.P, C.int(idx)),
	}

}

func (paths *pathsD) GetPath(idx int64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathsd_get_path(mem, paths.P, C.int(idx)),
	}

}

func (paths *paths64) GetPoint(path_idx int64, point_idx int64) point64 {

	return point64{
		P: C.clipper_paths64_get_point(paths.P, C.int(path_idx), C.int(point_idx)),
	}

}

func (paths *pathsD) GetPoint(path_idx int64, point_idx int64) pointD {

	return pointD{
		P: C.clipper_pathsd_get_point(paths.P, C.int(path_idx), C.int(point_idx)),
	}

}

func (path *path64) Translate(dx int64, dy int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_translate(mem, path.P, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (path *pathD) Translate(dx float64, dy float64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_translate(mem, path.P, C.double(dx), C.double(dy)),
	}

}

func (paths *paths64) Translate(dx int64, dy int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_translate(mem, paths.P, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (paths *pathsD) Translate(dx float64, dy float64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_translate(mem, paths.P, C.double(dx), C.double(dy)),
	}

}

func (path *path64) Scale(sx float64, sy float64) (*path64, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &path64{
		P: C.clipper_path64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *pathD) Scale(sx float64, sy float64) (*pathD, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &pathD{
		P: C.clipper_pathd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *paths64) Scale(sx float64, sy float64) (*paths64, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &paths64{
		P: C.clipper_paths64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *pathsD) Scale(sx float64, sy float64) (*pathsD, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &pathsD{
		P: C.clipper_pathsd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}
func (path *path64) TrimCollinear(is_open_path int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_trim_collinear(mem, path.P, C.int(is_open_path)),
	}

}

func (path *pathD) TrimCollinear(precision int64, is_open_path int64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_trim_collinear(mem, path.P, C.int(precision), C.int(is_open_path)),
	}

}

func (path *path64) Simplify(epsilon float64, is_open_path int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_simplify(mem, path.P, C.double(epsilon), C.int(is_open_path)),
	}

}

func (path *pathD) Simplify(epsilon float64, is_open_path int64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_simplify(mem, path.P, C.double(epsilon), C.int(is_open_path)),
	}

}

func (paths *paths64) Simplify(epsilon float64, is_open_paths int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_simplify(mem, paths.P, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (paths *pathsD) Simplify(epsilon float64, is_open_paths int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_simplify(mem, paths.P, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (path *path64) RamerDouglasPeucker(epsilon float64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_ramer_douglas_peucker(mem, path.P, C.double(epsilon)),
	}

}

func (path *pathD) RamerDouglasPeucker(epsilon float64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_ramer_douglas_peucker(mem, path.P, C.double(epsilon)),
	}

}

func (paths *paths64) RamerDouglasPeucker(epsilon float64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_ramer_douglas_peucker(mem, paths.P, C.double(epsilon)),
	}

}

func (paths *pathsD) RamerDouglasPeucker(epsilon float64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_ramer_douglas_peucker(mem, paths.P, C.double(epsilon)),
	}

}

func (path *path64) StripNearEqual(max_dist_sqrd float64, is_closed_path int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_strip_near_equal(mem, path.P, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (path *pathD) StripNearEqual(max_dist_sqrd float64, is_closed_path int64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_strip_near_equal(mem, path.P, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (paths *paths64) StripNearEqual(max_dist_sqrd float64, is_closed_paths int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_strip_near_equal(mem, paths.P, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (paths *pathsD) StripNearEqual(max_dist_sqrd float64, is_closed_paths int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_strip_near_equal(mem, paths.P, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (path *path64) StripDuplicates(is_closed_path int64) *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_path64_strip_duplicates(mem, path.P, C.int(is_closed_path)),
	}

}

func (path *pathD) StripDuplicates(is_closed_path int64) *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_pathd_strip_duplicates(mem, path.P, C.int(is_closed_path)),
	}

}

func (paths *paths64) StripDuplicates(is_closed_paths int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_paths64_strip_duplicates(mem, paths.P, C.int(is_closed_paths)),
	}

}

func (paths *pathsD) StripDuplicates(is_closed_paths int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathsd_strip_duplicates(mem, paths.P, C.int(is_closed_paths)),
	}

}

func (path *path64) ToPathd() *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_path64_to_pathd(mem, path.P),
	}

}

func (path *pathD) ToPath64() *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_pathd_to_path64(mem, path.P),
	}

}

func (paths *paths64) ToPathsd() *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_paths64_to_pathsd(mem, paths.P),
	}

}

func (paths *pathsD) ToPaths64() *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_pathsd_to_paths64(mem, paths.P),
	}

}

func (pattern *path64) MinkowskiSum(path path64, is_closed int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_path64_minkowski_sum(mem, pattern.P, path.P, C.int(is_closed)),
	}

}

func (pattern *pathD) MinkowskiSum(path pathD, is_closed int64, precision int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathd_minkowski_sum(mem, pattern.P, path.P, C.int(is_closed), C.int(precision)),
	}

}

func (pattern *path64) MinkowskiDiff(path path64, is_closed int64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_path64_minkowski_diff(mem, pattern.P, path.P, C.int(is_closed)),
	}

}

func (pattern *pathD) MinkowskiDiff(path pathD, is_closed int64, precision int64) *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_pathd_minkowski_diff(mem, pattern.P, path.P, C.int(is_closed), C.int(precision)),
	}

}

func (a *point64) Distance(b point64) float64 {

	return float64(C.clipper_point64_distance(a.P, b.P))
}

func (a *pointD) Distance(b pointD) float64 {

	return float64(C.clipper_pointd_distance(a.P, b.P))
}

func (a *point64) NearCollinear(b point64, c point64, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_point64_near_collinear(a.P, b.P, c.P, C.double(sin_sqrd_min_angle_rads)))
}

func (a *pointD) NearCollinear(b pointD, c pointD, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_pointd_near_collinear(a.P, b.P, c.P, C.double(sin_sqrd_min_angle_rads)))
}

func (path *pathD) Area() float64 {

	return float64(C.clipper_pathd_area(path.P))
}

func (path *path64) Area() float64 {

	return float64(C.clipper_path64_area(path.P))
}

func (paths *pathsD) Area() float64 {

	return float64(C.clipper_pathsd_area(paths.P))
}

func (paths *paths64) Area() float64 {

	return float64(C.clipper_paths64_area(paths.P))
}

func (path *pathD) IsPositive() int64 {

	return int64(C.clipper_pathd_is_positive(path.P))
}

func (path *path64) IsPositive() int64 {

	return int64(C.clipper_path64_is_positive(path.P))
}

func (pt *polyTree64) Parent() *polyTree64 {

	return &polyTree64{
		P: C.clipper_polytree64_parent(pt.P),
	}

}

func (pt *polyTree64) GetChild(idx int64) *polyTree64 {

	return &polyTree64{
		P: C.clipper_polytree64_get_child(pt.P, C.size_t(idx)),
	}

}

func (pt *polyTree64) AddChild(path path64) *polyTree64 {

	return &polyTree64{
		P: C.clipper_polytree64_add_child(pt.P, path.P),
	}

}

func (pt *polyTree64) Clear() {

	C.clipper_polytree64_clear(pt.P)
}

func (pt *polyTree64) Count() int64 {

	return int64(C.clipper_polytree64_count(pt.P))
}

func (pt *polyTree64) Level() int64 {

	return int64(C.clipper_polytree64_level(pt.P))
}

func (pt *polyTree64) IsHole() int64 {

	return int64(C.clipper_polytree64_is_hole(pt.P))
}

func (pt *polyTree64) Polygon() *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_polytree64_polygon(mem, pt.P),
	}

}

func (pt *polyTree64) Area() float64 {

	return float64(C.clipper_polytree64_area(pt.P))
}

func (pt *polyTree64) ToPaths() *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_polytree64_to_paths(mem, pt.P),
	}

}

func (pt *polyTree64) FullyContainsChildren() int64 {

	return int64(C.clipper_polytree64_fully_contains_children(pt.P))
}

func (pt *polyTreeD) Parent() *polyTreeD {

	return &polyTreeD{
		P: C.clipper_polytreed_parent(pt.P),
	}

}

func (pt *polyTreeD) GetChild(idx int64) *polyTreeD {

	return &polyTreeD{
		P: C.clipper_polytreed_get_child(pt.P, C.size_t(idx)),
	}

}

func (pt *polyTreeD) SetInvScale(value float64) {

	C.clipper_polytreed_set_inv_scale(pt.P, C.double(value))
}

func (pt *polyTreeD) InvScale() float64 {

	return float64(C.clipper_polytreed_inv_scale(pt.P))
}

func (pt *polyTreeD) AddChild(path path64) *polyTreeD {

	return &polyTreeD{
		P: C.clipper_polytreed_add_child(pt.P, path.P),
	}

}

func (pt *polyTreeD) Clear() {

	C.clipper_polytreed_clear(pt.P)
}

func (pt *polyTreeD) Count() int64 {

	return int64(C.clipper_polytreed_count(pt.P))
}

func (pt *polyTreeD) Level() int64 {

	return int64(C.clipper_polytreed_level(pt.P))
}

func (pt *polyTreeD) IsHole() int64 {

	return int64(C.clipper_polytreed_is_hole(pt.P))
}

func (pt *polyTreeD) Polygon() *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_polytreed_polygon(mem, pt.P),
	}

}

func (pt *polyTreeD) Area() float64 {

	return float64(C.clipper_polytreed_area(pt.P))
}

func (pt *polyTreeD) ToPaths() *pathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathsD{
		P: C.clipper_polytreed_to_paths(mem, pt.P),
	}

}

func (r *rect64) Width() int64 {

	return int64(C.clipper_rect64_width(r.P))
}

func (r *rect64) Height() int64 {

	return int64(C.clipper_rect64_height(r.P))
}

func (r *rect64) Midpoint() point64 {

	return point64{
		P: C.clipper_rect64_midpoint(r.P),
	}

}

func (r *rect64) AsPath() *path64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &path64{
		P: C.clipper_rect64_as_path(mem, r.P),
	}

}

func (r *rect64) ContainsPt(pt point64) int64 {

	return int64(C.clipper_rect64_contains_pt(r.P, pt.P))
}

func (a *rect64) ContainsRect(b rect64) int64 {

	return int64(C.clipper_rect64_contains_rect(a.P, b.P))
}

func (r *rect64) ScaleMut(scale float64) {

	C.clipper_rect64_scale_mut(r.P, C.double(scale))
}

func (r *rect64) Scale(scale float64) *rect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &rect64{
		P: C.clipper_rect64_scale(mem, r.P, C.double(scale)),
	}

}

func (r *rect64) IsEmpty() int64 {

	return int64(C.clipper_rect64_is_empty(r.P))
}

func (a *rect64) Intersects(b rect64) int64 {

	return int64(C.clipper_rect64_intersects(a.P, b.P))
}

func (r *rectD) Width() float64 {

	return float64(C.clipper_rectd_width(r.P))
}

func (r *rectD) Height() float64 {

	return float64(C.clipper_rectd_height(r.P))
}

func (r *rectD) Midpoint() pointD {

	return pointD{
		P: C.clipper_rectd_midpoint(r.P),
	}

}

func (r *rectD) AsPath() *pathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &pathD{
		P: C.clipper_rectd_as_path(mem, r.P),
	}

}

func (r *rectD) ContainsPt(pt pointD) int64 {

	return int64(C.clipper_rectd_contains_pt(r.P, pt.P))
}

func (a *rectD) ContainsRect(b rectD) int64 {

	return int64(C.clipper_rectd_contains_rect(a.P, b.P))
}

func (r *rectD) ScaleMut(scale float64) {

	C.clipper_rectd_scale_mut(r.P, C.double(scale))
}

func (r *rectD) Scale(scale float64) *rectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &rectD{
		P: C.clipper_rectd_scale(mem, r.P, C.double(scale)),
	}

}

func (r *rectD) IsEmpty() int64 {

	return int64(C.clipper_rectd_is_empty(r.P))
}

func (a *rectD) Intersects(b rectD) int64 {

	return int64(C.clipper_rectd_intersects(a.P, b.P))
}

func (rect *rect64) ToStruct() rect64 {
	return *rect
}

func (rect *rectD) ToStruct() rectD {
	return *rect
}

func (c *clipper64) SetPreserveCollinear(t int64) {

	C.clipper_clipper64_set_preserve_collinear(c.P, C.int(t))
}

func (c *clipper64) SetReverseSolution(t int64) {

	C.clipper_clipper64_set_reverse_solution(c.P, C.int(t))
}

func (c *clipper64) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipper64_get_preserve_collinear(c.P))
}

func (c *clipper64) GetReverseSolution() int64 {

	return int64(C.clipper_clipper64_get_reverse_solution(c.P))
}

func (c *clipper64) Clear() {

	C.clipper_clipper64_clear(c.P)
}

func (c *clipperD) SetPreserveCollinear(t int64) {

	C.clipper_clipperd_set_preserve_collinear(c.P, C.int(t))
}

func (c *clipperD) SetReverseSolution(t int64) {

	C.clipper_clipperd_set_reverse_solution(c.P, C.int(t))
}

func (c *clipperD) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipperd_get_preserve_collinear(c.P))
}

func (c *clipperD) GetReverseSolution() int64 {

	return int64(C.clipper_clipperd_get_reverse_solution(c.P))
}

func (c *clipperD) Clear() {

	C.clipper_clipperd_clear(c.P)
}

func (c *clipper64) AddSubject(subjects paths64) {

	C.clipper_clipper64_add_subject(c.P, subjects.P)
}

func (c *clipper64) AddOpenSubject(open_subjects paths64) {

	C.clipper_clipper64_add_open_subject(c.P, open_subjects.P)
}

func (c *clipper64) AddClip(clips paths64) {

	C.clipper_clipper64_add_clip(c.P, clips.P)
}

func (c64 *clipper64) Execute(ct ClipperClipType, fr ClipperFillRule, closed paths64, open paths64) int64 {

	return int64(C.clipper_clipper64_execute(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.P, open.P))
}

func (c64 *clipper64) ExecuteTree(ct ClipperClipType, fr ClipperFillRule, tree polyTree64) int64 {

	return int64(C.clipper_clipper64_execute_tree(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P))
}

func (c64 *clipper64) ExecuteTreeWithOpen(ct ClipperClipType, fr ClipperFillRule, tree polyTree64, open paths64) int64 {

	return int64(C.clipper_clipper64_execute_tree_with_open(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P, open.P))
}

func (c *clipperD) AddSubject(subjects pathsD) {

	C.clipper_clipperd_add_subject(c.P, subjects.P)
}

func (c *clipperD) AddOpenSubject(open_subjects pathsD) {

	C.clipper_clipperd_add_open_subject(c.P, open_subjects.P)
}

func (c *clipperD) AddClip(clips pathsD) {

	C.clipper_clipperd_add_clip(c.P, clips.P)
}

func (cD *clipperD) Execute(ct ClipperClipType, fr ClipperFillRule, closed pathsD, open pathsD) int64 {

	return int64(C.clipper_clipperd_execute(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.P, open.P))
}

func (cD *clipperD) ExecuteTree(ct ClipperClipType, fr ClipperFillRule, tree polyTreeD) int64 {

	return int64(C.clipper_clipperd_execute_tree(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P))
}

func (cD *clipperD) ExecuteTreeWithOpen(ct ClipperClipType, fr ClipperFillRule, tree polyTreeD, open pathsD) int64 {

	return int64(C.clipper_clipperd_execute_tree_with_open(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P, open.P))
}

func (c *clipperOffset) SetMiterLimit(l float64) {

	C.clipper_clipperoffset_set_miter_limit(c.P, C.double(l))
}

func (c *clipperOffset) SetArcTolerance(t float64) {

	C.clipper_clipperoffset_set_arc_tolerance(c.P, C.double(t))
}

func (c *clipperOffset) SetPreserveCollinear(t int64) {

	C.clipper_clipperoffset_set_preserve_collinear(c.P, C.int(t))
}

func (c *clipperOffset) SetReverseSolution(t int64) {

	C.clipper_clipperoffset_set_reverse_solution(c.P, C.int(t))
}

func (c *clipperOffset) GetMiterLimit() float64 {

	return float64(C.clipper_clipperoffset_get_miter_limit(c.P))
}

func (c *clipperOffset) GetArcTolerance() float64 {

	return float64(C.clipper_clipperoffset_get_arc_tolerance(c.P))
}

func (c *clipperOffset) GetPreserveCollinear() int64 {

	return int64(C.clipper_clipperoffset_get_preserve_collinear(c.P))
}

func (c *clipperOffset) GetReverseSolution() int64 {

	return int64(C.clipper_clipperoffset_get_reverse_solution(c.P))
}

func (c *clipperOffset) ErrorCode() int64 {

	return int64(C.clipper_clipperoffset_error_code(c.P))
}

func (c *clipperOffset) Clear() {

	C.clipper_clipperoffset_clear(c.P)
}

func (c *clipperOffset) AddPathd(p pathD, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_pathd(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *clipperOffset) AddPathsd(p pathsD, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_pathsd(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *clipperOffset) AddPath64(p path64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_path64(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *clipperOffset) AddPaths64(p paths64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_paths64(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *clipperOffset) Execute(delta float64) *paths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &paths64{
		P: C.clipper_clipperoffset_execute(mem, c.P, C.double(delta)),
	}

}
