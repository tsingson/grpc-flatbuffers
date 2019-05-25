package main

import (
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ser := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))

	serv := &server{
		id:    int32(0),
		books: make(map[int32]*book, 1),
	}

	bookmarks.RegisterBookmarksServiceServer(ser, serv)
	if err := ser.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
