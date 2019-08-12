package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sanity-io/litter"

	"github.com/tsingson/grpc-flatbuffers/frame"
)

type FrameObj struct {
	Ver  byte   // version
	Cmd  byte   // command
	Len  uint16 // Data (payload) length
	Sid  uint32 // stream ID
	Data []byte
}

func (f *FrameObj) Marshell(b *flatbuffers.Builder) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	// 这里需要注意的是，由于是 PrependByte，前置字节，所以循环的时候需要反向迭代
	frame.FrameAddInventory(b, 10)
	// Note: Since we prepend the bytes, this loop iterates in reverse.
	for i := 9; i >= 0; i-- {
		b.PrependByte(byte(i))
	}
	inv := b.EndVector(10)

	frame.FrameStart(b)

	frame.FrameAddVer(b, f.Ver)
	frame.FrameAddCmd(b, f.Cmd)
	frame.FrameAddLen(b, f.Len)
	frame.FrameAddSid(b, f.Sid)
	frame.FrameAddInventory(b, inv)
	// frame_position := frame.FrameEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(frame.FrameEnd(b))
	// return the byte slice containing encoded Data:
	return b.Bytes[b.Head():]
}

func UnMarshell(buf []byte) (fo *FrameObj, err error) {
	fo = &FrameObj{}

	f := frame.GetRootAsFrame(buf, 0)
	// if f.Ver() != byte(1) {
	// 	err = xerrors.New("parse error")
	// 	return
	// }
	fo.Ver = f.Ver()
	fo.Cmd = f.Cmd()
	fo.Len = f.Len()
	fo.Sid = f.Sid()
	// fo.Data = f.DataBytes()

	return
}

func main() {
	f := &FrameObj{
		Ver:  byte(1),
		Cmd:  byte(2),
		Len:  uint16(32),
		Sid:  uint32(32),
		Data: []byte("testtesttesttest"),
	}
	b := flatbuffers.NewBuilder(0)
	buf := f.Marshell(b)
	litter.Dump(string(buf))

	ff, err := UnMarshell(buf)
	if err != nil {

	}
	litter.Dump(ff)

}
