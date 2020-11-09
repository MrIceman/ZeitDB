package main

import (
	"unsafe"
)

func main() {
	var a = "sasfio"

	for i := 0; i < 10000; i++ {
		a += "asfjoisjfioasf"
	}

	println(unsafe.Sizeof(a))
}
