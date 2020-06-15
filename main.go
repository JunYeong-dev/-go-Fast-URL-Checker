package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	var c = make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.apple.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// channel의 directions 기본적으로 channel은 양방향이지만 매개변수의 방향표시를 통해 읽기, 쓰기 전용으로 만들 수도 있음
// chan<- : 데이터를 channel에 보낼 수만 있음
// <-chan : channel로 부터 데이터를 가져올 수만 있음
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "SUCCESS"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
