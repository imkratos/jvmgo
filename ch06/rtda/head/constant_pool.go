package head

import (
	"github.com/imkratos/jvmgo/ch06/classfile"
	"fmt"
)

type Constant interface {
}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		}

	}
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}