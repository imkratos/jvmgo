package references

import (
	"github.com/imkratos/jvmgo/ch06/instructions/base"
	"github.com/imkratos/jvmgo/ch06/rtda"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
