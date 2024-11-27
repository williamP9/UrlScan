package inflag

import (
	"UrlScan/config"
	"flag"
	"fmt"
	"net/url"
)

var parsedUrl config.Url

func init() {
	inputUrl := flag.String("url", "", "Input a full URL (e.g., http://example.com:80/path)")

	flag.Parse()
	if len(*inputUrl) == 0 {
		fmt.Println("No URL provided. Use --url to pass a URL.")
		return
	}
	// 解析 URL 并绑定到 Url 对象
	u, err := url.Parse(*inputUrl)
	if err != nil {
		fmt.Printf("Invalid URL: %v\n", err)
		return
	}
	// 将解析后的值绑定到结构体
	parsedUrl = config.Url{
		Scheme: u.Scheme,
		Host:   u.Hostname(),
		Port:   u.Port(),
		Path:   u.Path,
	}

}
func GetUrl() config.Url {

	return parsedUrl
}
