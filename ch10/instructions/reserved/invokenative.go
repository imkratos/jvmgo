package reserved

import (
	"github.com/imkratos/jvmgo/ch10/instructions/base"
	"github.com/imkratos/jvmgo/ch10/native"
	"github.com/imkratos/jvmgo/ch10/rtda"
)
import _ "github.com/imkratos/jvmgo/ch10/native/lang"
import _ "github.com/imkratos/jvmgo/ch10/native/sun/misc"

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
