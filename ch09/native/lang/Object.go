package lang

import (
	"github.com/imkratos/jvmgo/ch09/native"
	"github.com/imkratos/jvmgo/ch09/rtda"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}