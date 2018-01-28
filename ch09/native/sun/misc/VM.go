package misc

import (
	"github.com/imkratos/jvmgo/ch09/native"
	"github.com/imkratos/jvmgo/ch09/rtda"
	"github.com/imkratos/jvmgo/ch09/rtda/heap"
	"github.com/imkratos/jvmgo/ch09/instructions/base"
)

func init() {
	native.Register("sun/misc/VM", "initializ", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")

	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)

}
