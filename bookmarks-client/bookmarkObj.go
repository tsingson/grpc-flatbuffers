package main

import (
	"github.com/google/flatbuffers/go"

	"github.com/tsingson/grpc-flatbuffers/bookmarks"
)

type bookmarkObj struct {
	Url    string
	Title  string
	Status bookmarks.Status
}

func (b *bookmarkObj) Build() *flatbuffers.Builder {
	return BookmarkBuilder(b.Url, b.Title, b.Status)
}

func BookmarkBuilder(urlStr, titleStr string, status bookmarks.Status) *flatbuffers.Builder {
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

func bookmarkBuilder(urlStr, titleStr string, statusStr string) *flatbuffers.Builder {
	var sta bookmarks.Status
	sta = getStatus(statusStr)
	return BookmarkBuilder(urlStr, titleStr, sta)
}

func getStatus(statusStr string) (sta bookmarks.Status) {
	if len(statusStr) > 0 {
		switch statusStr {
		case "Online":
			sta = bookmarks.StatusOnline
		case "Offline":
			sta = bookmarks.StatusOffline
		default:
			sta = bookmarks.StatusUnAccessAble
		}
	} else {
		sta = bookmarks.StatusUnAccessAble
	}
	return
}
