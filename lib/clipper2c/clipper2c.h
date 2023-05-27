#include "types.h"
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

// Boolean Operations

ClipperPaths64 *clipper_paths64_boolean_op(void *mem, ClipperClipType cliptype,
                                           ClipperFillRule fillrule,
                                           ClipperPaths64 *subjects,
                                           ClipperPaths64 *clips);
void clipper_paths64_boolean_op_tree(ClipperClipType cliptype,
                                     ClipperFillRule fillrule,
                                     ClipperPaths64 *subjects,
                                     ClipperPaths64 *clips,
                                     ClipperPolyTree64 *solution);
ClipperPaths64 *clipper_paths64_intersect(void *mem, ClipperPaths64 *subjects,
                                          ClipperPaths64 *clips,
                                          ClipperFillRule fillrule);
ClipperPaths64 *clipper_paths64_union(void *mem, ClipperPaths64 *subjects,
                                      ClipperPaths64 *clips,
                                      ClipperFillRule fillrule);
ClipperPaths64 *clipper_paths64_difference(void *mem, ClipperPaths64 *subjects,
                                           ClipperPaths64 *clips,
                                           ClipperFillRule fillrule);
ClipperPaths64 *clipper_paths64_xor(void *mem, ClipperPaths64 *subjects,
                                    ClipperPaths64 *clips,
                                    ClipperFillRule fillrule);
ClipperPathsD *clipper_pathsd_boolean_op(void *mem, ClipperClipType cliptype,
                                         ClipperFillRule fillrule,
                                         ClipperPathsD *subjects,
                                         ClipperPathsD *clips,
                                         int decimal_prec);
void clipper_pathsd_boolean_op_tree(
        ClipperClipType cliptype, ClipperFillRule fillrule, ClipperPathsD *subjects,
        ClipperPathsD *clips, ClipperPolyTreeD *solution, int decimal_prec);
ClipperPathsD *clipper_pathsd_intersect(void *mem, ClipperPathsD *subjects,
                                        ClipperPathsD *clips,
                                        ClipperFillRule fillrule,
                                        int decimal_prec);
ClipperPathsD *clipper_pathsd_union(void *mem, ClipperPathsD *subjects,
                                    ClipperPathsD *clips,
                                    ClipperFillRule fillrule, int decimal_prec);
ClipperPathsD *clipper_pathsd_difference(void *mem, ClipperPathsD *subjects,
                                         ClipperPathsD *clips,
                                         ClipperFillRule fillrule,
                                         int decimal_prec);
ClipperPathsD *clipper_pathsd_xor(void *mem, ClipperPathsD *subjects,
                                  ClipperPathsD *clips,
                                  ClipperFillRule fillrule, int decimal_prec);

// Path Offsetting

ClipperPaths64 *clipper_paths64_inflate(void *mem, ClipperPaths64 *paths,
                                        double delta, ClipperJoinType jt,
                                        ClipperEndType et, double miter_limit);
ClipperPathsD *clipper_pathsd_inflate(void *mem, ClipperPathsD *paths,
                                      double delta, ClipperJoinType jt,
                                      ClipperEndType et, double miter_limit,
                                      int precision);

// Rect Clipping
ClipperRect64 *clipper_paths64_bounds(void *mem, ClipperPaths64 *paths);
ClipperPaths64 *clipper_paths64_rect_clip(void *mem, ClipperRect64 *rect,
                                          ClipperPaths64 *paths);
ClipperPaths64 *clipper_paths64_rect_clip_lines(void *mem, ClipperRect64 *rect,
                                                ClipperPaths64 *paths);

// Path Constructors
ClipperPath64 *clipper_path64(void *mem);
ClipperPathD *clipper_pathd(void *mem);
ClipperPath64 *clipper_path64_of_points(void *mem, ClipperPoint64 *pts,
                                        size_t len_pts);
ClipperPathD *clipper_pathd_of_points(void *mem, ClipperPointD *pts,
                                      size_t len_pts);
