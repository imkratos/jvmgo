package constants

import (
	"github.com/imkratos/jvmgo/ch08/instructions/base"
	"github.com/imkratos/jvmgo/ch08/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
