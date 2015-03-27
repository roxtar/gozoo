package gozoo

// #cgo LDFLAGS: -lzookeeper_mt -lm -lpthread
// #include <stdio.h>
// #include <errno.h>
// #include <stdlib.h>
// #include <zookeeper/zookeeper.h>
// typedef const char * const_char_ptr;
// typedef char * char_ptr;
// void gozoo_watcher(zhandle_t *zzh, int type, int state, const char *path, void *watcherCtx);
import "C"
import "unsafe"
import "fmt"

type WatcherCallback func(zooType int, zooState int, path string)

type ZooClient struct {
	handle       *C.zhandle_t
	BufferLength int
	Callback     WatcherCallback
}

func NewClient() ZooClient {
	return ZooClient{BufferLength: 1024}
}

func (z *ZooClient) Init(hostname string, recvTimeout int) error {
	chostname := (C.const_char_ptr)(C.CString(hostname))
	defer C.free(unsafe.Pointer(chostname))
	zhandle, err := C.zookeeper_init(chostname, C.watcher_fn(C.gozoo_watcher), C.int(recvTimeout), nil, unsafe.Pointer(z), 0)
	z.handle = zhandle
	if zhandle == nil {
		return err
	} else {
		return nil
	}
}

func (z *ZooClient) Close() error {
	err := C.zookeeper_close(z.handle)
	if err != 0 {
		return fmt.Errorf("%s", convertZookeeperError(err))
	}
	return nil
}

func (z *ZooClient) Create(path string, value []byte) (string, error) {
	return z.CreateWithFlags(path, value, 0)
}

func (z *ZooClient) CreateWithFlags(path string, value []byte, flags ZookeeperCreateFlag) (string, error) {
	bufferLength := z.BufferLength
	buffer := C.char_ptr(C.malloc(C.size_t(bufferLength)))
	defer C.free(unsafe.Pointer(buffer))
	var valuePtr C.const_char_ptr = nil
	if len(value) > 0 {
		valuePtr = (C.const_char_ptr)(unsafe.Pointer(&value[0]))
	}
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	err := C.zoo_create(z.handle, cpath, valuePtr, C.int(len(value)), &C.ZOO_OPEN_ACL_UNSAFE, C.int(flags), buffer, C.int(bufferLength))
	if err != 0 {
		return "", newZooError(convertZookeeperError(err))
	}
	return C.GoString(buffer), nil
}

func (z *ZooClient) Delete(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	err := C.zoo_delete(z.handle, cpath, -1)
	if err != 0 {
		return newZooError(convertZookeeperError(err))
	}
	return nil
}

func (z *ZooClient) Get(path string) ([]byte, error) {

	bufferLength := z.BufferLength
	buffer := C.char_ptr(C.malloc(C.size_t(bufferLength)))
	defer C.free(unsafe.Pointer(buffer))

	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var actualBufferLength C.int = C.int(bufferLength)

	err := C.zoo_get(z.handle, cpath, 1, buffer, &actualBufferLength, nil)
	if err != 0 {
		return []byte{}, newZooError(convertZookeeperError(err))
	}
	if actualBufferLength > 0 {
		value := C.GoBytes(unsafe.Pointer(buffer), actualBufferLength)
		return value, nil
	}
	return []byte{}, nil
}

func (z *ZooClient) Set(path string, value []byte) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	valuePtr := (C.const_char_ptr)(unsafe.Pointer(&value[0]))

	err := C.zoo_set(z.handle, cpath, valuePtr, C.int(len(value)), -1)
	if err != 0 {
		return fmt.Errorf("%s", convertZookeeperError(err))
	}
	return nil
}

//export goCallback
func goCallback(zooType int, zooState int, path C.const_char_ptr, context unsafe.Pointer) {
	//	gpath := C.GoString(path)
	//	z, ok := context.(*ZooClient)
	//	if ok && z.Callback != nil {
	//		z.Callback(zooType, zooState, gpath)
	//	}
}
