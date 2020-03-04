package gmb

import (
	"errors"
	"log"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/lxn/win"
)

const ExitNarg uintptr = 99

type dllAPI struct {
	nargs uintptr
	hFun  uintptr
	arg1  uintptr
	arg2  uintptr
	arg3  uintptr
	arg4  uintptr
	arg5  uintptr
	arg6  uintptr
}

type dllReturn struct {
	r1, r2 uintptr
	err    error
}

var (
	DLLFREED     chan bool
	_dllCommand  chan dllAPI
	_dllReturn   chan dllReturn
	_lc          sync.Mutex //操作回调函数记数
	_callBackNum uint32     //记录了运行的回调函数。（如果不为0，表示有操作是在回调函数里的，调用DLL时，需要判断线程）
	_mThreadId   uint32     //主线程ID
)

func doSysCall(dllCommand dllAPI) dllReturn {
	var ret dllReturn
	switch dllCommand.nargs {
	case 0:
		ret.r1, ret.r2, ret.err = syscall.Syscall(dllCommand.hFun, 0, 0,
			0, 0)
	case 1:
		ret.r1, ret.r2, ret.err = syscall.Syscall(dllCommand.hFun, 1, dllCommand.arg1,
			0, 0)
	case 2:
		ret.r1, ret.r2, ret.err = syscall.Syscall(dllCommand.hFun, 2, dllCommand.arg1,
			dllCommand.arg2, 0)
	case 3:
		ret.r1, ret.r2, ret.err = syscall.Syscall(dllCommand.hFun, 3, dllCommand.arg1,
			dllCommand.arg2, dllCommand.arg3)
	case 4:
		ret.r1, ret.r2, ret.err = syscall.Syscall6(dllCommand.hFun, 6, dllCommand.arg1,
			dllCommand.arg2, dllCommand.arg3, dllCommand.arg4,
			0, 0)
	case 5:
		ret.r1, ret.r2, ret.err = syscall.Syscall6(dllCommand.hFun, 6, dllCommand.arg1,
			dllCommand.arg2, dllCommand.arg3, dllCommand.arg4,
			dllCommand.arg5, 0)
	case 6:
		ret.r1, ret.r2, ret.err = syscall.Syscall6(dllCommand.hFun, 6, dllCommand.arg1,
			dllCommand.arg2, dllCommand.arg3, dllCommand.arg4,
			dllCommand.arg5, dllCommand.arg6)
	default:
		log.Panicln("argn????", dllCommand.nargs)
		ret.err = errors.New("错误")
		return ret
	}
	return ret
}

//MessageLoop并接受调用
func RunLoop() {
	DLLFREED = make(chan bool, 1)
	_dllCommand = make(chan dllAPI)
	_dllReturn = make(chan dllReturn)
	go func() {
		runtime.LockOSThread()
		_mThreadId = getCurrentThreadId()
		for {
			select {
			case dllCommand := <-_dllCommand:
				if dllCommand.nargs == ExitNarg {
					_dllReturn <- dllReturn{}
					goto EndLoop
				}
				_dllReturn <- doSysCall(dllCommand)
			default:
				var msg win.MSG
				time.Sleep(time.Millisecond * 2)
				if peekMessage(&msg, 0, 0, 0, 1) {
					translateMessage(&msg)
					dispatchMessage(&msg)
					if msg.Message == WM_QUIT {
						log.Println("WM_QUIT RECEIVED")
						goto EndLoop
					}
				}
			}
		}
	EndLoop:
		log.Println("卸载DLL!")
		freeMBDLL() //卸载DLL
		DLLFREED <- true
	}()
}

//内部支持
//调用DLL，主要解决在回调函数里操作时，channel 死锁的问题
func callDLLAPI(api dllAPI) dllReturn {
	_lc.Lock()
	callBackNum := _callBackNum
	_lc.Unlock()
	if callBackNum > 0 {
		//表示有操作是在回调函数里进行的
		if _mThreadId == getCurrentThreadId() {
			//当前在主线程，表示此操作是在回调里进行的
			//假设在回调里面不会主动Finalize
			if api.nargs == ExitNarg {
				postQuitMessage(0)
				return dllReturn{}
			}
			return doSysCall(api)
		} else {
			//非回调里进行的
			_dllCommand <- api
			return <-_dllReturn
		}
	} else {
		//非回调里进行的
		_dllCommand <- api
		return <-_dllReturn
	}
}

//回调函数开始时调用
func StartCallBack() {
	_lc.Lock()
	defer _lc.Unlock()
	_callBackNum++
}

//回调函数结束时调用
func EndCallBack() {
	_lc.Lock()
	defer _lc.Unlock()
	_callBackNum--
}
