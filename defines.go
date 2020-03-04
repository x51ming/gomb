package gmb

import (
	"github.com/lxn/win"
)

type WkeString uintptr
type JsValue int64
type JsExecState uintptr
type WkeNetJob uintptr

//复杂callback
//begin 返回 bool
//end 无返回

//返回1，不请求，返回0，继续请求
type WkeLoadUrlBeginCallback func(w WkeWebView, param uintptr, url *byte, job WkeNetJob) uintptr
type WkeLoadUrlEndCallback func(w WkeWebView, param uintptr, url *byte, job WkeNetJob, buf uintptr, length uint32) uintptr
type WkeLoadUrlFailCallback WkeLoadUrlEndCallback
type WkeDownloadCallback func(w WkeWebView, param uintptr, url *byte) uintptr
type WkeDownloadCallback2 func(
	w WkeWebView,
	param uintptr,
	expectedLength uintptr,
	url *byte,
	mime *byte,
	disposition *byte,
	job WkeNetJob,
	dataBind *WkeNetJobDataBind,
) WkeDownloadOpt
type WkeNetResponseCallback WkeLoadUrlBeginCallback
type WkeOnOtherLoadCallback func(w WkeWebView, param uintptr, otype WkeOtherLoadType, info *WkeTempCallbackInfo) uintptr
type WkeNetJobDataRecvCallback func(_ptr uintptr, job WkeNetJob, data *byte, length int32)
type WkeNetJobDataFinishCallback func(_ptr uintptr, job WkeNetJob, result WkeLoadingResult)

//CBtypes
type CB_view_p_void func(w WkeWebView, param uintptr) uintptr
type CB_view_p_bool CB_view_p_void
type CB_view_p_str_void func(w WkeWebView, param uintptr, str WkeString) uintptr
type CB_view_p_str_bool CB_view_p_str_void
type CB_view_p_fid_str_void func(w WkeWebView, param uintptr, frameid uintptr, str WkeString) uintptr
type CB_view_p_fid_bool func(w WkeWebView, param uintptr, frameid uintptr) uintptr

type WkeCreateViewCallback func(w WkeWebView, param uintptr, navitype uintptr, url WkeString, wf *WkeWindowFeatures) uintptr
type WkeWindowClosingCallback CB_view_p_bool
type WkeWindowDestroyCallback CB_view_p_void

type WkeOnShowDevtoolsCallback CB_view_p_void

type WkeAlertBoxCallback CB_view_p_str_void
type WkeConfirmBoxCallback CB_view_p_str_bool
type WkeNavigationCallback func(w WkeWebView, param uintptr, navigationType uint32, url WkeString) uintptr

type WkeDocumentReadyCallback CB_view_p_void
type WkeDocumentReady2Callback CB_view_p_fid_bool
type WkeTitleChangedCallback CB_view_p_str_void
type WkeTitleChangedCallback2 CB_view_p_fid_str_void
type WkeURLChangedCallback CB_view_p_str_void
type WkeURLChangedCallback2 CB_view_p_fid_str_void

//__fastcall
type JsNativeFunction func(es JsExecState) JsValue

//__cdecl
type WkeJsNativeFunction func(es JsExecState, param uintptr) uintptr

type Proxy struct {
	ptype    uint32
	hostname [100]byte
	port     uint32
	user     [50]byte
	pwd      [50]byte
}
type WkeWindowFeatures struct {
	x      int32
	y      int32
	width  int32
	height int32

	menuBarVisible     bool
	statusBarVisible   bool
	toolBarVisible     bool
	locationBarVisible bool
	scrollbarsVisible  bool
	resizable          bool
	fullscreen         bool
}
type WkeWebFrameHandle uintptr
type WkeWillSendRequestInfo struct {
	Url          WkeString
	NewUrl       WkeString
	ResType      WkeResourceType
	Httprespcode int32
	Method       WkeString
	Referer      WkeString
	Headers      uintptr
}

