package main

import (
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/oklog/run"
	"google.golang.org/grpc"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

func main() {
	var addr = "0.0.0.0:50051"
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

	var g run.Group

	g.Add(func() error {
		return ser.Serve(lis)
	}, func(e error) {
		ser.GracefulStop()
	})

	err = g.Run()

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
