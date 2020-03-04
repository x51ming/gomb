package gmb

import (
	"log"
	"unsafe"
)

//结束并释放node.dll
func Finalize() {
	//postQuitMessage(0)

	var api dllAPI
	api.nargs = 0
	// api.hFun = dllWkeFinalize
	// callDLLAPI(api)

	//结束事件循环
	api.nargs = ExitNarg
	callDLLAPI(api)
}

func Run(dll string) bool {
	loadMBDLL(dll)
	RunLoop()
	return WkeInit()
}

//是否已经初始化
func WkeIsInitialize() bool {
	var api dllAPI
	api.hFun = dllWkeIsInitialize
	api.nargs = 0
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}

//初始化
func WkeInit() bool {
	if !WkeIsInitialize() {
		var api dllAPI
		api.hFun = dllWkeInitialize
		api.nargs = 0
		callDLLAPI(api)
		if WkeIsInitialize() {
			log.Println("MB初始化成功")
			return true
		} else {
			log.Println("MB初始化失败")
			return false
		}
	} else {
		log.Println("MB已经初始化")
		return true
	}
}

func WkeToString(w WkeString) string {
	var api dllAPI
	api.hFun = dllWkeToString
	api.arg1 = uintptr(w)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}

func SetProxy1(host string, port uint32, ptype uint32) {
	var _proxy Proxy
	n := len(host)
	for i := 0; i < n; i++ {
		_proxy.hostname[i] = host[i]
	}
	_proxy.hostname[n] = 0
	_proxy.port = port
	_proxy.ptype = ptype
	_proxy.user[0] = 0
	_proxy.pwd[0] = 0

	SetProxy(&_proxy)
}

func SetProxy(p *Proxy) {
	var api dllAPI
	api.hFun = dllWkeSetProxy
	api.arg1 = uintptr(unsafe.Pointer(p))
	api.nargs = 1
	callDLLAPI(api)
}
