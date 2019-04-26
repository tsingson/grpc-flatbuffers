package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc/codes"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type book struct {
	id        int
	lastTitle string
	lastURL   string
	Status    int8
}

type server struct {
	id    int
	books map[int]*book
}

var addr = "0.0.0.0:50051"

func (s *server) Add(context context.Context, in *bookmarks.AddRequest) (*flatbuffers.Builder, error) {
	log.Println("Add called...")

	s.id++
	b := &book{}
	b.id = s.id
	b.lastTitle = string(in.Title())
	b.lastURL = string(in.URL())
	b.Status = in.Status()
	s.books[s.id] = b

	out := flatbuffers.NewBuilder(0)
	bookmarks.AddResponseStart(out)
	out.Finish(bookmarks.AddResponseEnd(out))
	return out, nil
}

func (s *server) LastAdded(context context.Context, in *bookmarks.LastAddedRequest) (*flatbuffers.Builder, error) {
	log.Println("LastAdded called...")

	b := flatbuffers.NewBuilder(0)
	id := b.CreateString(strconv.Itoa(s.id))
	_, ok := s.books[s.id]
	if ok {
		title := b.CreateString(s.books[s.id].lastTitle)
		url := b.CreateString(s.books[s.id].lastURL)
		sta := s.books[s.id].Status

		bookmarks.LastAddedResponseStart(b)
		bookmarks.LastAddedResponseAddID(b, id)
		bookmarks.LastAddedResponseAddTitle(b, title)
		bookmarks.LastAddedResponseAddURL(b, url)
		bookmarks.LastAddedResponseAddStatus(b, sta)
		b.Finish(bookmarks.LastAddedResponseEnd(b))
		return b, nil
	}
	err := status.Error(codes.NotFound, "id was not found")
	return nil, err
}

func (s *server) All(in *bookmarks.LastAddedRequest, serv bookmarks.BookmarksService_AllServer) error {
	log.Println("All called...")

	if s.id > 0 {

		for i := 0; i <= s.id+1; i++ {
			k, ok := s.books[i]
			if ok {
				b := flatbuffers.NewBuilder(0)
				id := b.CreateString(strconv.Itoa(k.id))
				title := b.CreateString(k.lastTitle)
				url := b.CreateString(k.lastURL)
				sta := k.Status

				bookmarks.LastAddedResponseStart(b)
				bookmarks.LastAddedResponseAddID(b, id)
				bookmarks.LastAddedResponseAddTitle(b, title)
				bookmarks.LastAddedResponseAddURL(b, url)
				bookmarks.LastAddedResponseAddStatus(b, sta)
				b.Finish(bookmarks.LastAddedResponseEnd(b))
				_ = serv.Send(b)
			}

		}
		return nil
	}
	err := status.Error(codes.NotFound, "id was ------------> not found")
	return err

}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ser := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))

	serv := &server{
		id:    0,
		books: make(map[int]*book, 1),
	}

	bookmarks.RegisterBookmarksServiceServer(ser, serv)
	if err := ser.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