void clipper_path64_add_point(ClipperPath64 *path, ClipperPoint64 pt);
void clipper_pathd_add_point(ClipperPathD *path, ClipperPointD pt);
ClipperPath64 *clipper_path64_ellipse(void *mem, ClipperPoint64 center,
                                      double radius_x, double radius_y,
                                      int steps);
ClipperPathD *clipper_pathd_ellipse(void *mem, ClipperPointD center,
                                    double radius_x, double radius_y,
                                    int steps);

ClipperPaths64 *clipper_paths64(void *mem);
ClipperPathsD *clipper_pathsd(void *mem);
ClipperPaths64 *clipper_paths64_of_paths(void *mem, ClipperPath64 **paths,
                                         size_t len_paths);
ClipperPathsD *clipper_pathsd_of_paths(void *mem, ClipperPathD **paths,
                                       size_t len_paths);
void clipper_paths64_add_path(ClipperPaths64 *paths, ClipperPath64 *p);
void clipper_pathsd_add_path(ClipperPathsD *paths, ClipperPathD *p);

// Path Conversions (to C)

size_t clipper_path64_length(ClipperPath64 *path);
size_t clipper_pathd_length(ClipperPathD *path);
ClipperPoint64 clipper_path64_get_point(ClipperPath64 *path, int idx);
ClipperPointD clipper_pathd_get_point(ClipperPathD *path, int idx);
ClipperPoint64 *clipper_path64_to_points(void *mem, ClipperPath64 *path);
ClipperPointD *clipper_pathd_to_points(void *mem, ClipperPathD *path);
size_t clipper_paths64_length(ClipperPaths64 *paths);
size_t clipper_pathsd_length(ClipperPathsD *paths);
size_t *clipper_paths64_lengths(void *mem, ClipperPaths64 *paths);
size_t *clipper_pathsd_lengths(void *mem, ClipperPathsD *paths);
size_t clipper_paths64_path_length(ClipperPaths64 *paths, int idx);
size_t clipper_pathsd_path_length(ClipperPathsD *paths, int idx);
ClipperPath64 *clipper_paths64_get_path(void *mem, ClipperPaths64 *paths,
                                        int idx);
ClipperPathD *clipper_pathsd_get_path(void *mem, ClipperPathsD *paths, int idx);
ClipperPoint64 clipper_paths64_get_point(ClipperPaths64 *paths, int path_idx,
                                         int point_idx);
ClipperPointD clipper_pathsd_get_point(ClipperPathsD *paths, int path_idx,
                                       int point_idx);
ClipperPoint64 **clipper_paths64_to_points(void **mem, ClipperPaths64 *paths);
ClipperPointD **clipper_pathsd_to_points(void **mem, ClipperPathsD *paths);

// Path Transformations

ClipperPath64 *clipper_path64_translate(void *mem, ClipperPath64 *path,
                                        int64_t dx, int64_t dy);
ClipperPathD *clipper_pathd_translate(void *mem, ClipperPathD *path, double dx,
                                      double dy);
ClipperPaths64 *clipper_paths64_translate(void *mem, ClipperPaths64 *paths,
                                          int64_t dx, int64_t dy);
ClipperPathsD *clipper_pathsd_translate(void *mem, ClipperPathsD *paths,
                                        double dx, double dy);
ClipperPath64 *clipper_path64_scale(void *mem, ClipperPath64 *path, double sx,
                                    double sy, int *error_code);
ClipperPathD *clipper_pathd_scale(void *mem, ClipperPathD *path, double sx,
                                  double sy, int *error_code);
ClipperPaths64 *clipper_paths64_scale(void *mem, ClipperPaths64 *paths,
                                      double sx, double sy, int *error_code);
ClipperPathsD *clipper_pathsd_scale(void *mem, ClipperPathsD *paths, double sx,
                                    double sy, int *error_code);
ClipperPath64 *clipper_path64_trim_collinear(void *mem, ClipperPath64 *path,
                                             int is_open_path);
ClipperPathD *clipper_pathd_trim_collinear(void *mem, ClipperPathD *path,
                                           int precision, int is_open_path);
