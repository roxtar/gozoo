#include "_cgo_export.h"
void gozoo_watcher(zhandle_t *zzh, int type, int state, const char *path, void *watcherCtx) {
   goCallback(type, state, path, watcherCtx);
}
