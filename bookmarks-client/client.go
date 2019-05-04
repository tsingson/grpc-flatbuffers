package main

import (
	"context"
	"log"
	"os"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"

	"google.golang.org/grpc"
)

type server struct{}

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
	}

}

func clientAll(client bookmarks.BookmarksServiceClient) (err error) {
	b := flatbuffers.NewBuilder(0)
	bookmarks.LastAddedRequestStart(b)
	b.Finish(bookmarks.LastAddedRequestEnd(b))
	out, err := client.All(context.Background(), b)
	if err != nil {
		log.Fatalf("Retrieve client failed: %v", err)
	}
	for {
		out, err := out.Recv()
		if err == nil {
			log.Println("ID: ", string(out.ID()))
			log.Println("URL: ", string(out.URL()))
			log.Println("Title: ", string(out.Title()))
			log.Println("Status: ", bookmarks.EnumNamesStatus[out.Status()])
			log.Println("LastTime",   time.Unix(out.LastTimes(), 0).Format("2006-01-02 15:04:05"))
			log.Println("---------------------------")
			log.Println(" ")
		} else {
			break
		}

	}
	log.Println("Done")
	return
}

func clientLastAdd(client bookmarks.BookmarksServiceClient) (err error) {
	b := flatbuffers.NewBuilder(0)
	bookmarks.LastAddedRequestStart(b)
	b.Finish(bookmarks.LastAddedRequestEnd(b))
	out, err := client.LastAdded(context.Background(), b)
	if err != nil {
		log.Fatalf("Retrieve client failed: %v", err)
		return
	}
	log.Println("ID: ", string(out.ID()))
	log.Println("URL: ", string(out.URL()))
	log.Println("Title: ", string(out.Title()))
	log.Println("Status: ", bookmarks.EnumNamesStatus[out.Status()])
	// log.Println("LastTime",out.LastTimes())
	log.Println("LastTime",   time.Unix(out.LastTimes(), 0).Format("2006-01-02 15:04:05"))
	return
}

func clientAdd(client bookmarks.BookmarksServiceClient) (err error) {

	if len(os.Args) < 4 {
		log.Fatalln("Insufficient args provided for add command..")
	}
	b := flatbuffers.NewBuilder(0)
	url := b.CreateString(os.Args[2])
	title := b.CreateString(os.Args[3])
	bookmarks.AddRequestStart(b)
	bookmarks.AddRequestAddURL(b, url)
	bookmarks.AddRequestAddTitle(b, title)

	var sta int8
	if len(os.Args[4]) > 0 {
		switch os.Args[4] {
		case "Online":
			sta = bookmarks.StatusOnline
		case "Offline":
			sta = bookmarks.StatusOffline
		default:
			sta = bookmarks.StatusUnAccessAble
		}

		bookmarks.AddRequestAddStatus(b, sta)
	}
	b.Finish(bookmarks.AddRequestEnd(b))
	_, err = client.Add(context.Background(), b)
	if err != nil {
		log.Fatalf("Retrieve client failed: %v", err)
	}
	return err
}
