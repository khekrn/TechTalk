package main

import (
	"fmt"
	"unsafe"

	"github.com/khekrn/cpu-cache/dod"
)

func main() {
	b := dod.I1{}

	fmt.Println(unsafe.Sizeof(b))

	c := dod.I2{}

	fmt.Println(unsafe.Sizeof(c))
}
