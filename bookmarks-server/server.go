package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type book struct {
	id        int32
	lastTitle string
	lastURL   string
	Status    int8
	LastTime  int64
}

type server struct {
	id    int32
	books map[int32]*book
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
	b.LastTime = time.Now().Unix()
	s.books[s.id] = b

	out := flatbuffers.NewBuilder(0)
	bookmarks.AddResponseStart(out)
	out.Finish(bookmarks.AddResponseEnd(out))
	return out, nil
}

func (s *server) LastAdded(context context.Context, in *bookmarks.LastAddedRequest) (*flatbuffers.Builder, error) {
	log.Println("LastAdded called...")

	b := flatbuffers.NewBuilder(0)
	id := b.CreateString(strconv.Itoa(int(s.id)))
	_, ok := s.books[s.id]
	if ok {
		title := b.CreateString(s.books[s.id].lastTitle)
		url := b.CreateString(s.books[s.id].lastURL)
		sta := s.books[s.id].Status
		lst := s.books[s.id].LastTime

		bookmarks.LastAddedResponseStart(b)
		bookmarks.LastAddedResponseAddID(b, id)
		bookmarks.LastAddedResponseAddTitle(b, title)
		bookmarks.LastAddedResponseAddURL(b, url)
		bookmarks.LastAddedResponseAddStatus(b, sta)
		bookmarks.LastAddedResponseAddLastTimes(b, lst)
		b.Finish(bookmarks.LastAddedResponseEnd(b))
		return b, nil
	}
	err := status.Error(codes.NotFound, "id was not found")
	return nil, err
}

func (s *server) All(in *bookmarks.LastAddedRequest, serv bookmarks.BookmarksService_AllServer) error {
	log.Println("All called...")

	if s.id > 0 {

		for i := 0; i <= int(s.id+1); i++ {
			k, ok := s.books[int32(i)]
			if ok {
				b := flatbuffers.NewBuilder(0)
				id := b.CreateString(strconv.Itoa(int(k.id)))
				title := b.CreateString(k.lastTitle)
				url := b.CreateString(k.lastURL)
				sta := k.Status
				lst := s.books[s.id].LastTime

				bookmarks.LastAddedResponseStart(b)
				bookmarks.LastAddedResponseAddID(b, id)
				bookmarks.LastAddedResponseAddTitle(b, title)
				bookmarks.LastAddedResponseAddURL(b, url)
				bookmarks.LastAddedResponseAddStatus(b, sta)
				bookmarks.LastAddedResponseAddLastTimes(b, lst)
				b.Finish(bookmarks.LastAddedResponseEnd(b))
				_ = serv.Send(b)
			}

		}
		return nil
	}
	err := status.Error(codes.NotFound, "id was ------------> not found")
	return err

}
func (s *server) GetAll(context context.Context, in *bookmarks.AllRequest) (all *flatbuffers.Builder, err error) {
	log.Println("getAll called...")
	all = s.buildAllResponse()
	return

}

func (s *server) buildAllResponse() (all *flatbuffers.Builder) {

	// 初始化 builder
	all = flatbuffers.NewBuilder(0)

	var offset flatbuffers.UOffsetT
	var data = make(map[int]flatbuffers.UOffsetT, 0)

	var count int
	if s.id > 0 {

		for i := int(s.id + 1); i >= 0; i-- {
			k, ok := s.books[int32(i)]
			if ok {
				id := all.CreateString(strconv.Itoa(int(k.id)))
				title := all.CreateString(k.lastTitle)
				url := all.CreateString(k.lastURL)
				sta := k.Status
				lst := s.books[s.id].LastTime

				bookmarks.LastAddedResponseStart(all)
				bookmarks.LastAddedResponseAddID(all, id)
				bookmarks.LastAddedResponseAddTitle(all, title)
				bookmarks.LastAddedResponseAddURL(all, url)
				bookmarks.LastAddedResponseAddStatus(all, sta)
				bookmarks.LastAddedResponseAddLastTimes(all, lst)
				off := bookmarks.LastAddedResponseEnd(all)
				data[count] = off
				count++
			}
		}
		bookmarks.AllResponseStartDataVector(all, count)
		for j := 0; j < count; j++ {
			all.PrependUOffsetT(data[j])
		}
		offset = all.EndVector(count)
	}
	bookmarks.AllResponseStart(all)
	bookmarks.AllResponseAddTotal(all, s.id)
	bookmarks.AllResponseAddData(all, offset)
	all.Finish(bookmarks.AllResponseEnd(all))
	return
}