type WkeMemBuf struct {
	size   int32
	data   uintptr
	length uintptr
}
type WkePostBodyElement struct {
	size      int32
	bodytype  WkeHttBodyElementType
	data      *WkeMemBuf
	filepath  WkeString
	filestart int64
	filelen   int64
}
type WkePostBodyElements struct {
	size        int32
	element     **WkePostBodyElement
	elementsize uintptr
	isdirty     bool
}
type WkeTempCallbackInfo struct {
	Size     int32
	Frame    WkeWebFrameHandle
	Info     *WkeWillSendRequestInfo
	Url      *byte
	PostBody *WkePostBodyElements
	Job      WkeNetJob
}
type WkeNetJobDataBind struct {
	Param    uintptr
	Recvcb   uintptr
	Finishcb uintptr
}

const (
	WKE_LBUTTON = 1
	WKE_RBUTTON = 2
	WKE_SHIFT   = 4
	WKE_CONTROL = 8
	WKE_MBUTTON = 16

	WM_PAINT                = 15
	WM_SIZE                 = 5
	WM_QUIT                 = win.WM_QUIT
	WM_KEYDOWN              = 256
	WM_KEYUP                = 257
	WM_CHAR                 = 258
	WM_LBUTTONDOWN          = 513
	WM_MBUTTONDOWN          = 519
	WM_RBUTTONDOWN          = 516
	WM_LBUTTONDBLCLK        = 515
	WM_MBUTTONDBLCLK        = 521
	WM_RBUTTONDBLCLK        = 518
	WM_LBUTTONUP            = 514
	WM_MBUTTONUP            = 520
	WM_RBUTTONUP            = 517
	WM_MOUSEMOVE            = 512
	MK_CONTROL              = 8
	MK_SHIFT                = 4
	MK_LBUTTON              = 1
	MK_MBUTTON              = 16
	MK_RBUTTON              = 2
	WM_CONTEXTMENU          = 123
	WM_MOUSEWHEEL           = 522
	WM_SETFOCUS             = 7
	WM_KILLFOCUS            = 8
	WM_IME_STARTCOMPOSITION = 269
	CFS_EXCLUDE             = 128
	WS_CHILD                = 1073741824
	WS_VISIBLE              = 268435456
	SS_NOTIFY               = 256
	WS_EX_LAYERED           = 524288
	WM_NCHITTEST            = 132
	WM_GETMINMAXINFO        = 36
	HTTOP                   = 12
	HTLEFT                  = 10
	HTRIGHT                 = 11
	HTBOTTOM                = 15
	HTTOPLEFT               = 13
	HTTOPRIGHT              = 14
	HTBOTTOMLEFT            = 16
	HTBOTTOMRIGHT           = 17
	SPI_GETWORKAREA         = 48
	WM_DESTROY              = 2
	WKE_EXTENDED            = 256
	WKE_REPEAT              = 16384
	GWL_EXSTYLE             = -20
	SRCCOPY                 = 13369376
	WM_TIMER                = 275
	WM_ERASEBKGND           = 20
	WM_SETCURSOR            = 32
	OBJ_BITMAP              = 7
	CAPTUREBLT              = 1073741824
	GA_ROOT                 = 2
	WM_NCDESTROY            = 130
	WM_CLOSE                = 16
	CP_UTF8                 = 65001
	CP_936                  = 936
	CFS_POINT               = 2
	CFS_FORCE_POSITION      = 32
)
const (
	WKE_PROXY_NONE uint32 = iota
	WKE_PROXY_HTTP
	WKE_PROXY_SOCKS4
	WKE_PROXY_SOCKS4A
	WKE_PROXY_SOCKS5
	WKE_PROXY_SOCKS5HOSTNAME
)
const (
	WKE_NAVIGATION_TYPE_LINKCLICK uint32 = iota
	WKE_NAVIGATION_TYPE_FORMSUBMITTE
	WKE_NAVIGATION_TYPE_BACKFORWARD
	WKE_NAVIGATION_TYPE_RELOAD
	WKE_NAVIGATION_TYPE_FORMRESUBMITT
	WKE_NAVIGATION_TYPE_OTHER
)
