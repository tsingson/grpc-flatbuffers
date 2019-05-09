package model

import (
	"github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

type AddBookmark struct {
	Url    string
	Title  string
	Status bookmarks.Status
}

func (b *AddBookmark) Build() *flatbuffers.Builder {
	return builder(b.Url, b.Title, b.Status)
}

func AddRequest(urlStr, titleStr string, statusStr string) *flatbuffers.Builder {
	var sta bookmarks.Status
	sta = getStatus(statusStr)
	return builder(urlStr, titleStr, sta)
}

func builder(urlStr, titleStr string, status bookmarks.Status) *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	b.Reset()

	url := b.CreateString(urlStr)
	title := b.CreateString(titleStr)

	bookmarks.AddRequestStart(b)
	bookmarks.AddRequestAddURL(b, url)
	bookmarks.AddRequestAddTitle(b, title)

	// var sta bookmarks.Status
	// sta = getStatus(statusStr)
	bookmarks.AddRequestAddStatus(b, status)

	b.Finish(bookmarks.AddRequestEnd(b))

	return b
}

func getStatus(stStr string) (st bookmarks.Status) {
	if len(stStr) > 0 {
		switch stStr {
		case "Online":
			st = bookmarks.StatusOnline
		case "Offline":
			st = bookmarks.StatusOffline
		default:
			st = bookmarks.StatusUnAccessAble
		}
	} else {
		st = bookmarks.StatusUnAccessAble
	}
	return
}
