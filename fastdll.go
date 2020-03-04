package gmb

import (
	"syscall"
)

func (mb *WebView) f_void(fun uintptr) uintptr {
	var api dllAPI
	api.hFun = fun
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	return callDLLAPI(api).r1
}
func (mb *WebView) f_bool_void(fun uintptr, b bool) {
	var api dllAPI
	api.hFun = fun
	api.arg1 = uintptr(mb.hWkeWebView)
	if b {
		api.arg2 = 1
	} else {
		api.arg2 = 0
	}
	api.nargs = 2
	callDLLAPI(api)
}
func f_js_to(fun uintptr, es JsExecState, v JsValue) uintptr {
	var api dllAPI
	api.hFun = fun
	api.arg1 = uintptr(es)
	if Is32 {
		api.arg2 = uintptr(uint64(v) & 0xffffffff)
		api.arg3 = uintptr(uint64(v) >> 32)
		api.nargs = 3
	} else {
		api.arg2 = uintptr(v)
		api.nargs = 2
	}
	ret := callDLLAPI(api)
	return ret.r1
}
func (mb *WebView) f_cb_p(fun uintptr, cb interface{}, param uintptr) {
	var api dllAPI
	api.hFun = fun
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = syscall.NewCallbackCDecl(cb)
	api.arg3 = param
	api.nargs = 3
	callDLLAPI(api)
}
