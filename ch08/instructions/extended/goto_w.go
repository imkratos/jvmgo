package extended

import (
	"github.com/imkratos/jvmgo/ch08/instructions/base"
	"github.com/imkratos/jvmgo/ch08/rtda"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
