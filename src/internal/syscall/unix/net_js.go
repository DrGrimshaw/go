// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package unix

import (
	"syscall"
	_ "unsafe"
)

func RecvfromInet4(fd int, p []byte, flags int, from *syscall.SockaddrInet4) (int, error) {
	return 0, syscall.ENOSYS
}

func RecvfromInet6(fd int, p []byte, flags int, from *syscall.SockaddrInet6) (n int, err error) {
	return 0, syscall.ENOSYS
}

func SendtoInet4(fd int, p []byte, flags int, to syscall.SockaddrInet4) (err error) {
	return syscall.ENOSYS
}

func SendtoInet6(fd int, p []byte, flags int, to syscall.SockaddrInet6) (err error) {
	return syscall.ENOSYS
}