ClipperPath64 *clipper_path64_simplify(void *mem, ClipperPath64 *path,
                                       double epsilon, int is_open_path);
ClipperPathD *clipper_pathd_simplify(void *mem, ClipperPathD *path,
                                     double epsilon, int is_open_path);
ClipperPaths64 *clipper_paths64_simplify(void *mem, ClipperPaths64 *paths,
                                         double epsilon, int is_open_paths);
ClipperPathsD *clipper_pathsd_simplify(void *mem, ClipperPathsD *paths,
                                       double epsilon, int is_open_paths);
ClipperPath64 *clipper_path64_ramer_douglas_peucker(void *mem,
                                                    ClipperPath64 *path,
                                                    double epsilon);
ClipperPathD *clipper_pathd_ramer_douglas_peucker(void *mem, ClipperPathD *path,
                                                  double epsilon);
ClipperPaths64 *clipper_paths64_ramer_douglas_peucker(void *mem,
                                                      ClipperPaths64 *paths,
                                                      double epsilon);
ClipperPathsD *clipper_pathsd_ramer_douglas_peucker(void *mem,
                                                    ClipperPathsD *paths,
                                                    double epsilon);
ClipperPath64 *clipper_path64_strip_near_equal(void *mem, ClipperPath64 *path,
                                               double max_dist_sqrd,
                                               int is_closed_path);
ClipperPathD *clipper_pathd_strip_near_equal(void *mem, ClipperPathD *path,
                                             double max_dist_sqrd,
                                             int is_closed_path);
ClipperPaths64 *clipper_paths64_strip_near_equal(void *mem,
                                                 ClipperPaths64 *paths,
                                                 double max_dist_sqrd,
                                                 int is_closed_paths);
ClipperPathsD *clipper_pathsd_strip_near_equal(void *mem, ClipperPathsD *paths,
                                               double max_dist_sqrd,
                                               int is_closed_paths);
ClipperPath64 *clipper_path64_strip_duplicates(void *mem, ClipperPath64 *path,
                                               int is_closed_path);
ClipperPathD *clipper_pathd_strip_duplicates(void *mem, ClipperPathD *path,
                                             int is_closed_path);
ClipperPaths64 *clipper_paths64_strip_duplicates(void *mem,
                                                 ClipperPaths64 *paths,
                                                 int is_closed_paths);
ClipperPathsD *clipper_pathsd_strip_duplicates(void *mem, ClipperPathsD *paths,
                                               int is_closed_paths);

// Path Conversions

ClipperPathD *clipper_path64_to_pathd(void *mem, ClipperPath64 *path);
ClipperPath64 *clipper_pathd_to_path64(void *mem, ClipperPathD *path);
ClipperPathsD *clipper_paths64_to_pathsd(void *mem, ClipperPaths64 *paths);
ClipperPaths64 *clipper_pathsd_to_paths64(void *mem, ClipperPathsD *paths);
ClipperPathD *clipper_scale_path64_to_pathd(void *mem, ClipperPath64 *path,
                                            double sx, double sy,
                                            int *error_code);
ClipperPath64 *clipper_scale_pathd_to_path64(void *mem, ClipperPathD *path,
                                             double sx, double sy,
                                             int *error_code);
ClipperPathsD *clipper_scale_paths64_to_pathsd(void *mem, ClipperPaths64 *paths,
                                               double sx, double sy,
                                               int *error_code);
ClipperPaths64 *clipper_scale_pathsd_to_paths64(void *mem, ClipperPathsD *paths,
                                                double sx, double sy,
                                                int *error_code);

// Minkowski

ClipperPaths64 *clipper_path64_minkowski_sum(void *mem, ClipperPath64 *pattern,
                                             ClipperPath64 *path,
                                             int is_closed);
ClipperPathsD *clipper_pathd_minkowski_sum(void *mem, ClipperPathD *pattern,
                                           ClipperPathD *path, int is_closed,
                                           int precision);
