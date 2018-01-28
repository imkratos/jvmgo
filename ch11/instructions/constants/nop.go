package constants

import (
	"github.com/imkratos/jvmgo/ch11/instructions/base"
	"github.com/imkratos/jvmgo/ch11/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
