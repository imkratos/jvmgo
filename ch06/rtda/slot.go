package rtda

import "github.com/imkratos/jvmgo/ch06/rtda/head"

type Slot struct {
	num int32
	ref *head.Object
}
