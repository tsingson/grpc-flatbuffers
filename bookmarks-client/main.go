package main

import (
	"log"
	"os"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

var addr = "0.0.0.0:50051"

func main() {

	if len(os.Args) < 2 {
		log.Fatalln("Insufficient args provided")
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := bookmarks.NewBookmarksServiceClient(conn)

	cmd := os.Args[1]

	if cmd == "add" {

		_ = clientAdd(client)

	} else if cmd == "last-added" {

		_ = clientLastAdd(client)

	} else if cmd == "all" {

		_ = clientAll(client)
	} else if cmd == "getall" {
		_ = clientGetAll(client)
	}

}
