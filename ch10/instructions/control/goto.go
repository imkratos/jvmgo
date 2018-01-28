package control

import (
	"github.com/imkratos/jvmgo/ch10/instructions/base"
	"github.com/imkratos/jvmgo/ch10/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
