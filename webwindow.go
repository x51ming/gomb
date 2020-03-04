package gmb

var Is32 = true

//c int -> int32
type HWND uintptr

func CreateWebWindow(wtype int32, parent HWND, x, y, w, h int32) WebView {
	var api dllAPI
	api.nargs = 6
	api.hFun = dllWkeCreateWebWindow
	api.arg1 = uintptr(wtype)
	api.arg2 = uintptr(parent)
	api.arg3 = uintptr(x)
	api.arg4 = uintptr(y)
	api.arg5 = uintptr(w)
	api.arg6 = uintptr(h)
	ret := callDLLAPI(api)
	return WebView{hWkeWebView: WkeWebView(ret.r1)}
}
