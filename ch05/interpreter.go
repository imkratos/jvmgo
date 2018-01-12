package main

import (
	"github.com/imkratos/jvmgo/ch05/classfile"
	"github.com/imkratos/jvmgo/ch05/rtda"
	"fmt"
	"github.com/imkratos/jvmgo/ch05/instructions"
	"github.com/imkratos/jvmgo/ch05/instructions/base"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)

	loop(thread, bytecode)

}
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		fmt.Printf("opcode: %v: \n",opcode)
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)

		fmt.Printf("pc:%2d inst:%T %v %v\n", pc, inst, inst,frame)
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
