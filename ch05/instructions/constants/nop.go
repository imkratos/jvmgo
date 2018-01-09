package constants

import (
	"github.com/imkratos/jvmgo/ch05/instructions/base"
	"github.com/imkratos/jvmgo/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
