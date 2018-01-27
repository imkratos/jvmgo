package heap

import (
	"github.com/imkratos/jvmgo/ch09/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref

}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.ResolvedMethodRef()
	}
	return self.method
}
func (self *MethodRef) ResolvedMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)

	if method == nil {
		panic("java.langNoSuchMethodError")
	}

	if method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method

}
func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInclass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
