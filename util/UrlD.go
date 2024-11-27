package util

import (
	"UrlScan/config"
	"fmt"
	"net/url"
)

func ValidURL(Url *config.Url) (string, error) {
	u := fmt.Sprintf("%s://%s:%s", Url.Scheme, Url.Host, Url.Port)

	// 解析 URL
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	// 校验 Scheme,仅支持http,https协议
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("invalid scheme: %s", parsedURL.Scheme)
	}

	// 校验 Host
	if parsedURL.Host == "" {
		return "", fmt.Errorf("invalid host: %s", parsedURL.Host)
	}

	return parsedURL.String(), nil
}
