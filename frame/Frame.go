// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package frame

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Frame struct {
	_tab flatbuffers.Table
}

func GetRootAsFrame(buf []byte, offset flatbuffers.UOffsetT) *Frame {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Frame{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Frame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Frame) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Frame) Ver() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frame) MutateVer(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *Frame) Cmd() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frame) MutateCmd(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Frame) Len() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frame) MutateLen(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *Frame) Sid() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frame) MutateSid(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *Frame) Inventory(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Frame) InventoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Frame) InventoryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Frame) MutateInventory(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func FrameStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func FrameAddVer(builder *flatbuffers.Builder, ver byte) {
	builder.PrependByteSlot(0, ver, 0)
}
func FrameAddCmd(builder *flatbuffers.Builder, cmd byte) {
	builder.PrependByteSlot(1, cmd, 0)
}
func FrameAddLen(builder *flatbuffers.Builder, len uint16) {
	builder.PrependUint16Slot(2, len, 0)
}
func FrameAddSid(builder *flatbuffers.Builder, sid uint32) {
	builder.PrependUint32Slot(3, sid, 0)
}
func FrameAddInventory(builder *flatbuffers.Builder, inventory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(inventory), 0)
}
func FrameStartInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FrameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
