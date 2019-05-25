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

	client := bookmarks.NewBookmarksServiceClient(conn)

	cmd := os.Args[1]
	if len(cmd) > 0 {
		switch cmd {
		case "add":
			_ = clientAdd(client)

		case "last-added":

			_ = clientLastAdd(client)

		case "all":
			_ = clientAll(client)

		case "getall":
			_ = clientGetAll(client)

		default:
			_ = clientLastAdd(client)
		}
	}
	_ = conn.Close()

}
