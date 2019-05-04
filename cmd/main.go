package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.UnixNano())
	var tt int64
	tt = t.Unix()
	// t1 := time.Unix(tt, 0)
	fmt.Println(time.Unix(tt, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(tt )
}
