// fetchall 并发获取URL并报告它们的时间和大小
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道ch接收
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}

	f, err := os.Create(url.QueryEscape(uri))
	if err != nil {
		ch <- err.Error()
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, uri)
}
