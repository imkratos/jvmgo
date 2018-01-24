package control

import (
	"github.com/imkratos/jvmgo/ch08/instructions/base"
	"github.com/imkratos/jvmgo/ch08/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
