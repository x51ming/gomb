package gmb

import "syscall"

func freeMBDLL() {
	syscall.FreeLibrary(mb)
}
