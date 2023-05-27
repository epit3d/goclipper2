#include "callback.h"

// define delta callback
double delta_callback(
    uintptr_t h,
    ClipperPath64 *path, ClipperPathD *path_normals,
    size_t curr_idx, size_t prev_idx)
{
    return goDeltaCallback64(h, path, path_normals, curr_idx, prev_idx);
}