ClipperPaths64 *clipper_path64_minkowski_diff(void *mem, ClipperPath64 *pattern,
                                              ClipperPath64 *path,
                                              int is_closed);
ClipperPathsD *clipper_pathd_minkowski_diff(void *mem, ClipperPathD *pattern,
                                            ClipperPathD *path, int is_closed,
                                            int precision);

// Geometry

double clipper_point64_distance(ClipperPoint64 a, ClipperPoint64 b);
double clipper_pointd_distance(ClipperPointD a, ClipperPointD b);
int clipper_point64_near_collinear(ClipperPoint64 a, ClipperPoint64 b,
                                   ClipperPoint64 c,
                                   double sin_sqrd_min_angle_rads);
int clipper_pointd_near_collinear(ClipperPointD a, ClipperPointD b,
                                  ClipperPointD c,
                                  double sin_sqrd_min_angle_rads);
double clipper_pathd_area(ClipperPathD *path);
double clipper_path64_area(ClipperPath64 *path);
double clipper_pathsd_area(ClipperPathsD *paths);
double clipper_paths64_area(ClipperPaths64 *paths);
int clipper_pathd_is_positive(ClipperPathD *path);
int clipper_path64_is_positive(ClipperPath64 *path);
ClipperPointInPolygonResult clipper_point_in_path64(ClipperPath64 *path,
                                                    ClipperPoint64 pt);
ClipperPointInPolygonResult clipper_point_in_pathd(ClipperPathD *path,
                                                   ClipperPointD pt);

// Class Interfaces

// PolyTree Constructors

ClipperPolyTree64 *clipper_polytree64(void *mem, ClipperPolyTree64 *parent);
ClipperPolyTreeD *clipper_polytreed(void *mem, ClipperPolyTreeD *parent);

// PolyTree64 Methods

const ClipperPolyTree64 *clipper_polytree64_parent(ClipperPolyTree64 *pt);
const ClipperPolyTree64 *clipper_polytree64_get_child(ClipperPolyTree64 *pt,
                                                      size_t idx);
ClipperPolyTree64 *clipper_polytree64_add_child(ClipperPolyTree64 *pt,
                                                ClipperPath64 *path);
void clipper_polytree64_clear(ClipperPolyTree64 *pt);
size_t clipper_polytree64_count(ClipperPolyTree64 *pt);
int clipper_polytree64_level(ClipperPolyTree64 *pt);
int clipper_polytree64_is_hole(ClipperPolyTree64 *pt);
ClipperPath64 *clipper_polytree64_polygon(void *mem, ClipperPolyTree64 *pt);
double clipper_polytree64_area(ClipperPolyTree64 *pt);
ClipperPaths64 *clipper_polytree64_to_paths(void *mem, ClipperPolyTree64 *pt);
int clipper_polytree64_fully_contains_children(ClipperPolyTree64 *pt);

// PolyTreeD Methods

const ClipperPolyTreeD *clipper_polytreed_parent(ClipperPolyTreeD *pt);
const ClipperPolyTreeD *clipper_polytreed_get_child(ClipperPolyTreeD *pt,
                                                    size_t idx);
void clipper_polytreed_set_inv_scale(ClipperPolyTreeD *pt, double value);
double clipper_polytreed_inv_scale(ClipperPolyTreeD *pt);
ClipperPolyTreeD *clipper_polytreed_add_child(ClipperPolyTreeD *pt,
                                              ClipperPath64 *path);
void clipper_polytreed_clear(ClipperPolyTreeD *pt);
size_t clipper_polytreed_count(ClipperPolyTreeD *pt);
int clipper_polytreed_level(ClipperPolyTreeD *pt);
int clipper_polytreed_is_hole(ClipperPolyTreeD *pt);
ClipperPathD *clipper_polytreed_polygon(void *mem, ClipperPolyTreeD *pt);
double clipper_polytreed_area(ClipperPolyTreeD *pt);
ClipperPathsD *clipper_polytreed_to_paths(void *mem, ClipperPolyTreeD *pt);

// Rect Constructors

