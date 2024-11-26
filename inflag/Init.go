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
	//// 动态生成 URL
	//u := parsedUrl
	//
	//// 拼接 URL，确保只有在端口存在时才添加 ":port"
	//var urlString string
	//if u.Port != "" {
	//	urlString = fmt.Sprintf("%s://%s:%s", u.Scheme, u.Host, u.Port)
	//} else {
	//	urlString = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	//}
	//
	//// 如果路径存在，拼接路径
	//if u.Path != "" {
	//	urlString += u.Path
	//}
	//
	//return urlString
	return parsedUrl
}
