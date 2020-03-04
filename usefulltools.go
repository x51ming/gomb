package gmb

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lxn/win"
)

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

//需要位于当前目录
func MyGetDevtoolsPath(s string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	f := filepath.Join(dir, s)
	return "file:///" + strings.ReplaceAll(f, "\\", "/")
}

func Cmpsuf3(s, t string) bool {
	i := strings.IndexByte(s, '?')
	if i == -1 {
		i = len(s)
	}
	return s[i-4] == '.' &&
		s[i-3] == t[0] &&
		s[i-2] == t[1] &&
		s[i-1] == t[2]
}
func (mb *WebView) MyBindClosing() {
	mb.WkeOnWindowClosing(
		func(w WkeWebView, param uintptr) uintptr {
			StartCallBack()
			h := win.HWND(mb.WkeGetWindowHandle())
			EndCallBack()
			if win.MessageBox(h,
				win.StringToBSTR("是否退出？"), win.StringToBSTR("提示"),
				win.MB_OKCANCEL) == win.IDOK {
				return 1
			} else {
				return 0
			}
		}, 0)
}
func (mb *WebView) MyDisablePic() {
	mb.WkeOnLoadUrlBegin(
		func(w WkeWebView, param uintptr, url *byte, job WkeNetJob) uintptr {
			uu := Byteptr2str(url)
			if Cmpsuf3(uu, "jpg") ||
				Cmpsuf3(uu, "png") ||
				Cmpsuf3(uu, "gif") ||
				Cmpsuf3(uu, "bmp") {
				StartCallBack()
				WkeNetCancelRequest(job)
				EndCallBack()
			}
			return 0
		},
		0)
}
func (mb *WebView) MyBindDestroy() chan bool {
	c := make(chan bool, 1)
	mb.WkeOnWindowDestroy(func(w WkeWebView, p uintptr) uintptr {
		log.Println("WebView", w, "Destroy")
		c <- true
		defer close(c)
		return 0
	}, 0)
	return c
}
func (mb *WebView) MyBindNavi() {
	mb.WkeOnNavigation(
		func(w WkeWebView, p uintptr, tp uint32, url WkeString) uintptr {
			log.Println("WebView", w, "Navi", tp)
			return 1
		},
		0)
}
func (mb *WebView) MyBindReady() {
	mb.BindDocumentReady2(func(w WkeWebView, p uintptr, f uintptr) uintptr {
		log.Println("WebView", w, "Frame", f, "Ready.")
		return 0
	}, 0)
}
func (mb *WebView) MyWaitReady(timeout int) chan bool {
	time.Sleep(time.Millisecond * 50)
	c := make(chan bool, 1)
	go func() {
		defer close(c)
		i := 0
		for !mb.DOMReady() {
			if i == timeout*10 {
				c <- false
				return
			}
			i++
			time.Sleep(time.Millisecond * 100)
		}
		c <- true
	}()
	return c
}

//
func MyGetJsArgBool(es JsExecState, i int32, b *bool) bool {
	if JsArgCount(es) > i && JsArgType(es, i) == JSTYPE_BOOLEAN {
		*b = JsToBoolean(es, JsArg(es, i))
		return true
	}
	return false
}

//
func MyGetJsArgString(es JsExecState, i int32, s *string) bool {
	if JsArgCount(es) > i && JsArgType(es, i) == JSTYPE_STRING {
		*s = JsToString(es, JsArg(es, i))
		return true
	}
	return false
}

//第二个返回值ok
func MyJsArg(es JsExecState, i int32) (interface{}, bool) {
	v := JsArg(es, i)
	switch JsArgType(es, i) {
	case JSTYPE_NUMBER:
		return nil, false
	case JSTYPE_STRING:
		return JsToString(es, v), true
	case JSTYPE_BOOLEAN:
		return JsToBoolean(es, v), true
	case JSTYPE_OBJECT:
		return nil, false
	case JSTYPE_FUNCTION:
		return nil, false
	case JSTYPE_UNDEFINED:
		return nil, false
	case JSTYPE_ARRAY:
		return nil, false
	case JSTYPE_NULL:
		return nil, true
	}
	return nil, false
}