ClipperRect64 *clipper_rect64(void *mem, int64_t left, int64_t top,
                              int64_t right, int64_t bottom);

// Rect64 Methods

int64_t clipper_rect64_width(ClipperRect64 *r);
int64_t clipper_rect64_height(ClipperRect64 *r);
ClipperPoint64 clipper_rect64_midpoint(ClipperRect64 *r);
ClipperPath64 *clipper_rect64_as_path(void *mem, ClipperRect64 *r);
int clipper_rect64_contains_pt(ClipperRect64 *r, ClipperPoint64 pt);
int clipper_rect64_contains_rect(ClipperRect64 *a, ClipperRect64 *b);
void clipper_rect64_scale_mut(ClipperRect64 *r, double scale);
ClipperRect64 *clipper_rect64_scale(void *mem, ClipperRect64 *r, double scale);
int clipper_rect64_is_empty(ClipperRect64 *r);
int clipper_rect64_intersects(ClipperRect64 *a, ClipperRect64 *b);

// Rect Conversions (to C)

struct ClipperRect64 clipper_rect64_to_struct(ClipperRect64 *rect);

// Clipper Contsructors

ClipperClipper64 *clipper_clipper64(void *mem);
ClipperClipperD *clipper_clipperd(void *mem, int precision);

// Clipper64 Setters / Getters

void clipper_clipper64_set_preserve_collinear(ClipperClipper64 *c, int t);
void clipper_clipper64_set_reverse_solution(ClipperClipper64 *c, int t);
int clipper_clipper64_get_preserve_collinear(ClipperClipper64 *c);
int clipper_clipper64_get_reverse_solution(ClipperClipper64 *c);
void clipper_clipper64_clear(ClipperClipper64 *c);

// ClipperD Setters / Getters
//
void clipper_clipperd_set_preserve_collinear(ClipperClipperD *c, int t);
void clipper_clipperd_set_reverse_solution(ClipperClipperD *c, int t);
int clipper_clipperd_get_preserve_collinear(ClipperClipperD *c);
int clipper_clipperd_get_reverse_solution(ClipperClipperD *c);
void clipper_clipperd_clear(ClipperClipperD *c);

// Clipper64 Methods

void clipper_clipper64_add_subject(ClipperClipper64 *c,
                                   ClipperPaths64 *subjects);
void clipper_clipper64_add_open_subject(ClipperClipper64 *c,
                                        ClipperPaths64 *open_subjects);
void clipper_clipper64_add_clip(ClipperClipper64 *c, ClipperPaths64 *clips);
int clipper_clipper64_execute(ClipperClipper64 *c64, ClipperClipType ct,
                              ClipperFillRule fr, ClipperPaths64 *closed,
                              ClipperPaths64 *open);
int clipper_clipper64_execute_tree(ClipperClipper64 *c64, ClipperClipType ct,
                                   ClipperFillRule fr, ClipperPolyTree64 *tree);
int clipper_clipper64_execute_tree_with_open(ClipperClipper64 *c64,
                                             ClipperClipType ct,
                                             ClipperFillRule fr,
                                             ClipperPolyTree64 *tree,
                                             ClipperPaths64 *open);

// ClipperD Methods

void clipper_clipperd_add_subject(ClipperClipperD *c, ClipperPathsD *subjects);
void clipper_clipperd_add_open_subject(ClipperClipperD *c,
                                       ClipperPathsD *open_subjects);
void clipper_clipperd_add_clip(ClipperClipperD *c, ClipperPathsD *clips);
int clipper_clipperd_execute(ClipperClipperD *cD, ClipperClipType ct,
                             ClipperFillRule fr, ClipperPathsD *closed,
                             ClipperPathsD *open);
int clipper_clipperd_execute_tree(ClipperClipperD *cD, ClipperClipType ct,
                                  ClipperFillRule fr, ClipperPolyTreeD *tree);
int clipper_clipperd_execute_tree_with_open(ClipperClipperD *cD,
                                            ClipperClipType ct,
                                            ClipperFillRule fr,
                                            ClipperPolyTreeD *tree,
                                            ClipperPathsD *open);

