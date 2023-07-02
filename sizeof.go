#! /bin/sh
TMPGO="$(mktemp).go" && tail -n +3 "$0" > "$TMPGO" && go run "$TMPGO"; rm "$TMPGO"; exit
// File: "sizeof.go"

package main

import (
	"errors"
	"fmt"
	"io"
	"time"
	"unsafe"
  "runtime"
)

type rd int
type rw int

type any interface{} // old golang

func (rd *rd) Read(p []byte) (n int, err error) { // io.Reader interface
	return 0, nil
}

func (rw *rw) Read(p []byte) (n int, err error) { // io.ReadWriter interface
	return 0, nil
}

func (rw *rw) Write(p []byte) (n int, err error) { // io.ReadWriter interface
	return 0, nil
}

func main() {
	fmt.Println(Indian(), "Indian")
  
  fmt.Println("GOOS                  =", runtime.GOOS)
  fmt.Println("GOARCH                =", runtime.GOARCH)
  fmt.Println("GOMAXPROCS            =", runtime.GOMAXPROCS(0))
  fmt.Println("NumCPU                =", runtime.NumCPU())

	fmt.Println("sizeof(bool)          =", unsafe.Sizeof(bool(false)))
	fmt.Println("sizeof(byte)          =", unsafe.Sizeof(byte(0)))
	fmt.Println("sizeof(int8)          =", unsafe.Sizeof(int8(0)))
	fmt.Println("sizeof(int16)         =", unsafe.Sizeof(int16(0)))
	fmt.Println("sizeof(int32)         =", unsafe.Sizeof(int32(0)))
	fmt.Println("sizeof(int64)         =", unsafe.Sizeof(int64(0)))
	fmt.Println("sizeof(int)           =", unsafe.Sizeof(int(0)))
	fmt.Println("sizeof(rune)          =", unsafe.Sizeof(rune(0)))
	fmt.Println("sizeof(uintptr)       =", unsafe.Sizeof(uintptr(0)))
	fmt.Println("sizeof(*int)          =", unsafe.Sizeof(new(int)))
	fmt.Println("sizeof(float32)       =", unsafe.Sizeof(float32(0.)))
	fmt.Println("sizeof(float64)       =", unsafe.Sizeof(float64(0.)))
	fmt.Println("sizeof(complex64)     =", unsafe.Sizeof(complex64(0+0i)))
	fmt.Println("sizeof(complex128)    =", unsafe.Sizeof(complex128(0+0i)))
	fmt.Println("sizeof(string)        =", unsafe.Sizeof(string("")))
	fmt.Println("sizeof([1]int)        =", unsafe.Sizeof([1]int{0}))
	fmt.Println("sizeof([]int)         =", unsafe.Sizeof([]int{}))
	fmt.Println("sizeof(map[int]int)   =", unsafe.Sizeof(map[int]int{0: 0}))
	fmt.Println("sizeof(func(){})      =", unsafe.Sizeof(func() {}))

	fmt.Println("sizeof(error(nil))    =", unsafe.Sizeof(error(nil)))
	fmt.Println(`sizeof(error(""))     =`, unsafe.Sizeof(errors.New("")))

	fmt.Println("sizeof(any(0))        =", unsafe.Sizeof(any(0)))
	fmt.Println("sizeof(io.Reader)     =", unsafe.Sizeof(io.Reader(new(rd))))
	fmt.Println("sizeof(io.ReadWriter) =", unsafe.Sizeof(io.ReadWriter(new(rw))))

	fmt.Println("sizeof(time.Time)     =", unsafe.Sizeof(time.Now()))

	c := make(chan int)
	fmt.Println("sizeof(chan int)      =", unsafe.Sizeof(c))
}

func Indian() string {
	i := uint32(0x12345678)
	p := unsafe.Pointer(&i)
	w := *(*uint16)(p)
	switch w {
	case 0x1234:
		return "Big"
	case 0x5678:
		return "Little"
	default:
		return "Unknown"
	}
}

// EOF
