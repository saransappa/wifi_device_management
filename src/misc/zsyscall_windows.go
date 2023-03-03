// Code generated by 'go generate'; DO NOT EDIT.

package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwlanapi = windows.NewLazySystemDLL("wlanapi.dll")

	procWlanFreeMemory               = modwlanapi.NewProc("WlanFreeMemory")
	procWlanOpenHandle               = modwlanapi.NewProc("WlanOpenHandle")
	procWlanCloseHandle              = modwlanapi.NewProc("WlanCloseHandle")
	procWlanEnumInterfaces           = modwlanapi.NewProc("WlanEnumInterfaces")
	procWlanSetProfileEapXmlUserData = modwlanapi.NewProc("WlanSetProfileEapXmlUserData")
)

func wlanFreeMemory(memory unsafe.Pointer) {
	syscall.Syscall(procWlanFreeMemory.Addr(), 1, uintptr(memory), 0, 0)
	return
}

func wlanOpenHandle(clientVersion uint32, reserved uintptr, negotiatedVersion *uint32, clientHandle *windows.Handle) (ret error) {
	r0, _, _ := syscall.Syscall6(procWlanOpenHandle.Addr(), 4, uintptr(clientVersion), uintptr(reserved), uintptr(unsafe.Pointer(negotiatedVersion)), uintptr(unsafe.Pointer(clientHandle)), 0, 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func wlanCloseHandle(clientHandle windows.Handle, reserved uintptr) (ret error) {
	r0, _, _ := syscall.Syscall(procWlanCloseHandle.Addr(), 2, uintptr(clientHandle), uintptr(reserved), 0)
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func wlanEnumInterfaces(clientHandle windows.Handle, reserved uintptr, interfaceList **InterfaceInfoList) (ret error) {
	r0, _, _ := syscall.Syscall(procWlanEnumInterfaces.Addr(), 3, uintptr(clientHandle), uintptr(reserved), uintptr(unsafe.Pointer(interfaceList)))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}

func wlanSetProfileEAPXMLUserData(clientHandle windows.Handle, interfaceGUID *windows.GUID, profileName *uint16, flags uint32, eapXMLUserData *uint16, reserved uintptr) (ret error) {
	r0, _, _ := syscall.Syscall6(procWlanSetProfileEapXmlUserData.Addr(), 6, uintptr(clientHandle), uintptr(unsafe.Pointer(interfaceGUID)), uintptr(unsafe.Pointer(profileName)), uintptr(flags), uintptr(unsafe.Pointer(eapXMLUserData)), uintptr(reserved))
	if r0 != 0 {
		ret = syscall.Errno(r0)
	}
	return
}
