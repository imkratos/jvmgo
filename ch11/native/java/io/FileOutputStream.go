package io

import (
	"github.com/imkratos/jvmgo/ch11/rtda"
	"os"
	"unsafe"
)

func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)

	jBytes := b.Data().([]int8)
	goBytes := castInt8ToUint8s(jBytes)
	goBytes = goBytes[off:off+len]
	os.Stdout.Write(goBytes)
}
func castInt8ToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
