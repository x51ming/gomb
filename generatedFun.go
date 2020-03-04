package gmb

import (
	"syscall"
	"unsafe"
)

func (mb *WebView) GetHandleVerifier() {
	var api dllAPI
	api.hFun = dllGetHandleVerifier
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsAddRef() {
	var api dllAPI
	api.hFun = dllJsAddRef
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsArg(es JsExecState, argIdx int32) JsValue {
	var api dllAPI
	api.hFun = dllJsArg
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(argIdx)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsArgCount(es JsExecState) int32 {
	var api dllAPI
	api.hFun = dllJsArgCount
	api.arg1 = uintptr(es)
	api.nargs = 1
	ret := callDLLAPI(api)
	return int32(ret.r1)
}
func JsArgType(es JsExecState, argIdx int32) JsType {
	var api dllAPI
	api.hFun = dllJsArgType
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(argIdx)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return JsType(ret.r1)
}

//char *buffer,size_t size
func JsArrayBuffer(es JsExecState, buffer uintptr, size uintptr) JsValue {
	var api dllAPI
	api.hFun = dllJsArrayBuffer
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(buffer)
	api.arg3 = uintptr(size)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}

//使用了__fastcall
/*
func JsBindFunction(name string, f JsNativeFunction, argcount uint32) {
	var api dllAPI
	api.hFun = dllJsBindFunction
	api.arg1 = str2ptr(name)
	api.arg2 =
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsBindGetter() {
	var api dllAPI
	api.hFun = dllJsBindGetter
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsBindSetter() {
	var api dllAPI
	api.hFun = dllJsBindSetter
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
*/
func JsBoolean(b bool) JsValue {
	var api dllAPI
	api.hFun = dllJsBoolean
	api.arg1 = bool2ptr(b)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}

func JsCall(es JsExecState, fun JsValue, this JsValue, args *JsValue, argcount int32) JsValue {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsCall
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return JsValue(0)
}
func JsCallGlobal() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsCallGlobal
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}

//上次到这里
func JsDeleteObjectProp() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsDeleteObjectProp
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsDouble() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsDouble
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsEmptyArray() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsEmptyArray
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsEmptyObject() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsEmptyObject
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsEval(es JsExecState, str string) JsValue {
	var api dllAPI
	api.hFun = dllJsEval
	api.arg1 = uintptr(es)
	api.arg2 = str2ptr(str)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsEvalExW(es JsExecState, str string, isinclosure bool) JsValue {
	var api dllAPI
	api.hFun = dllJsEvalExW
	api.arg1 = uintptr(es)
	api.arg2 = str2wcharptr(str)
	if isinclosure {
		api.arg3 = 1
	} else {
		api.arg3 = 0
	}
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsEvalW(es JsExecState, str string) JsValue {
	var api dllAPI
	api.hFun = dllJsEvalW
	api.arg1 = uintptr(es)
	api.arg2 = str2wcharptr(str)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsFalse() JsValue {
	var api dllAPI
	api.hFun = dllJsFalse
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsFloat(f float32) JsValue {
	var api dllAPI
	api.hFun = dllJsFloat
	api.arg1 = uintptr(f)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}

//JsData很复杂
/*
func JsFunction(es JsExecState, obj *JsData) JsValue {
	var api dllAPI
	api.hFun = dllJsFunction
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(unsafe.Pointer(obj))
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
*/
func JsGC() {
	var api dllAPI
	api.hFun = dllJsGC
	api.nargs = 0
	callDLLAPI(api)
	return
}
func JsGet(es JsExecState, obj JsValue, prop string) JsValue {
	var api dllAPI
	api.hFun = dllJsGet
	api.arg1 = uintptr(es)
	if Is32 {
		api.arg2 = uintptr(obj & 0xffffffff)
		api.arg3 = uintptr(obj >> 32)
		api.arg4 = str2ptr(prop)
		api.nargs = 4
	} else {
		api.arg2 = uintptr(obj)
		api.arg3 = str2ptr(prop)
		api.nargs = 3
	}
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsGetArrayBuffer() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetArrayBuffer
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetAt() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetAt
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetCallstack() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetCallstack
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetData() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetData
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetGlobal(es JsExecState, prop string) JsValue {
	var api dllAPI
	api.hFun = dllJsGetGlobal
	api.arg1 = uintptr(es)
	api.arg2 = str2ptr(prop)
	api.nargs = 2
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsGetKeys() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetKeys
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetLastErrorIfException() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetLastErrorIfException
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsGetLength() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsGetLength
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}

//通过execstate获得WkeWebView
func JsGetWebView(es JsExecState) WkeWebView {
	var api dllAPI
	api.hFun = dllJsGetWebView
	api.arg1 = uintptr(es)
	api.nargs = 1
	ret := callDLLAPI(api)
	return WkeWebView(ret.r1)
}
func JsGlobalObject(es JsExecState) JsValue {
	var api dllAPI
	api.hFun = dllJsGlobalObject
	api.arg1 = uintptr(es)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsInt(n int32) JsValue {
	var api dllAPI
	api.hFun = dllJsInt
	api.arg1 = uintptr(n)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsIsArray(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsArray
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsBoolean(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsBoolean
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsFalse(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsFalse
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsFunction(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsFunction
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsJsValueValid(es JsExecState, obj JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsJsValueValid
	api.arg1 = uintptr(es)
	api.arg2 = uintptr(obj & 0xffffffff)
	api.arg3 = uintptr(obj >> 32)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsNull(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsNull
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsNumber(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsNumber
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsObject(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsObject
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsString(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsString
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsTrue(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsTrue
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsUndefined(v JsValue) bool {
	var api dllAPI
	api.hFun = dllJsIsUndefined
	api.arg1 = uintptr(v & 0xffffffff)
	api.arg2 = uintptr(v >> 32)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsIsValidExecState(es JsExecState) bool {
	var api dllAPI
	api.hFun = dllJsIsValidExecState
	api.arg1 = uintptr(es)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func JsNull() JsValue {
	var api dllAPI
	api.hFun = dllJsNull
	api.nargs = 0
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsObject() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsObject
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsReleaseRef() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsReleaseRef
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsSet() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsSet
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsSetAt() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsSetAt
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsSetGlobal() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsSetGlobal
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsSetLength() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsSetLength
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsString(es JsExecState, s string) JsValue {
	var api dllAPI
	api.hFun = dllJsString
	api.arg1 = uintptr(es)
	api.arg2 = str2ptr(s)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsStringW(es JsExecState, s string) JsValue {
	var api dllAPI
	api.hFun = dllJsStringW
	api.arg1 = uintptr(es)
	api.arg2 = str2wcharptr(s)
	api.arg3 = uintptr(0)
	api.nargs = 3
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsThrowException() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsThrowException
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsToBoolean(es JsExecState, v JsValue) bool {
	return ptr2bool(f_js_to(dllJsToBoolean, es, v))
}
func JsToDouble() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsToDouble
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsToFloat(es JsExecState, v JsValue) float32 {
	return float32(f_js_to(dllJsToFloat, es, v))
}
func JsToInt(es JsExecState, v JsValue) int32 {
	return int32(f_js_to(dllJsToInt, es, v))
}
func JsToString(es JsExecState, v JsValue) string {
	return ptr2str(f_js_to(dllJsToString, es, v))
}
func JsToStringW() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsToStringW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsToTempString() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsToTempString
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsToTempStringW() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsToTempStringW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsToV8Value() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsToV8Value
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsTrue() JsValue {
	var api dllAPI
	api.hFun = dllJsTrue
	api.nargs = 0
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}
func JsTypeOf() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllJsTypeOf
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func JsUndefined() JsValue {
	var api dllAPI
	api.hFun = dllJsUndefined
	api.nargs = 0
	ret := callDLLAPI(api)
	return ret2JsValue(ret)
}

/*
func (mb *WebView) WkeAddDirtyArea() {
	var api dllAPI
	api.hFun = dllWkeAddDirtyArea
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeAddNpapiPlugin() {
	var api dllAPI
	api.hFun = dllWkeAddNpapiPlugin
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeAddPluginDirectory() {
	var api dllAPI
	api.hFun = dllWkeAddPluginDirectory
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
*/
func (mb *WebView) WkeAwaken() {
	mb.f_void(dllWkeAwaken)
}
func (mb *WebView) WkeCanGoBack() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllWkeCanGoBack
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCanGoForward() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllWkeCanGoForward
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeClearCookie() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllWkeClearCookie
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeConfigure() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllWkeConfigure
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeContentsHeight() int32 {
	var api dllAPI
	api.hFun = dllWkeContentsHeight
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return int32(ret.r1)
}
func (mb *WebView) WkeContentsWidth() int32 {
	var api dllAPI
	api.hFun = dllWkeContentsWidth
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return int32(ret.r1)
}
func (mb *WebView) WkeContextMenuEvent() {
	panic("未实现的函数\n")
	var api dllAPI
	api.hFun = dllWkeContextMenuEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCookieEnabled() bool {
	var api dllAPI
	api.hFun = dllWkeCookieEnabled
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2bool(ret.r1)
}
func (mb *WebView) WkeCopy() {
	mb.f_void(dllWkeCopy)
}
func (mb *WebView) WkeCreateMemBuf() {
	var api dllAPI
	api.hFun = dllWkeCreateMemBuf
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCreateString(s string) WkeString {
	var api dllAPI
	api.hFun = dllWkeCreateString
	api.arg1 = uintptr(str2ptr(s))
	api.arg2 = uintptr(len(s))
	api.nargs = 2
	ret := callDLLAPI(api)
	return WkeString(ret.r1)
}
func (mb *WebView) WkeCreateStringW() {
	var api dllAPI
	api.hFun = dllWkeCreateStringW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCreateStringWithoutNullTermination() {
	var api dllAPI
	api.hFun = dllWkeCreateStringWithoutNullTermination
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCreateWebCustomWindow() {
	var api dllAPI
	api.hFun = dllWkeCreateWebCustomWindow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeCreateWebView() {
	var api dllAPI
	api.nargs = 0
	api.hFun = dllWkeCreateWebView
	ret := callDLLAPI(api)
	mb.hWkeWebView = WkeWebView(ret.r1)
}
func (mb *WebView) WkeCreateWebWindow(wtype int32, parent HWND, x, y, w, h int32) {
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
	mb.hWkeWebView = WkeWebView(ret.r1)
}
func (mb *WebView) WkeCut() {
	mb.f_void(dllWkeCut)
}
func (mb *WebView) WkeDelete() {
	mb.f_void(dllWkeDelete)
}
func (mb *WebView) WkeDeleteString(s WkeString) {
	var api dllAPI
	api.hFun = dllWkeDeleteString
	api.arg1 = uintptr(s)
	api.nargs = 1
	callDLLAPI(api)
}
func (mb *WebView) WkeDestroyWebView() {
	var api dllAPI
	api.hFun = dllWkeDestroyWebView
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDestroyWebWindow() {
	var api dllAPI
	api.hFun = dllWkeDestroyWebWindow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDragTargetDragEnter() {
	var api dllAPI
	api.hFun = dllWkeDragTargetDragEnter
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDragTargetDragLeave() {
	var api dllAPI
	api.hFun = dllWkeDragTargetDragLeave
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDragTargetDragOver() {
	var api dllAPI
	api.hFun = dllWkeDragTargetDragOver
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDragTargetDrop() {
	var api dllAPI
	api.hFun = dllWkeDragTargetDrop
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeDragTargetEnd() {
	var api dllAPI
	api.hFun = dllWkeDragTargetEnd
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorCopy() {
	var api dllAPI
	api.hFun = dllWkeEditorCopy
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorCut() {
	var api dllAPI
	api.hFun = dllWkeEditorCut
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorDelete() {
	var api dllAPI
	api.hFun = dllWkeEditorDelete
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorPaste() {
	var api dllAPI
	api.hFun = dllWkeEditorPaste
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorRedo() {
	var api dllAPI
	api.hFun = dllWkeEditorRedo
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorSelectAll() {
	var api dllAPI
	api.hFun = dllWkeEditorSelectAll
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorUnSelect() {
	var api dllAPI
	api.hFun = dllWkeEditorUnSelect
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEditorUndo() {
	var api dllAPI
	api.hFun = dllWkeEditorUndo
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeEnableWindow(b bool) {
	mb.f_bool_void(dllWkeEnableWindow, b)
}
func (mb *WebView) WkeFinalize() {
	var api dllAPI
	api.hFun = dllWkeFinalize
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireContextMenuEvent() {
	var api dllAPI
	api.hFun = dllWkeFireContextMenuEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireKeyDownEvent() {
	var api dllAPI
	api.hFun = dllWkeFireKeyDownEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireKeyPressEvent() {
	var api dllAPI
	api.hFun = dllWkeFireKeyPressEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireKeyUpEvent() {
	var api dllAPI
	api.hFun = dllWkeFireKeyUpEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireMouseEvent() {
	var api dllAPI
	api.hFun = dllWkeFireMouseEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireMouseWheelEvent() {
	var api dllAPI
	api.hFun = dllWkeFireMouseWheelEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFireWindowsMessage() {
	var api dllAPI
	api.hFun = dllWkeFireWindowsMessage
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFocus() {
	var api dllAPI
	api.hFun = dllWkeFocus
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeFreeMemBuf() {
	var api dllAPI
	api.hFun = dllWkeFreeMemBuf
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGC() {
	var api dllAPI
	api.hFun = dllWkeGC
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetBlinkMainThreadIsolate() {
	var api dllAPI
	api.hFun = dllWkeGetBlinkMainThreadIsolate
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetCaret() {
	var api dllAPI
	api.hFun = dllWkeGetCaret
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetCaretRect() {
	var api dllAPI
	api.hFun = dllWkeGetCaretRect
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetClientHandler() {
	var api dllAPI
	api.hFun = dllWkeGetClientHandler
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetContentAsMarkup() {
	var api dllAPI
	api.hFun = dllWkeGetContentAsMarkup
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetContentHeight() {
	var api dllAPI
	api.hFun = dllWkeGetContentHeight
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetContentWidth() {
	var api dllAPI
	api.hFun = dllWkeGetContentWidth
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetCookie() {
	var api dllAPI
	api.hFun = dllWkeGetCookie
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetCookieW() {
	var api dllAPI
	api.hFun = dllWkeGetCookieW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetCursorInfoType() {
	var api dllAPI
	api.hFun = dllWkeGetCursorInfoType
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetDebugConfig() {
	var api dllAPI
	api.hFun = dllWkeGetDebugConfig
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetDocumentCompleteURL() {
	var api dllAPI
	api.hFun = dllWkeGetDocumentCompleteURL
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetFrameUrl() {
	var api dllAPI
	api.hFun = dllWkeGetFrameUrl
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetGlobalExecByFrame() {
	var api dllAPI
	api.hFun = dllWkeGetGlobalExecByFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetHeight() {
	var api dllAPI
	api.hFun = dllWkeGetHeight
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetHostHWND() {
	var api dllAPI
	api.hFun = dllWkeGetHostHWND
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetMediaVolume() {
	var api dllAPI
	api.hFun = dllWkeGetMediaVolume
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetName() {
	var api dllAPI
	api.hFun = dllWkeGetName
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetSource() {
	var api dllAPI
	api.hFun = dllWkeGetSource
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetString(s WkeString) string {
	var api dllAPI
	api.hFun = dllWkeGetString
	api.arg1 = uintptr(s)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}
func (mb *WebView) WkeGetStringLen(s WkeString) uintptr {
	var api dllAPI
	api.hFun = dllWkeGetStringLen
	api.arg1 = uintptr(s)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ret.r1
}
func (mb *WebView) WkeGetStringW(s WkeString) *uint16 {
	var api dllAPI
	api.hFun = dllWkeGetStringW
	api.arg1 = uintptr(s)
	api.nargs = 1
	ret := callDLLAPI(api)
	return (*uint16)(ptr(ret.r1))
}

func (mb *WebView) WkeGetTempCallbackInfo() *WkeTempCallbackInfo {
	return (*WkeTempCallbackInfo)(ptr(mb.f_void(dllWkeGetTempCallbackInfo)))
}
func (mb *WebView) WkeGetTitle() {
	var api dllAPI
	api.hFun = dllWkeGetTitle
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetTitleW() {
	var api dllAPI
	api.hFun = dllWkeGetTitleW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetURL() {
	var api dllAPI
	api.hFun = dllWkeGetURL
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetUserAgent() {
	var api dllAPI
	api.hFun = dllWkeGetUserAgent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetUserKeyValue() {
	var api dllAPI
	api.hFun = dllWkeGetUserKeyValue
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetVersion() {
	var api dllAPI
	api.hFun = dllWkeGetVersion
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetVersionString() {
	var api dllAPI
	api.hFun = dllWkeGetVersionString
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetViewDC() {
	var api dllAPI
	api.hFun = dllWkeGetViewDC
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetWebViewByNData() {
	var api dllAPI
	api.hFun = dllWkeGetWebViewByNData
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetWebViewForCurrentContext() {
	var api dllAPI
	api.hFun = dllWkeGetWebViewForCurrentContext
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetWebviewId() {
	var api dllAPI
	api.hFun = dllWkeGetWebviewId
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetWidth() {
	var api dllAPI
	api.hFun = dllWkeGetWidth
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGetWindowHandle() HWND {
	return HWND(mb.f_void(dllWkeGetWindowHandle))
}
func (mb *WebView) WkeGetZoomFactor() {
	var api dllAPI
	api.hFun = dllWkeGetZoomFactor
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGlobalExec() {
	var api dllAPI
	api.hFun = dllWkeGlobalExec
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGoBack() {
	var api dllAPI
	api.hFun = dllWkeGoBack
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGoForward() {
	var api dllAPI
	api.hFun = dllWkeGoForward
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGoToIndex() {
	var api dllAPI
	api.hFun = dllWkeGoToIndex
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeGoToOffset() {
	var api dllAPI
	api.hFun = dllWkeGoToOffset
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeHeight() {
	var api dllAPI
	api.hFun = dllWkeHeight
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeInit() {
	var api dllAPI
	api.hFun = dllWkeInit
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeInitialize() {
	var api dllAPI
	api.hFun = dllWkeInitialize
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeInitializeEx() {
	var api dllAPI
	api.hFun = dllWkeInitializeEx
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeInsertCSSByFrame() {
	var api dllAPI
	api.hFun = dllWkeInsertCSSByFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsAwake() {
	var api dllAPI
	api.hFun = dllWkeIsAwake
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsCookieEnabled() {
	var api dllAPI
	api.hFun = dllWkeIsCookieEnabled
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsDirty() {
	var api dllAPI
	api.hFun = dllWkeIsDirty
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsDocumentReady() {
	var api dllAPI
	api.hFun = dllWkeIsDocumentReady
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsInitialize() {
	var api dllAPI
	api.hFun = dllWkeIsInitialize
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoadComplete() {
	var api dllAPI
	api.hFun = dllWkeIsLoadComplete
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoadFailed() {
	var api dllAPI
	api.hFun = dllWkeIsLoadFailed
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoaded() {
	var api dllAPI
	api.hFun = dllWkeIsLoaded
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoading() {
	var api dllAPI
	api.hFun = dllWkeIsLoading
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoadingCompleted() {
	var api dllAPI
	api.hFun = dllWkeIsLoadingCompleted
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoadingFailed() {
	var api dllAPI
	api.hFun = dllWkeIsLoadingFailed
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsLoadingSucceeded() {
	var api dllAPI
	api.hFun = dllWkeIsLoadingSucceeded
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsMainFrame() {
	var api dllAPI
	api.hFun = dllWkeIsMainFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsProcessingUserGesture() {
	var api dllAPI
	api.hFun = dllWkeIsProcessingUserGesture
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsTransparent() {
	var api dllAPI
	api.hFun = dllWkeIsTransparent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsWebRemoteFrame() {
	var api dllAPI
	api.hFun = dllWkeIsWebRemoteFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsWebviewAlive() {
	var api dllAPI
	api.hFun = dllWkeIsWebviewAlive
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeIsWebviewValid() {
	var api dllAPI
	api.hFun = dllWkeIsWebviewValid
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func WkeJsBindFunctionX(name string, fn interface{}, param uintptr, argCount uint32) {
	var api dllAPI
	api.hFun = dllWkeJsBindFunction
	api.arg1 = str2ptr(name)
	api.arg2 = syscall.NewCallbackCDecl(fn)
	api.arg3 = param
	api.arg4 = uintptr(argCount)
	api.nargs = 4
	callDLLAPI(api)
}
func WkeJsBindFunction(name string, fn WkeJsNativeFunction, param uintptr, argCount uint32) {
	var api dllAPI
	api.hFun = dllWkeJsBindFunction
	api.arg1 = str2ptr(name)
	api.arg2 = syscall.NewCallbackCDecl(fn)
	api.arg3 = param
	api.arg4 = uintptr(argCount)
	api.nargs = 4
	callDLLAPI(api)
}
func WkeJsBindGetter(name string, fn WkeJsNativeFunction, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeJsBindGetter
	api.arg1 = str2ptr(name)
	api.arg2 = syscall.NewCallbackCDecl(fn)
	api.arg3 = param
	api.nargs = 3
	callDLLAPI(api)
}
func WkeJsBindSetter(name string, fn WkeJsNativeFunction, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeJsBindSetter
	api.arg1 = str2ptr(name)
	api.arg2 = syscall.NewCallbackCDecl(fn)
	api.arg3 = param
	api.nargs = 3
	callDLLAPI(api)
}
func (mb *WebView) WkeKeyDown() {
	var api dllAPI
	api.hFun = dllWkeKeyDown
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeKeyPress() {
	var api dllAPI
	api.hFun = dllWkeKeyPress
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeKeyUp() {
	var api dllAPI
	api.hFun = dllWkeKeyUp
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeKillFocus() {
	var api dllAPI
	api.hFun = dllWkeKillFocus
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLayoutIfNeeded() {
	var api dllAPI
	api.hFun = dllWkeLayoutIfNeeded
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadFile(fn string) {
	var api dllAPI
	api.hFun = dllWkeLoadFile
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(fn)
	api.nargs = 2
	callDLLAPI(api)
}
func (mb *WebView) WkeLoadFileW(fn string) {
	var api dllAPI
	api.hFun = dllWkeLoadFileW
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2wcharptr(fn)
	api.nargs = 2
	callDLLAPI(api)
}
func (mb *WebView) WkeLoadHTML() {
	var api dllAPI
	api.hFun = dllWkeLoadHTML
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadHTMLW() {
	var api dllAPI
	api.hFun = dllWkeLoadHTMLW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadHtmlWithBaseUrl() {
	var api dllAPI
	api.hFun = dllWkeLoadHtmlWithBaseUrl
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadURL() {
	var api dllAPI
	api.hFun = dllWkeLoadURL
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadURLW() {
	var api dllAPI
	api.hFun = dllWkeLoadURLW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeLoadW() {
	var api dllAPI
	api.hFun = dllWkeLoadW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeMediaVolume() {
	var api dllAPI
	api.hFun = dllWkeMediaVolume
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeMouseEvent() {
	var api dllAPI
	api.hFun = dllWkeMouseEvent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeMouseWheel() {
	var api dllAPI
	api.hFun = dllWkeMouseWheel
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeMoveToCenter() {
	var api dllAPI
	api.hFun = dllWkeMoveToCenter
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeMoveWindow() {
	var api dllAPI
	api.hFun = dllWkeMoveWindow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetAddHTTPHeaderFieldToUrlRequest() {
	var api dllAPI
	api.hFun = dllWkeNetAddHTTPHeaderFieldToUrlRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func WkeNetCancelRequest(job WkeNetJob) {
	var api dllAPI
	api.hFun = dllWkeNetCancelRequest
	api.arg1 = uintptr(job)
	api.nargs = 1
	callDLLAPI(api)
}
func (mb *WebView) WkeNetCancelWebUrlRequest() {
	var api dllAPI
	api.hFun = dllWkeNetCancelWebUrlRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetChangeRequestUrl() {
	var api dllAPI
	api.hFun = dllWkeNetChangeRequestUrl
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetContinueJob() {
	var api dllAPI
	api.hFun = dllWkeNetContinueJob
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetCopyWebUrlRequest() {
	var api dllAPI
	api.hFun = dllWkeNetCopyWebUrlRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetCreatePostBodyElement() {
	var api dllAPI
	api.hFun = dllWkeNetCreatePostBodyElement
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetCreatePostBodyElements() {
	var api dllAPI
	api.hFun = dllWkeNetCreatePostBodyElements
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetCreateWebUrlRequest() {
	var api dllAPI
	api.hFun = dllWkeNetCreateWebUrlRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetCreateWebUrlRequest2() {
	var api dllAPI
	api.hFun = dllWkeNetCreateWebUrlRequest2
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetDeleteBlinkWebURLRequestPtr() {
	var api dllAPI
	api.hFun = dllWkeNetDeleteBlinkWebURLRequestPtr
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetFreePostBodyElement() {
	var api dllAPI
	api.hFun = dllWkeNetFreePostBodyElement
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetFreePostBodyElements() {
	var api dllAPI
	api.hFun = dllWkeNetFreePostBodyElements
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetExpectedContentLength() {
	var api dllAPI
	api.hFun = dllWkeNetGetExpectedContentLength
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetFavicon() {
	var api dllAPI
	api.hFun = dllWkeNetGetFavicon
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetHTTPHeaderField(job WkeNetJob, key string) string {
	var api dllAPI
	api.hFun = dllWkeNetGetHTTPHeaderField
	api.arg1 = uintptr(job)
	api.arg2 = uintptr(str2ptr(key))
	api.nargs = 2
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}
func (mb *WebView) WkeNetGetHTTPHeaderFieldFromResponse(job WkeNetJob, key string) string {
	var api dllAPI
	api.hFun = dllWkeNetGetHTTPHeaderFieldFromResponse
	api.arg1 = uintptr(job)
	api.arg2 = uintptr(str2ptr(key))
	api.nargs = 2
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}
func (mb *WebView) WkeNetGetHttpStatusCode() {
	var api dllAPI
	api.hFun = dllWkeNetGetHttpStatusCode
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetMIMEType() {
	var api dllAPI
	api.hFun = dllWkeNetGetMIMEType
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetPostBody() {
	var api dllAPI
	api.hFun = dllWkeNetGetPostBody
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetRawHttpHead() {
	var api dllAPI
	api.hFun = dllWkeNetGetRawHttpHead
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetRawResponseHead() {
	var api dllAPI
	api.hFun = dllWkeNetGetRawResponseHead
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetGetRequestMethod(job WkeNetJob) WkeRequestType {
	var api dllAPI
	api.hFun = dllWkeNetGetRequestMethod
	api.arg1 = uintptr(job)
	api.nargs = 1
	ret := callDLLAPI(api)
	return WkeRequestType(ret.r1)
}
func (mb *WebView) WkeNetGetResponseUrl() {
	var api dllAPI
	api.hFun = dllWkeNetGetResponseUrl
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func WkeNetGetUrlByJob(job WkeNetJob) string {
	var api dllAPI
	api.hFun = dllWkeNetGetUrlByJob
	api.arg1 = uintptr(job)
	api.nargs = 1
	ret := callDLLAPI(api)
	return ptr2str(ret.r1)
}
func (mb *WebView) WkeNetHoldJobToAsynCommit() {
	var api dllAPI
	api.hFun = dllWkeNetHoldJobToAsynCommit
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetHookRequest() {
	var api dllAPI
	api.hFun = dllWkeNetHookRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetOnResponse() {
	var api dllAPI
	api.hFun = dllWkeNetOnResponse
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetSetData() {
	var api dllAPI
	api.hFun = dllWkeNetSetData
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetSetHTTPHeaderField() {
	var api dllAPI
	api.hFun = dllWkeNetSetHTTPHeaderField
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetSetMIMEType() {
	var api dllAPI
	api.hFun = dllWkeNetSetMIMEType
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNetStartUrlRequest() {
	var api dllAPI
	api.hFun = dllWkeNetStartUrlRequest
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeNodeOnCreateProcess() {
	var api dllAPI
	api.hFun = dllWkeNodeOnCreateProcess
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnAlertBox(cb WkeAlertBoxCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnAlertBox, cb, param)
}
func (mb *WebView) WkeOnConfirmBox(cb WkeConfirmBoxCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnConfirmBox, cb, param)
}
func (mb *WebView) WkeOnConsole() {
	var api dllAPI
	api.hFun = dllWkeOnConsole
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnContextMenuItemClick() {
	var api dllAPI
	api.hFun = dllWkeOnContextMenuItemClick
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnCreateView(cb WkeCreateViewCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnCreateView, cb, param)
}
func (mb *WebView) WkeOnDidCreateScriptContext() {
	var api dllAPI
	api.hFun = dllWkeOnDidCreateScriptContext
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}

func (mb *WebView) WkeOnDocumentReady(cb WkeDocumentReadyCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnDocumentReady, cb, param)
}
func (mb *WebView) WkeOnDocumentReady2(cb WkeDocumentReady2Callback, param uintptr) {
	mb.f_cb_p(dllWkeOnDocumentReady2, cb, param)
}
func (mb *WebView) WkeOnDownload(cb WkeDownloadCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnDownload, cb, param)
}
func (mb *WebView) WkeOnDownload2(cb WkeDownloadCallback2, param uintptr) {
	mb.f_cb_p(dllWkeOnDownload2, cb, param)
}
func (mb *WebView) WkeOnDraggableRegionsChanged() {
	var api dllAPI
	api.hFun = dllWkeOnDraggableRegionsChanged
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnLoadUrlBegin(cb WkeLoadUrlBeginCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnLoadUrlBegin, cb, param)
}
func (mb *WebView) WkeOnLoadUrlEnd(cb WkeLoadUrlEndCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnLoadUrlEnd, cb, param)
}
func (mb *WebView) WkeOnLoadUrlFail(cb WkeLoadUrlFailCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnLoadUrlFail, cb, param)
}
func (mb *WebView) WkeOnLoadingFinish() {
	var api dllAPI
	api.hFun = dllWkeOnLoadingFinish
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnMouseOverUrlChanged() {
	var api dllAPI
	api.hFun = dllWkeOnMouseOverUrlChanged
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}

//CB返回1表示继续浏览，0表示取消
func (mb *WebView) WkeOnNavigation(cb WkeNavigationCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnNavigation, cb, param)
}
func (mb *WebView) WkeOnOtherLoad(cb WkeOnOtherLoadCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnOtherLoad, cb, param)
}
func (mb *WebView) WkeOnPaintBitUpdated() {
	var api dllAPI
	api.hFun = dllWkeOnPaintBitUpdated
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnPaintUpdated() {
	var api dllAPI
	api.hFun = dllWkeOnPaintUpdated
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnPluginFind() {
	var api dllAPI
	api.hFun = dllWkeOnPluginFind
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnPrint() {
	var api dllAPI
	api.hFun = dllWkeOnPrint
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnPromptBox() {
	var api dllAPI
	api.hFun = dllWkeOnPromptBox
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnStartDragging() {
	var api dllAPI
	api.hFun = dllWkeOnStartDragging
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnTitleChanged(cb WkeTitleChangedCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnTitleChanged, cb, param)
}
func (mb *WebView) WkeOnURLChanged(cb WkeURLChangedCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnURLChanged, cb, param)
}
func (mb *WebView) WkeOnURLChanged2(cb WkeURLChangedCallback2, param uintptr) {
	mb.f_cb_p(dllWkeOnURLChanged2, cb, param)
}
func (mb *WebView) WkeOnWillMediaLoad() {
	var api dllAPI
	api.hFun = dllWkeOnWillMediaLoad
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnWillReleaseScriptContext() {
	var api dllAPI
	api.hFun = dllWkeOnWillReleaseScriptContext
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeOnWindowClosing(cb WkeWindowClosingCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnWindowClosing, cb, param)
}
func (mb *WebView) WkeOnWindowDestroy(cb WkeWindowDestroyCallback, param uintptr) {
	mb.f_cb_p(dllWkeOnWindowDestroy, cb, param)
}
func (mb *WebView) WkePaint() {
	var api dllAPI
	api.hFun = dllWkePaint
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePaint2() {
	var api dllAPI
	api.hFun = dllWkePaint2
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePaste() {
	mb.f_void(dllWkePaste)
}
func (mb *WebView) WkePerformCookieCommand() {
	var api dllAPI
	api.hFun = dllWkePerformCookieCommand
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePluginListBuilderAddFileExtensionToLastMediaType() {
	var api dllAPI
	api.hFun = dllWkePluginListBuilderAddFileExtensionToLastMediaType
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePluginListBuilderAddMediaTypeToLastPlugin() {
	var api dllAPI
	api.hFun = dllWkePluginListBuilderAddMediaTypeToLastPlugin
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePluginListBuilderAddPlugin() {
	var api dllAPI
	api.hFun = dllWkePluginListBuilderAddPlugin
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePostURL(url string, data *byte, length uint32) {
	var api dllAPI
	api.hFun = dllWkePostURL
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(url)
	api.arg3 = uintptr(unsafe.Pointer(data))
	api.arg4 = uintptr(length)
	api.nargs = 4
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePostURLW() {
	panic("W")
	var api dllAPI
	api.hFun = dllWkePostURLW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkePrintToBitmap() {
	var api dllAPI
	api.hFun = dllWkePrintToBitmap
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeRegisterEmbedderCustomElement() {
	var api dllAPI
	api.hFun = dllWkeRegisterEmbedderCustomElement
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeReload() {
	mb.f_void(dllWkeReload)
}
func (mb *WebView) WkeRepaintIfNeeded() {
	var api dllAPI
	api.hFun = dllWkeRepaintIfNeeded
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeResize() {
	var api dllAPI
	api.hFun = dllWkeResize
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeResizeWindow() {
	var api dllAPI
	api.hFun = dllWkeResizeWindow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeRunJS() {
	var api dllAPI
	api.hFun = dllWkeRunJS
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeRunJSW() {
	var api dllAPI
	api.hFun = dllWkeRunJSW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeRunJsByFrame() {
	var api dllAPI
	api.hFun = dllWkeRunJsByFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeRunMessageLoop() {
	var api dllAPI
	api.hFun = dllWkeRunMessageLoop
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSaveMemoryCache() {
	var api dllAPI
	api.hFun = dllWkeSaveMemoryCache
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeScreenshot() {
	var api dllAPI
	api.hFun = dllWkeScreenshot
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSelectAll() {
	var api dllAPI
	api.hFun = dllWkeSelectAll
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetClientHandler() {
	var api dllAPI
	api.hFun = dllWkeSetClientHandler
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetContextMenuEnabled() {
	var api dllAPI
	api.hFun = dllWkeSetContextMenuEnabled
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetContextMenuItemShow() {
	var api dllAPI
	api.hFun = dllWkeSetContextMenuItemShow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetCookie() {
	var api dllAPI
	api.hFun = dllWkeSetCookie
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetCookieEnabled() {
	var api dllAPI
	api.hFun = dllWkeSetCookieEnabled
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetCookieJarFullPath() {
	var api dllAPI
	api.hFun = dllWkeSetCookieJarFullPath
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetCookieJarPath() {
	var api dllAPI
	api.hFun = dllWkeSetCookieJarPath
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetCspCheckEnable(b bool) {
	mb.f_bool_void(dllWkeSetCspCheckEnable, b)
}
func (mb *WebView) WkeSetCursorInfoType() {
	var api dllAPI
	api.hFun = dllWkeSetCursorInfoType
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDebugConfig(debugstr string, param string) {
	var api dllAPI
	api.hFun = dllWkeSetDebugConfig
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(debugstr)
	api.arg3 = str2ptr(param)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDeviceParameter() {
	var api dllAPI
	api.hFun = dllWkeSetDeviceParameter
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDirty() {
	var api dllAPI
	api.hFun = dllWkeSetDirty
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDragDropEnable() {
	var api dllAPI
	api.hFun = dllWkeSetDragDropEnable
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDragEnable() {
	var api dllAPI
	api.hFun = dllWkeSetDragEnable
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetDragFiles() {
	var api dllAPI
	api.hFun = dllWkeSetDragFiles
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetEditable() {
	var api dllAPI
	api.hFun = dllWkeSetEditable
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetFileSystem() {
	var api dllAPI
	api.hFun = dllWkeSetFileSystem
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetFocus() {
	var api dllAPI
	api.hFun = dllWkeSetFocus
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetHandle() {
	var api dllAPI
	api.hFun = dllWkeSetHandle
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetHandleOffset() {
	var api dllAPI
	api.hFun = dllWkeSetHandleOffset
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetHeadlessEnabled() {
	var api dllAPI
	api.hFun = dllWkeSetHeadlessEnabled
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetLanguage() {
	var api dllAPI
	api.hFun = dllWkeSetLanguage
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetLocalStorageFullPath() {
	var api dllAPI
	api.hFun = dllWkeSetLocalStorageFullPath
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetMediaPlayerFactory() {
	var api dllAPI
	api.hFun = dllWkeSetMediaPlayerFactory
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetMediaVolume() {
	var api dllAPI
	api.hFun = dllWkeSetMediaVolume
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetMemoryCacheEnable(b bool) {
	mb.f_bool_void(dllWkeSetMemoryCacheEnable, b)
}
func (mb *WebView) WkeSetMouseEnabled(b bool) {
	mb.f_bool_void(dllWkeSetMouseEnabled, b)
}
func (mb *WebView) WkeSetName() {
	var api dllAPI
	api.hFun = dllWkeSetName
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetNavigationToNewWindowEnable(b bool) {
	mb.f_bool_void(dllWkeSetNavigationToNewWindowEnable, b)
}
func (mb *WebView) WkeSetNpapiPluginsEnabled(b bool) {
	mb.f_bool_void(dllWkeSetNpapiPluginsEnabled, b)
}
func (mb *WebView) WkeSetProxy() {
	var api dllAPI
	api.hFun = dllWkeSetProxy
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetResourceGc() {
	var api dllAPI
	api.hFun = dllWkeSetResourceGc
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetString() {
	var api dllAPI
	api.hFun = dllWkeSetString
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetStringW() {
	var api dllAPI
	api.hFun = dllWkeSetStringW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetStringWithoutNullTermination() {
	var api dllAPI
	api.hFun = dllWkeSetStringWithoutNullTermination
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetTouchEnabled(b bool) {
	mb.f_bool_void(dllWkeSetTouchEnabled, b)
}
func (mb *WebView) WkeSetTransparent() {
	var api dllAPI
	api.hFun = dllWkeSetTransparent
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetUIThreadCallback() {
	var api dllAPI
	api.hFun = dllWkeSetUIThreadCallback
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetUserAgent(s string) {
	var api dllAPI
	api.hFun = dllWkeSetUserAgent
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2ptr(s)
	api.nargs = 2
	callDLLAPI(api)
}
func (mb *WebView) WkeSetUserAgentW() {
	var api dllAPI
	api.hFun = dllWkeSetUserAgentW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetUserKeyValue() {
	var api dllAPI
	api.hFun = dllWkeSetUserKeyValue
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetViewNetInterface() {
	var api dllAPI
	api.hFun = dllWkeSetViewNetInterface
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetViewProxy() {
	var api dllAPI
	api.hFun = dllWkeSetViewProxy
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetViewSettings() {
	var api dllAPI
	api.hFun = dllWkeSetViewSettings
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetWebViewName() {
	var api dllAPI
	api.hFun = dllWkeSetWebViewName
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetWindowTitle() {
	var api dllAPI
	api.hFun = dllWkeSetWindowTitle
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetWindowTitleW() {
	var api dllAPI
	api.hFun = dllWkeSetWindowTitleW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSetZoomFactor() {
	var api dllAPI
	api.hFun = dllWkeSetZoomFactor
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeShowDevtools(path string, cb WkeOnShowDevtoolsCallback, param uintptr) {
	var api dllAPI
	api.hFun = dllWkeShowDevtools
	api.arg1 = uintptr(mb.hWkeWebView)
	api.arg2 = str2wcharptr(path)
	api.arg3 = uintptr(syscall.NewCallbackCDecl(cb))
	api.arg4 = param
	api.nargs = 4
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeShowWindow() {
	var api dllAPI
	api.hFun = dllWkeShowWindow
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeShutdown() {
	var api dllAPI
	api.hFun = dllWkeShutdown
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeShutdownForDebug() {
	var api dllAPI
	api.hFun = dllWkeShutdownForDebug
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeSleep() {
	var api dllAPI
	api.hFun = dllWkeSleep
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeStopLoading() {
	var api dllAPI
	api.hFun = dllWkeStopLoading
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeTitle() {
	var api dllAPI
	api.hFun = dllWkeTitle
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeTitleW() {
	var api dllAPI
	api.hFun = dllWkeTitleW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeToString() {
	var api dllAPI
	api.hFun = dllWkeToString
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeToStringW() {
	var api dllAPI
	api.hFun = dllWkeToStringW
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUnfocus() {
	var api dllAPI
	api.hFun = dllWkeUnfocus
	api.arg1 = uintptr(mb.hWkeWebView)
	api.nargs = 1
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUpdate() {
	var api dllAPI
	api.hFun = dllWkeUpdate
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilBase64Decode() {
	var api dllAPI
	api.hFun = dllWkeUtilBase64Decode
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilBase64Encode() {
	var api dllAPI
	api.hFun = dllWkeUtilBase64Encode
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilCreateV8Snapshot() {
	var api dllAPI
	api.hFun = dllWkeUtilCreateV8Snapshot
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilDecodeURLEscape() {
	var api dllAPI
	api.hFun = dllWkeUtilDecodeURLEscape
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilEncodeURLEscape() {
	var api dllAPI
	api.hFun = dllWkeUtilEncodeURLEscape
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilPrintToPdf() {
	var api dllAPI
	api.hFun = dllWkeUtilPrintToPdf
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilRelasePrintPdfDatas() {
	var api dllAPI
	api.hFun = dllWkeUtilRelasePrintPdfDatas
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilSerializeToMHTML() {
	var api dllAPI
	api.hFun = dllWkeUtilSerializeToMHTML
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeUtilSetUiCallback() {
	var api dllAPI
	api.hFun = dllWkeUtilSetUiCallback
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeVersion() {
	var api dllAPI
	api.hFun = dllWkeVersion
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeVersionString() {
	var api dllAPI
	api.hFun = dllWkeVersionString
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeVisitAllCookie() {
	var api dllAPI
	api.hFun = dllWkeVisitAllCookie
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeWake() {
	var api dllAPI
	api.hFun = dllWkeWake
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeWebFrameGetMainFrame() {
	var api dllAPI
	api.hFun = dllWkeWebFrameGetMainFrame
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeWebFrameGetMainWorldScriptContext() {
	var api dllAPI
	api.hFun = dllWkeWebFrameGetMainWorldScriptContext
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeWebViewName() {
	var api dllAPI
	api.hFun = dllWkeWebViewName
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeWidth() {
	var api dllAPI
	api.hFun = dllWkeWidth
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
func (mb *WebView) WkeZoomFactor() {
	var api dllAPI
	api.hFun = dllWkeZoomFactor
	api.arg1 = uintptr(0)
	api.arg2 = uintptr(0)
	api.arg3 = uintptr(0)
	api.nargs = 3
	callDLLAPI(api)
	return
}
