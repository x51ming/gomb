package gmb

import (
	"bytes"
	"log"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

type ptr unsafe.Pointer

var (
	peekMessage        = win.PeekMessage
	translateMessage   = win.TranslateMessage
	dispatchMessage    = win.DispatchMessage
	postQuitMessage    = win.PostQuitMessage
	getCurrentThreadId = win.GetCurrentThreadId
)

func NewWebView(h WkeWebView) WebView {
	var w WebView
	w.hWkeWebView = h
	return w
}
func bool2ptr(b bool) uintptr {
	if b {
		return 1
	} else {
		return 0
	}
}
func ptr2str(p uintptr) string {
	buf := bytes.NewBufferString("")
	for p1 := (*byte)(ptr(p)); *p1 != byte(0); p1 = (*byte)(ptr(uintptr(ptr(p1)) + 1)) {
		buf.WriteByte(*p1)
	}
	return buf.String()
}

//var Ptr2str = ptr2str

func Byteptr2str(b *byte) string { return ptr2str(uintptr(ptr(b))) }

func ptr2bool(p uintptr) bool {
	if p&0xff == 0 {
		return false
	} else {
		return true
	}
}
func str2ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}
func str2wcharptr(s string) uintptr {
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}
func ToolLoadLibrary(lib string) syscall.Handle {
	ret, err := syscall.LoadLibrary(lib)
	if err != nil {
		log.Println("[ERR]LoadLibrary", lib, "", err)
		return syscall.Handle(0)
	} else {
		return ret
	}
}
func ToolLoadFunc(dll syscall.Handle, funcname string) uintptr {
	ret, err := syscall.GetProcAddress(dll, funcname)
	if err != nil {
		log.Println("LoadDllFunction:", funcname, "not loaded.")
		return uintptr(0)
	} else {
		return ret
	}
}
func ret2uint64(r dllReturn) uint64 {
	return uint64(r.r1) + (uint64(r.r2) << 32)
}
func ret2JsValue(r dllReturn) JsValue {
	if Is32 {
		return JsValue(ret2uint64(r))
	} else {
		return JsValue(r.r1)
	}
}
