#ifndef CLIBRARY_H
#define CLIBRARY_H

#include <stddef.h>
#include "../lib/clipper2c/clipper2c.h"

typedef double (*callbackfunc) (
    uintptr_t h, ClipperPath64 *path, ClipperPathD *path_normals,
                    size_t curr_idx, size_t prev_idx);

double delta_callback(
    uintptr_t h,
    ClipperPath64 *path, ClipperPathD *path_normals,
    size_t curr_idx, size_t prev_idx);

extern double goDeltaCallback64(uintptr_t h, ClipperPath64 *path,
                                ClipperPathD *path_normals, size_t curr_idx,
                                size_t prev_idx);

#endif