package goclipper2

// #cgo LDFLAGS: /usr/local/lib/libclipper2c.so
// #include "clipper2c/clipper2c.h"
import "C"
import "unsafe"

func (subjects *ClipperPaths64) Clipper_paths64_intersect(clips ClipperPaths64, fillrule ClipperFillRule) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_intersect(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *ClipperPaths64) Clipper_paths64_union(clips ClipperPaths64, fillrule ClipperFillRule) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_union(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *ClipperPaths64) Clipper_paths64_difference(clips ClipperPaths64, fillrule ClipperFillRule) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_difference(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *ClipperPaths64) Clipper_paths64_xor(clips ClipperPaths64, fillrule ClipperFillRule) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_xor(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule)),
	}

}

func (subjects *ClipperPathsD) Clipper_pathsd_intersect(clips ClipperPathsD, fillrule ClipperFillRule, decimal_prec int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_intersect(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *ClipperPathsD) Clipper_pathsd_union(clips ClipperPathsD, fillrule ClipperFillRule, decimal_prec int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_union(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *ClipperPathsD) Clipper_pathsd_difference(clips ClipperPathsD, fillrule ClipperFillRule, decimal_prec int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_difference(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (subjects *ClipperPathsD) Clipper_pathsd_xor(clips ClipperPathsD, fillrule ClipperFillRule, decimal_prec int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_xor(mem, subjects.P, clips.P, C.ClipperFillRule(fillrule), C.int(decimal_prec)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_inflate(mem, paths.P, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_inflate(delta float64, jt ClipperJoinType, et ClipperEndType, miter_limit float64, precision int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_inflate(mem, paths.P, C.double(delta), C.ClipperJoinType(jt), C.ClipperEndType(et), C.double(miter_limit), C.int(precision)),
	}

}

func (path *ClipperPath64) Clipper_path64_bounds() *ClipperRect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRect64{
		P: C.clipper_path64_bounds(mem, path.P),
	}

}

func (path *ClipperPathD) Clipper_pathd_bounds() *ClipperRectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRectD{
		P: C.clipper_pathd_bounds(mem, path.P),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_bounds() *ClipperRect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRect64{
		P: C.clipper_paths64_bounds(mem, paths.P),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_bounds() *ClipperRectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRectD{
		P: C.clipper_pathsd_bounds(mem, paths.P),
	}

}

func (path *ClipperPath64) Clipper_path64_add_point(pt ClipperPoint64) {

	C.clipper_path64_add_point(path.P, pt.P)
}

func (path *ClipperPathD) Clipper_pathd_add_point(pt ClipperPointD) {

	C.clipper_pathd_add_point(path.P, pt.P)
}

func (paths *ClipperPaths64) Clipper_paths64_add_path(p ClipperPath64) {

	C.clipper_paths64_add_path(paths.P, p.P)
}

func (paths *ClipperPathsD) Clipper_pathsd_add_path(p ClipperPathD) {

	C.clipper_pathsd_add_path(paths.P, p.P)
}

func (path *ClipperPath64) Clipper_path64_length() int64 {

	return int64(C.clipper_path64_length(path.P))
}

func (path *ClipperPathD) Clipper_pathd_length() int64 {

	return int64(C.clipper_pathd_length(path.P))
}

func (path *ClipperPath64) Clipper_path64_get_point(idx int64) ClipperPoint64 {

	return ClipperPoint64{
		P: C.clipper_path64_get_point(path.P, C.int(idx)),
	}

}

func (path *ClipperPathD) Clipper_pathd_get_point(idx int64) ClipperPointD {

	return ClipperPointD{
		P: C.clipper_pathd_get_point(path.P, C.int(idx)),
	}

}

func (path *ClipperPath64) Clipper_path64_to_points() []ClipperPoint64 {
	var res []ClipperPoint64
	l := path.Clipper_path64_length()
	var i int64

	for i = 0; i < l; i++ {
		n := path.Clipper_path64_get_point(i)
		res = append(res, n)
	}
	return res
}

func (path *ClipperPathD) Clipper_pathd_to_points() []ClipperPointD {
	var res []ClipperPointD
	l := path.Clipper_pathd_length()
	var i int64

	for i = 0; i < l; i++ {
		n := path.Clipper_pathd_get_point(i)
		res = append(res, n)
	}
	return res
}

func (paths *ClipperPaths64) Clipper_paths64_length() int64 {

	return int64(C.clipper_paths64_length(paths.P))
}

func (paths *ClipperPathsD) Clipper_pathsd_length() int64 {

	return int64(C.clipper_pathsd_length(paths.P))
}

func (paths *ClipperPaths64) Clipper_paths64_lengths() []int64 {
	var lengths []int64

	l := paths.Clipper_paths64_length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.Clipper_paths64_get_path(i).Clipper_path64_length())
	}

	return lengths
}

func (paths *ClipperPathsD) Clipper_pathsd_lengths() []int64 {
	var lengths []int64

	l := paths.Clipper_pathsd_length()

	var i int64
	for i = 0; i < l; i++ {
		lengths = append(lengths, paths.Clipper_pathsd_get_path(i).Clipper_pathd_length())
	}

	return lengths
}

func (paths *ClipperPaths64) Clipper_paths64_path_length(idx int64) int64 {

	return int64(C.clipper_paths64_path_length(paths.P, C.int(idx)))
}

func (paths *ClipperPathsD) Clipper_pathsd_path_length(idx int64) int64 {

	return int64(C.clipper_pathsd_path_length(paths.P, C.int(idx)))
}

func (paths *ClipperPaths64) Clipper_paths64_get_path(idx int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_paths64_get_path(mem, paths.P, C.int(idx)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_get_path(idx int64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathsd_get_path(mem, paths.P, C.int(idx)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_get_point(path_idx int64, point_idx int64) ClipperPoint64 {

	return ClipperPoint64{
		P: C.clipper_paths64_get_point(paths.P, C.int(path_idx), C.int(point_idx)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_get_point(path_idx int64, point_idx int64) ClipperPointD {

	return ClipperPointD{
		P: C.clipper_pathsd_get_point(paths.P, C.int(path_idx), C.int(point_idx)),
	}

}

func (path *ClipperPath64) Clipper_path64_translate(dx int64, dy int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_translate(mem, path.P, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (path *ClipperPathD) Clipper_pathd_translate(dx float64, dy float64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_translate(mem, path.P, C.double(dx), C.double(dy)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_translate(dx int64, dy int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_translate(mem, paths.P, C.int64_t(dx), C.int64_t(dy)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_translate(dx float64, dy float64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_translate(mem, paths.P, C.double(dx), C.double(dy)),
	}

}

func (path *ClipperPath64) Clipper_path64_scale(sx float64, sy float64) (*ClipperPath64, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &ClipperPath64{
		P: C.clipper_path64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (path *ClipperPathD) Clipper_pathd_scale(sx float64, sy float64) (*ClipperPathD, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &ClipperPathD{
		P: C.clipper_pathd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (path *ClipperPaths64) Clipper_paths64_scale(sx float64, sy float64) (*ClipperPaths64, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (path *ClipperPathsD) Clipper_pathsd_scale(sx float64, sy float64) (*ClipperPathsD, int) {
	mem := C.malloc(0)
	error_code := C.int(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_scale(mem, path.P, C.double(sx), C.double(sy), &error_code),
	}, int(error_code)
}

func (path *ClipperPath64) Clipper_path64_trim_collinear(is_open_path int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_trim_collinear(mem, path.P, C.int(is_open_path)),
	}

}

func (path *ClipperPathD) Clipper_pathd_trim_collinear(precision int64, is_open_path int64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_trim_collinear(mem, path.P, C.int(precision), C.int(is_open_path)),
	}

}

func (path *ClipperPath64) Clipper_path64_simplify(epsilon float64, is_open_path int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_simplify(mem, path.P, C.double(epsilon), C.int(is_open_path)),
	}

}

func (path *ClipperPathD) Clipper_pathd_simplify(epsilon float64, is_open_path int64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_simplify(mem, path.P, C.double(epsilon), C.int(is_open_path)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_simplify(epsilon float64, is_open_paths int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_simplify(mem, paths.P, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_simplify(epsilon float64, is_open_paths int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_simplify(mem, paths.P, C.double(epsilon), C.int(is_open_paths)),
	}

}

func (path *ClipperPath64) Clipper_path64_ramer_douglas_peucker(epsilon float64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_ramer_douglas_peucker(mem, path.P, C.double(epsilon)),
	}

}

func (path *ClipperPathD) Clipper_pathd_ramer_douglas_peucker(epsilon float64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_ramer_douglas_peucker(mem, path.P, C.double(epsilon)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_ramer_douglas_peucker(epsilon float64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_ramer_douglas_peucker(mem, paths.P, C.double(epsilon)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_ramer_douglas_peucker(epsilon float64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_ramer_douglas_peucker(mem, paths.P, C.double(epsilon)),
	}

}

func (path *ClipperPath64) Clipper_path64_strip_near_equal(max_dist_sqrd float64, is_closed_path int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_strip_near_equal(mem, path.P, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (path *ClipperPathD) Clipper_pathd_strip_near_equal(max_dist_sqrd float64, is_closed_path int64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_strip_near_equal(mem, path.P, C.double(max_dist_sqrd), C.int(is_closed_path)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_strip_near_equal(max_dist_sqrd float64, is_closed_paths int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_strip_near_equal(mem, paths.P, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_strip_near_equal(max_dist_sqrd float64, is_closed_paths int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_strip_near_equal(mem, paths.P, C.double(max_dist_sqrd), C.int(is_closed_paths)),
	}

}

func (path *ClipperPath64) Clipper_path64_strip_duplicates(is_closed_path int64) *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_path64_strip_duplicates(mem, path.P, C.int(is_closed_path)),
	}

}

func (path *ClipperPathD) Clipper_pathd_strip_duplicates(is_closed_path int64) *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_pathd_strip_duplicates(mem, path.P, C.int(is_closed_path)),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_strip_duplicates(is_closed_paths int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_paths64_strip_duplicates(mem, paths.P, C.int(is_closed_paths)),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_strip_duplicates(is_closed_paths int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathsd_strip_duplicates(mem, paths.P, C.int(is_closed_paths)),
	}

}

func (path *ClipperPath64) Clipper_path64_to_pathd() *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_path64_to_pathd(mem, path.P),
	}

}

func (path *ClipperPathD) Clipper_pathd_to_pathd() *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_pathd_to_pathd(mem, path.P),
	}

}

func (paths *ClipperPaths64) Clipper_paths64_to_pathsd() *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_paths64_to_pathsd(mem, paths.P),
	}

}

func (paths *ClipperPathsD) Clipper_pathsd_to_pathsd() *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_pathsd_to_pathsd(mem, paths.P),
	}

}

func (pattern *ClipperPath64) Clipper_path64_minkowski_sum(path ClipperPath64, is_closed int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_path64_minkowski_sum(mem, pattern.P, path.P, C.int(is_closed)),
	}

}

func (pattern *ClipperPathD) Clipper_pathd_minkowski_sum(path ClipperPathD, is_closed int64, precision int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathd_minkowski_sum(mem, pattern.P, path.P, C.int(is_closed), C.int(precision)),
	}

}

func (pattern *ClipperPath64) Clipper_path64_minkowski_diff(path ClipperPath64, is_closed int64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_path64_minkowski_diff(mem, pattern.P, path.P, C.int(is_closed)),
	}

}

func (pattern *ClipperPathD) Clipper_pathd_minkowski_diff(path ClipperPathD, is_closed int64, precision int64) *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_pathd_minkowski_diff(mem, pattern.P, path.P, C.int(is_closed), C.int(precision)),
	}

}

func (a *ClipperPoint64) Clipper_point64_distance(b ClipperPoint64) float64 {

	return float64(C.clipper_point64_distance(a.P, b.P))
}

func (a *ClipperPointD) Clipper_pointd_distance(b ClipperPointD) float64 {

	return float64(C.clipper_pointd_distance(a.P, b.P))
}

func (a *ClipperPoint64) Clipper_point64_near_collinear(b ClipperPoint64, c ClipperPoint64, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_point64_near_collinear(a.P, b.P, c.P, C.double(sin_sqrd_min_angle_rads)))
}

func (a *ClipperPointD) Clipper_pointd_near_collinear(b ClipperPointD, c ClipperPointD, sin_sqrd_min_angle_rads float64) int64 {

	return int64(C.clipper_pointd_near_collinear(a.P, b.P, c.P, C.double(sin_sqrd_min_angle_rads)))
}

func (path *ClipperPathD) Clipper_pathd_area() float64 {

	return float64(C.clipper_pathd_area(path.P))
}

func (path *ClipperPath64) Clipper_path64_area() float64 {

	return float64(C.clipper_path64_area(path.P))
}

func (paths *ClipperPathsD) Clipper_pathsd_area() float64 {

	return float64(C.clipper_pathsd_area(paths.P))
}

func (paths *ClipperPaths64) Clipper_paths64_area() float64 {

	return float64(C.clipper_paths64_area(paths.P))
}

func (path *ClipperPathD) Clipper_pathd_is_positive() int64 {

	return int64(C.clipper_pathd_is_positive(path.P))
}

func (path *ClipperPath64) Clipper_path64_is_positive() int64 {

	return int64(C.clipper_path64_is_positive(path.P))
}

func (pt *ClipperPolyTree64) Clipper_polytree64_parent() *ClipperPolyTree64 {

	return &ClipperPolyTree64{
		P: C.clipper_polytree64_parent(pt.P),
	}

}

func (pt *ClipperPolyTree64) Clipper_polytree64_get_child(idx int64) *ClipperPolyTree64 {

	return &ClipperPolyTree64{
		P: C.clipper_polytree64_get_child(pt.P, C.size_t(idx)),
	}

}

func (pt *ClipperPolyTree64) Clipper_polytree64_add_child(path ClipperPath64) *ClipperPolyTree64 {

	return &ClipperPolyTree64{
		P: C.clipper_polytree64_add_child(pt.P, path.P),
	}

}

func (pt *ClipperPolyTree64) Clipper_polytree64_clear() {

	C.clipper_polytree64_clear(pt.P)
}

func (pt *ClipperPolyTree64) Clipper_polytree64_count() int64 {

	return int64(C.clipper_polytree64_count(pt.P))
}

func (pt *ClipperPolyTree64) Clipper_polytree64_level() int64 {

	return int64(C.clipper_polytree64_level(pt.P))
}

func (pt *ClipperPolyTree64) Clipper_polytree64_is_hole() int64 {

	return int64(C.clipper_polytree64_is_hole(pt.P))
}

func (pt *ClipperPolyTree64) Clipper_polytree64_polygon() *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_polytree64_polygon(mem, pt.P),
	}

}

func (pt *ClipperPolyTree64) Clipper_polytree64_area() float64 {

	return float64(C.clipper_polytree64_area(pt.P))
}

func (pt *ClipperPolyTree64) Clipper_polytree64_to_paths() *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_polytree64_to_paths(mem, pt.P),
	}

}

func (pt *ClipperPolyTree64) Clipper_polytree64_fully_contains_children() int64 {

	return int64(C.clipper_polytree64_fully_contains_children(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_parent() *ClipperPolyTreeD {

	return &ClipperPolyTreeD{
		P: C.clipper_polytreed_parent(pt.P),
	}

}

func (pt *ClipperPolyTreeD) Clipper_polytreed_get_child(idx int64) *ClipperPolyTreeD {

	return &ClipperPolyTreeD{
		P: C.clipper_polytreed_get_child(pt.P, C.size_t(idx)),
	}

}

func (pt *ClipperPolyTreeD) Clipper_polytreed_set_inv_scale(value float64) {

	C.clipper_polytreed_set_inv_scale(pt.P, C.double(value))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_inv_scale() float64 {

	return float64(C.clipper_polytreed_inv_scale(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_add_child(path ClipperPath64) *ClipperPolyTreeD {

	return &ClipperPolyTreeD{
		P: C.clipper_polytreed_add_child(pt.P, path.P),
	}

}

func (pt *ClipperPolyTreeD) Clipper_polytreed_clear() {

	C.clipper_polytreed_clear(pt.P)
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_count() int64 {

	return int64(C.clipper_polytreed_count(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_level() int64 {

	return int64(C.clipper_polytreed_level(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_is_hole() int64 {

	return int64(C.clipper_polytreed_is_hole(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_polygon() *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_polytreed_polygon(mem, pt.P),
	}

}

func (pt *ClipperPolyTreeD) Clipper_polytreed_area() float64 {

	return float64(C.clipper_polytreed_area(pt.P))
}

func (pt *ClipperPolyTreeD) Clipper_polytreed_to_paths() *ClipperPathsD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathsD{
		P: C.clipper_polytreed_to_paths(mem, pt.P),
	}

}

func (r *ClipperRect64) Clipper_rect64_width() int64 {

	return int64(C.clipper_rect64_width(r.P))
}

func (r *ClipperRect64) Clipper_rect64_height() int64 {

	return int64(C.clipper_rect64_height(r.P))
}

func (r *ClipperRect64) Clipper_rect64_midpoint() ClipperPoint64 {

	return ClipperPoint64{
		P: C.clipper_rect64_midpoint(r.P),
	}

}

func (r *ClipperRect64) Clipper_rect64_as_path() *ClipperPath64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPath64{
		P: C.clipper_rect64_as_path(mem, r.P),
	}

}

func (r *ClipperRect64) Clipper_rect64_contains_pt(pt ClipperPoint64) int64 {

	return int64(C.clipper_rect64_contains_pt(r.P, pt.P))
}

func (a *ClipperRect64) Clipper_rect64_contains_rect(b ClipperRect64) int64 {

	return int64(C.clipper_rect64_contains_rect(a.P, b.P))
}

func (r *ClipperRect64) Clipper_rect64_scale_mut(scale float64) {

	C.clipper_rect64_scale_mut(r.P, C.double(scale))
}

func (r *ClipperRect64) Clipper_rect64_scale(scale float64) *ClipperRect64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRect64{
		P: C.clipper_rect64_scale(mem, r.P, C.double(scale)),
	}

}

func (r *ClipperRect64) Clipper_rect64_is_empty() int64 {

	return int64(C.clipper_rect64_is_empty(r.P))
}

func (a *ClipperRect64) Clipper_rect64_intersects(b ClipperRect64) int64 {

	return int64(C.clipper_rect64_intersects(a.P, b.P))
}

func (r *ClipperRectD) Clipper_rectd_width() float64 {

	return float64(C.clipper_rectd_width(r.P))
}

func (r *ClipperRectD) Clipper_rectd_height() float64 {

	return float64(C.clipper_rectd_height(r.P))
}

func (r *ClipperRectD) Clipper_rectd_midpoint() ClipperPointD {

	return ClipperPointD{
		P: C.clipper_rectd_midpoint(r.P),
	}

}

func (r *ClipperRectD) Clipper_rectd_as_path() *ClipperPathD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPathD{
		P: C.clipper_rectd_as_path(mem, r.P),
	}

}

func (r *ClipperRectD) Clipper_rectd_contains_pt(pt ClipperPointD) int64 {

	return int64(C.clipper_rectd_contains_pt(r.P, pt.P))
}

func (a *ClipperRectD) Clipper_rectd_contains_rect(b ClipperRectD) int64 {

	return int64(C.clipper_rectd_contains_rect(a.P, b.P))
}

func (r *ClipperRectD) Clipper_rectd_scale_mut(scale float64) {

	C.clipper_rectd_scale_mut(r.P, C.double(scale))
}

func (r *ClipperRectD) Clipper_rectd_scale(scale float64) *ClipperRectD {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperRectD{
		P: C.clipper_rectd_scale(mem, r.P, C.double(scale)),
	}

}

func (r *ClipperRectD) Clipper_rectd_is_empty() int64 {

	return int64(C.clipper_rectd_is_empty(r.P))
}

func (a *ClipperRectD) Clipper_rectd_intersects(b ClipperRectD) int64 {

	return int64(C.clipper_rectd_intersects(a.P, b.P))
}

func (rect *ClipperRect64) Clipper_rect64_to_struct() ClipperRect64 {
	return *rect
}

func (rect *ClipperRectD) Clipper_rectd_to_struct() ClipperRectD {
	return *rect
}

func (c *ClipperClipper64) Clipper_clipper64_set_preserve_collinear(t int64) {

	C.clipper_clipper64_set_preserve_collinear(c.P, C.int(t))
}

func (c *ClipperClipper64) Clipper_clipper64_set_reverse_solution(t int64) {

	C.clipper_clipper64_set_reverse_solution(c.P, C.int(t))
}

func (c *ClipperClipper64) Clipper_clipper64_get_preserve_collinear() int64 {

	return int64(C.clipper_clipper64_get_preserve_collinear(c.P))
}

func (c *ClipperClipper64) Clipper_clipper64_get_reverse_solution() int64 {

	return int64(C.clipper_clipper64_get_reverse_solution(c.P))
}

func (c *ClipperClipper64) Clipper_clipper64_clear() {

	C.clipper_clipper64_clear(c.P)
}

func (c *ClipperClipperD) Clipper_clipperd_set_preserve_collinear(t int64) {

	C.clipper_clipperd_set_preserve_collinear(c.P, C.int(t))
}

func (c *ClipperClipperD) Clipper_clipperd_set_reverse_solution(t int64) {

	C.clipper_clipperd_set_reverse_solution(c.P, C.int(t))
}

func (c *ClipperClipperD) Clipper_clipperd_get_preserve_collinear() int64 {

	return int64(C.clipper_clipperd_get_preserve_collinear(c.P))
}

func (c *ClipperClipperD) Clipper_clipperd_get_reverse_solution() int64 {

	return int64(C.clipper_clipperd_get_reverse_solution(c.P))
}

func (c *ClipperClipperD) Clipper_clipperd_clear() {

	C.clipper_clipperd_clear(c.P)
}

func (c *ClipperClipper64) Clipper_clipper64_add_subject(subjects ClipperPaths64) {

	C.clipper_clipper64_add_subject(c.P, subjects.P)
}

func (c *ClipperClipper64) Clipper_clipper64_add_open_subject(open_subjects ClipperPaths64) {

	C.clipper_clipper64_add_open_subject(c.P, open_subjects.P)
}

func (c *ClipperClipper64) Clipper_clipper64_add_clip(clips ClipperPaths64) {

	C.clipper_clipper64_add_clip(c.P, clips.P)
}

func (c64 *ClipperClipper64) Clipper_clipper64_execute(ct ClipperClipType, fr ClipperFillRule, closed ClipperPaths64, open ClipperPaths64) int64 {

	return int64(C.clipper_clipper64_execute(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.P, open.P))
}

func (c64 *ClipperClipper64) Clipper_clipper64_execute_tree(ct ClipperClipType, fr ClipperFillRule, tree ClipperPolyTree64) int64 {

	return int64(C.clipper_clipper64_execute_tree(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P))
}

func (c64 *ClipperClipper64) Clipper_clipper64_execute_tree_with_open(ct ClipperClipType, fr ClipperFillRule, tree ClipperPolyTree64, open ClipperPaths64) int64 {

	return int64(C.clipper_clipper64_execute_tree_with_open(c64.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P, open.P))
}

func (c *ClipperClipperD) Clipper_clipperd_add_subject(subjects ClipperPathsD) {

	C.clipper_clipperd_add_subject(c.P, subjects.P)
}

func (c *ClipperClipperD) Clipper_clipperd_add_open_subject(open_subjects ClipperPathsD) {

	C.clipper_clipperd_add_open_subject(c.P, open_subjects.P)
}

func (c *ClipperClipperD) Clipper_clipperd_add_clip(clips ClipperPathsD) {

	C.clipper_clipperd_add_clip(c.P, clips.P)
}

func (cD *ClipperClipperD) Clipper_clipperd_execute(ct ClipperClipType, fr ClipperFillRule, closed ClipperPathsD, open ClipperPathsD) int64 {

	return int64(C.clipper_clipperd_execute(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), closed.P, open.P))
}

func (cD *ClipperClipperD) Clipper_clipperd_execute_tree(ct ClipperClipType, fr ClipperFillRule, tree ClipperPolyTreeD) int64 {

	return int64(C.clipper_clipperd_execute_tree(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P))
}

func (cD *ClipperClipperD) Clipper_clipperd_execute_tree_with_open(ct ClipperClipType, fr ClipperFillRule, tree ClipperPolyTreeD, open ClipperPathsD) int64 {

	return int64(C.clipper_clipperd_execute_tree_with_open(cD.P, C.ClipperClipType(ct), C.ClipperFillRule(fr), tree.P, open.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_set_miter_limit(l float64) {

	C.clipper_clipperoffset_set_miter_limit(c.P, C.double(l))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_set_arc_tolerance(t float64) {

	C.clipper_clipperoffset_set_arc_tolerance(c.P, C.double(t))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_set_preserve_collinear(t int64) {

	C.clipper_clipperoffset_set_preserve_collinear(c.P, C.int(t))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_set_reverse_solution(t int64) {

	C.clipper_clipperoffset_set_reverse_solution(c.P, C.int(t))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_get_miter_limit() float64 {

	return float64(C.clipper_clipperoffset_get_miter_limit(c.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_get_arc_tolerance() float64 {

	return float64(C.clipper_clipperoffset_get_arc_tolerance(c.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_get_preserve_collinear() int64 {

	return int64(C.clipper_clipperoffset_get_preserve_collinear(c.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_get_reverse_solution() int64 {

	return int64(C.clipper_clipperoffset_get_reverse_solution(c.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_error_code() int64 {

	return int64(C.clipper_clipperoffset_error_code(c.P))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_clear() {

	C.clipper_clipperoffset_clear(c.P)
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_add_pathd(p ClipperPathD, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_pathd(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_add_pathsd(p ClipperPathsD, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_pathsd(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_add_path64(p ClipperPath64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_path64(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_add_paths64(p ClipperPaths64, jt ClipperJoinType, et ClipperEndType) {

	C.clipper_clipperoffset_add_paths64(c.P, p.P, C.ClipperJoinType(jt), C.ClipperEndType(et))
}

func (c *ClipperClipperOffset) Clipper_clipperoffset_execute(delta float64) *ClipperPaths64 {
	var mem unsafe.Pointer = C.malloc(0)

	return &ClipperPaths64{
		P: C.clipper_clipperoffset_execute(mem, c.P, C.double(delta)),
	}

}
