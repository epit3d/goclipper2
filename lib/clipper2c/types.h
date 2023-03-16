#pragma once
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct ClipperClipper64 ClipperClipper64;
typedef struct ClipperClipperD ClipperClipperD;
typedef struct ClipperClipperOffset ClipperClipperOffset;
typedef struct ClipperPath64 ClipperPath64;
typedef struct ClipperPathD ClipperPathD;
typedef struct ClipperPaths64 ClipperPaths64;
typedef struct ClipperPathsD ClipperPathsD;
typedef struct ClipperRect64 ClipperRect64;
typedef struct ClipperRectD ClipperRectD;
typedef struct ClipperPolyTree64 ClipperPolyTree64;
typedef struct ClipperPolyTreeD ClipperPolyTreeD;

typedef struct ClipperPointD {
  double x;
  double y;
} ClipperPointD;

typedef struct ClipperPoint64 {
  int64_t x;
  int64_t y;
} ClipperPoint64;

struct ClipperRect64 {
  int64_t left;
  int64_t top;
  int64_t right;
  int64_t bottom;
};

struct ClipperRectD {
  double left;
  double top;
  double right;
  double bottom;
};

typedef enum ClipperFillRule {
  EVEN_ODD,
  NON_ZERO,
  POSITIVE,
  NEGATIVE
} ClipperFillRule;

typedef enum ClipperClipType {
  NONE,
  INTERSECTION,
  UNION,
  DIFFERENCE,
  XOR
} ClipperClipType;

typedef enum ClipperPathType { SUBJECT, CLIP } ClipperPathType;

typedef enum ClipperJoinType {
  SQUARE_JOIN,
  ROUND_JOIN,
  MITER_JOIN
} ClipperJoinType;

typedef enum ClipperEndType {
  POLYGON_END,
  JOINED_END,
  BUTT_END,
  SQUARE_END,
  ROUND_END
} ClipperEndType;

typedef enum ClipperPointInPolygonResult {
  IS_ON,
  IS_INSIDE,
  IS_OUTSIDE
} ClipperPointInPolygonResult;

#ifdef __cplusplus
}
#endif
