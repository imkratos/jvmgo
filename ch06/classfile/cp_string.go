package classfile

import "fmt"

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
	fmt.Printf("aaa-------- %s", self.stringIndex)
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)

}
