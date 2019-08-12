package main

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

func (s *server) buildAllResponse() (all *flatbuffers.Builder) {

	// 初始化 builder
	all = flatbuffers.NewBuilder(0)

	var offset flatbuffers.UOffsetT
	var data = make(map[int]flatbuffers.UOffsetT, 0)

	var count int
	if s.id > 0 {
		// 倒序处理 books
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
