package references

import (
	"github.com/imkratos/jvmgo/ch11/instructions/base"
	"github.com/imkratos/jvmgo/ch11/rtda"
	"github.com/imkratos/jvmgo/ch11/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	class := classRef.ResolvedClass()

	if ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}

}
