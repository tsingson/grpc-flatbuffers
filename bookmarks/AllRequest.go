// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bookmarks

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AllRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsAllRequest(buf []byte, offset flatbuffers.UOffsetT) *AllRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AllRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AllRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AllRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func AllRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func AllRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
