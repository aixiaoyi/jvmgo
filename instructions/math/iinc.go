package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IINC struct {
	Index    uint
	ConstVal int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.ConstVal = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	localVars.SetInt(self.Index, val+self.ConstVal)
}
