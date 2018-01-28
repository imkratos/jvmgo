package misc

import (
	"github.com/imkratos/jvmgo/ch11/instructions/base"
	"github.com/imkratos/jvmgo/ch11/native"
	"github.com/imkratos/jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initializ", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initalizeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)

}
