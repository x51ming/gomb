package gmb

import (
	"syscall"
)

type WkeWebView uintptr
type WebView struct {
	hWkeWebView WkeWebView
}

//
func (mb *WebView) GetHandle() WkeWebView { return mb.hWkeWebView }

//
func (mb *WebView) BindWindowClosing(f WkeWindowClosingCallback, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeOnWindowClosing
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = syscall.NewCallbackCDecl(f)
	api.arg3 = param
	api.nargs = 3
	callDLLAPI(api)
}

//创建WebView
func CreateWebView() WebView {
	var api dllAPI
	api.nargs = 0
	api.hFun = dllWkeCreateWebView
	ret := callDLLAPI(api)
	return WebView{hWkeWebView: WkeWebView(ret.r1)}
}

//LoadURL
func (mb *WebView) LoadURL(url string) {
	var api dllAPI
	api.hFun = dllWkeLoadURL
	api.nargs = 2
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(url)
	callDLLAPI(api)
}

//RunJs
func (mb *WebView) RunJS(js string) JsValue {
	var api dllAPI
	api.hFun = dllWkeRunJS
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(js)
	api.nargs = 2
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}

//wkeGlobalExec
func (mb *WebView) GetExecState() JsExecState {
	var api dllAPI
	api.hFun = dllWkeGlobalExec
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return JsExecState(ret.r1)
}

//GetSource
func (mb *WebView) GetSource() string {
	var api dllAPI
	api.hFun = dllWkeGetSource
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}

//BindDocumentReady2
func (mb *WebView) BindDocumentReady2(f WkeDocumentReady2Callback, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeOnDocumentReady2
	api.arg1 = syscall.NewCallbackCDecl(f)
	api.arg2 = param
	api.nargs = 2
	callDLLAPI(api)
}

//DOMReady
func (mb *WebView) DOMReady() bool {
	var api dllAPI
	api.hFun = dllWkeIsDocumentReady
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}

// 0 点击a标签
// 1 提交表单
// 2 后退前进
// 3 Reload
// 4 重新提交表单
// 5 其他
// @return 1 继续
// @return 0 阻止
func (mb *WebView) BindNavigation(f WkeNavigationCallback, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeOnNavigation
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = syscall.NewCallbackCDecl(f)
	api.arg3 = param
	api.nargs = 3
	callDLLAPI(api)
}

//SetHeadlessEnabled
func (mb *WebView) SetHeadlessEnabled(b bool) {
	var api dllAPI
	api.hFun = dllWkeSetHeadlessEnabled
	api.arg1 = uintptr(mb.hWkeWebView)
	if b {
		api.arg2 = 1
	} else {
		api.arg2 = 0
	}
	api.nargs = 2
	callDLLAPI(api)
}

//Reload
func (mb *WebView) Reload() {
	var api dllAPI
	api.hFun = dllWkeReload
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
}

//调用前要确保有WebWindow
func (mb *WebView) ShowWindow(show bool) {
	var api dllAPI
	api.nargs = 3
	api.hFun = dllWkeShowWindow
	api.arg1 = uintptr(mb.hWkeWebView)
	if show {
		api.arg2 = 1
	} else {
		api.arg2 = 0
	}
	callDLLAPI(api)
}

func (mb *WebView) DestroyWebView() {
	var api dllAPI
	api.hFun = dllWkeDestroyWebView
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
}
