package references

import (
	"github.com/imkratos/jvmgo/ch08/rtda"
	"github.com/imkratos/jvmgo/ch08/rtda/heap"
	"github.com/imkratos/jvmgo/ch08/instructions/base"
)

type INVOKE_INTERFACE struct {
	index uint
	count uint8
	zero  uint8
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)

	resolvedMethod := methodRef.ResolvedInterfaceMethod()

	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)

	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInclass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)

}
