#include "_cgo_export.h"
void gozoo_watcher(zhandle_t *zzh, int type, int state, const char *path, void *watcherCtx) {
   goCallback(type, state, path, watcherCtx);
}
char* get_string(char** strings, int index){
   return strings[index];
}

