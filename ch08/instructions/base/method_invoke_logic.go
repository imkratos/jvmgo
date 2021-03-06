package base

import (
	"fmt"
	"github.com/imkratos/jvmgo/ch08/rtda"
	"github.com/imkratos/jvmgo/ch08/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {

	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())

	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}

}
