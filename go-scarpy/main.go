package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(url string, ch chan<- string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(ioutil.Discard, res.Body)

	ch <- fmt.Sprintf("%7d  %s", res.StatusCode, url)
}
func main() {
	urls := []string{"https://google.com.tw", "https://tw.yahoo.com"}

	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
