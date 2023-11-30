package main

import (
	"fmt"
	"time"
)

func main() {
	// Lấy thời điểm hiện tại
	now := time.Now()
	otherTime := now.Add(60 * time.Second)
	nowTimestamp := now.Unix()
	otherTimestamp := otherTime.Unix()
	if otherTimestamp-nowTimestamp <= 60 {
		fmt.Println("Hai thời điểm cách nhau chính xác 60 giây.", "*")
	} else {
		fmt.Println("Hai thời điểm không cách nhau đúng 60 giây.", "#")
	}
}
