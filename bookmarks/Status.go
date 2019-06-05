// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bookmarks

import "strconv"

type Status byte

const (
	StatusOnline       Status = 1
	StatusOffline      Status = 2
	StatusUnAccessAble Status = 4
)

var EnumNamesStatus = map[Status]string{
	StatusOnline:       "Online",
	StatusOffline:      "Offline",
	StatusUnAccessAble: "UnAccessAble",
}

var EnumValuesStatus = map[string]Status{
	"Online":       StatusOnline,
	"Offline":      StatusOffline,
	"UnAccessAble": StatusUnAccessAble,
}

func (v Status) String() string {
	if s, ok := EnumNamesStatus[v]; ok {
		return s
	}
	return "Status(" + strconv.FormatInt(int64(v), 10) + ")"
}
