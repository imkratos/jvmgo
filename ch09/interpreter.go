package main

import (
	"fmt"
	"github.com/imkratos/jvmgo/ch09/instructions"
	"github.com/imkratos/jvmgo/ch09/instructions/base"
	"github.com/imkratos/jvmgo/ch09/rtda"
	"github.com/imkratos/jvmgo/ch09/rtda/heap"
)

func interpret(method *heap.Method, logInst bool, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catchErr(thread)
	loop(thread, logInst)

}
func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}

	return argsArr

}
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}

}
func logInstruction(frame *rtda.Frame, ins base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d   \n", className, methodName, pc)
}
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())

	}
}
