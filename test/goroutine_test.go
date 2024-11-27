package test

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"sync"
	"testing"
)

// 提取链接的函数
func extractLinks(url string, wg *sync.WaitGroup, results chan<- []string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		results <- []string{}
		return
	}
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	var links []string

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			results <- links
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						links = append(links, a.Val)
					}
				}
			}
		}
	}
}
func Test_goroutine(t *testing.T) {
	urls := []string{
		"https://www.github.com",
		"https://liushuchun.gitbooks.io/golang/content/go_concurence.html",
		// 添加更多URL
	}

	var wg sync.WaitGroup
	results := make(chan []string)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			extractLinks(u, &wg, results)
		}(url)
	}

	// 独立输出每个URL的结果
	go func() {
		wg.Wait()
		close(results)
	}()
	i := 1
	for links := range results {
		fmt.Printf("第%d个链接提取到的：\n", i)
		i++
		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println("----")
	}
}
