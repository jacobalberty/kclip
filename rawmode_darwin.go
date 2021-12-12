package main

import (
	"os"
	"syscall"
	"unsafe"
)

func getTermios(fd uintptr) *syscall.Termios {
	var t syscall.Termios
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TIOCGETA,
		uintptr(unsafe.Pointer(&t)),
		0, 0, 0)

	if err != 0 {
		panic("err")
	}

	return &t
}

func setTermios(fd uintptr, term *syscall.Termios) {
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TIOCSETA,
		uintptr(unsafe.Pointer(term)),
		0, 0, 0)
	if err != 0 {
		panic("err")
	}
}
