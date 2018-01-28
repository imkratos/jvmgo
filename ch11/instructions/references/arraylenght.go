package references

import (
	"github.com/imkratos/jvmgo/ch11/instructions/base"
	"github.com/imkratos/jvmgo/ch11/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.long.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
