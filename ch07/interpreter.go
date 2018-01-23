package main

import (
	"fmt"
	"github.com/imkratos/jvmgo/ch06/instructions"
	"github.com/imkratos/jvmgo/ch06/instructions/base"
	"github.com/imkratos/jvmgo/ch06/rtda"
	"github.com/imkratos/jvmgo/ch06/rtda/heap"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())

}
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}

}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v \n", frame.LocalVars())
		fmt.Printf("OperandStack: %v \n", frame.OperandStack())
		panic(r)
	}
}
