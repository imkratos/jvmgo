package control

import (
	"github.com/imkratos/jvmgo/ch05/instructions/base"
	"github.com/imkratos/jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
