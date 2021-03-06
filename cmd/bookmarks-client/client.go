package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sanity-io/litter"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
	"github.com/tsingson/grpc-flatbuffers/model"
)

type server struct{}

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
			log.Println("LastTime", time.Unix(out.LastTimes(), 0).Format("2006-01-02 15:04:05"))
			log.Println("---------------------------")
			log.Println(" ")
		} else {
			break
		}

	}
	log.Println("Done")
	return
}

func clientGetAll(client bookmarks.BookmarksServiceClient) (err error) {
	b := flatbuffers.NewBuilder(0)
	bookmarks.AllRequestStart(b)
	b.Finish(bookmarks.AllRequestEnd(b))
	out, err := client.GetAll(context.Background(), b)
	if err != nil {
		log.Fatalf("Retrieve client failed: %v", err)
	}
	if out.Total() > 0 {

		litter.Dump(out.DataLength())
		for i := 0; i < int(out.Total()); i++ {

			obj := &bookmarks.LastAddedResponse{}

			if out.Data(obj, i) {
				fmt.Println(i)
				log.Println("ID: ", string(obj.ID()))
				log.Println("URL: ", string(obj.URL()))
				log.Println("Title: ", string(obj.Title()))
				log.Println("Status: ", bookmarks.EnumNamesStatus[obj.Status()])
				// log.Println("LastTime",out.LastTimes())
				log.Println("LastTime", time.Unix(obj.LastTimes(), 0).Format("2006-01-02 15:04:05"))
			}
		}
	}

	fmt.Println("")
	fmt.Println("call server Done")
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
	log.Println("LastTime", time.Unix(out.LastTimes(), 0).Format("2006-01-02 15:04:05"))
	return
}

func clientAdd(client bookmarks.BookmarksServiceClient) (err error) {
	if len(os.Args) < 4 {
		log.Fatalln("Insufficient args provided for add command..")
	}
	urlStr := os.Args[2]
	titleStr := os.Args[3]
	statusStr := os.Args[4]

	b := model.AddRequest(urlStr, titleStr, statusStr)

	_, err = client.Add(context.Background(), b)
	if err != nil {
		log.Fatalf("Retrieve client failed: %v", err)
	}
	return err
}