// ClipperOffset Constructors

ClipperClipperOffset *clipper_clipperoffset(void *mem, double miter_limit,
                                            double arc_tolerance,
                                            int preserve_collinear,
                                            int reverse_solution);

// ClipperOffset Setters / Getters

void clipper_clipperoffset_set_miter_limit(ClipperClipperOffset *c, double l);
void clipper_clipperoffset_set_arc_tolerance(ClipperClipperOffset *c, double t);
void clipper_clipperoffset_set_preserve_collinear(ClipperClipperOffset *c,
                                                  int t);
void clipper_clipperoffset_set_reverse_solution(ClipperClipperOffset *c, int t);
double clipper_clipperoffset_get_miter_limit(ClipperClipperOffset *c);
double clipper_clipperoffset_get_arc_tolerance(ClipperClipperOffset *c);
int clipper_clipperoffset_get_preserve_collinear(ClipperClipperOffset *c);
int clipper_clipperoffset_get_reverse_solution(ClipperClipperOffset *c);
int clipper_clipperoffset_error_code(ClipperClipperOffset *c);
void clipper_clipperoffset_clear(ClipperClipperOffset *c);

// ClipperOffset Methods
void clipper_clipperoffset_add_path64(ClipperClipperOffset *c, ClipperPath64 *p,
                                      ClipperJoinType jt, ClipperEndType et);
void clipper_clipperoffset_add_paths64(ClipperClipperOffset *c,
                                       ClipperPaths64 *p, ClipperJoinType jt,
                                       ClipperEndType et);
ClipperPaths64 *
clipper_clipperoffset_execute(void *mem, ClipperClipperOffset *c, double delta);

// execute with callback
ClipperPaths64 *
clipper_clipperoffset_execute_callback(void *mem, ClipperClipperOffset *c,
                                       ClipperDeltaCallback64 cb);
#ifdef GO_BINDINGS
ClipperPaths64 *
clipper_clipperoffset_execute_gocallback(void *mem, ClipperClipperOffset *c,
                                         ClipperDeltaGoCallback64 cb);
#endif
// memory size

size_t clipper_path64_size();
size_t clipper_pathd_size();
size_t clipper_paths64_size();
size_t clipper_pathsd_size();
size_t clipper_rect64_size();
size_t clipper_rectd_size();
size_t clipper_polytree64_size();
size_t clipper_polytreed_size();
size_t clipper_clipper64_size();
size_t clipper_clipperd_size();
size_t clipper_clipperoffset_size();

// destruction

void clipper_destruct_path64(ClipperPath64 *p);
void clipper_destruct_pathd(ClipperPathD *p);
void clipper_destruct_paths64(ClipperPaths64 *p);
void clipper_destruct_pathsd(ClipperPathsD *p);
void clipper_destruct_rect64(ClipperRect64 *p);
void clipper_destruct_rectd(ClipperRectD *p);
void clipper_destruct_polytree64(ClipperPolyTree64 *p);
void clipper_destruct_polytreed(ClipperPolyTreeD *p);
void clipper_destruct_clipper64(ClipperClipper64 *p);
void clipper_destruct_clipperd(ClipperClipperD *p);
void clipper_destruct_clipperoffset(ClipperClipperOffset *p);

// pointer free + destruction

void clipper_delete_path64(ClipperPath64 *p);
void clipper_delete_pathd(ClipperPathD *p);
void clipper_delete_paths64(ClipperPaths64 *p);
void clipper_delete_pathsd(ClipperPathsD *p);
void clipper_delete_rect64(ClipperRect64 *p);
void clipper_delete_rectd(ClipperRectD *p);
void clipper_delete_polytree64(ClipperPolyTree64 *p);
void clipper_delete_polytreed(ClipperPolyTreeD *p);
void clipper_delete_clipper64(ClipperClipper64 *p);
void clipper_delete_clipperd(ClipperClipperD *p);
void clipper_delete_clipperoffset(ClipperClipperOffset *p);

#ifdef __cplusplus
}
#endif
