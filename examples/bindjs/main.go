package main

import (
	"github.com/lxn/win"
	"github.com/x51ming/gomb"
)

func main() {
	gmb.Run("../node.dll")
	w := gmb.CreateWebWindow(0, 0, 0, 0, 800, 600)
	w.ShowWindow(true)
	w.MyBindClosing()
	w.MyBindNavi()
	w.MyBindReady()
	w.MyDisablePic()
	gmb.WkeJsBindFunction("f1",
		func(es gmb.JsExecState, p uintptr) uintptr {
			s := ""
			gmb.StartCallBack()
			defer gmb.EndCallBack()
			gmb.MyGetJsArgString(es, 0, &s)
			win.MessageBox(win.HWND(w.WkeGetWindowHandle()),
				win.StringToBSTR(s), win.StringToBSTR("你输入了"),
				win.MB_OK)
			return 0
		}, 0, 1)
	w.WkeLoadFile("test.html")
	<-w.MyBindDestroy()
	gmb.Finalize()
}
