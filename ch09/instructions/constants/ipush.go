package constants

import (
	"github.com/imkratos/jvmgo/ch09/instructions/base"
	"github.com/imkratos/jvmgo/ch09/rtda"
)

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
