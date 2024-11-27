package main

import (
	"UrlScan/inflag"
	"UrlScan/util"
	"fmt"
)

func main() {
	url := inflag.GetUrl()
	u, err := util.ValidURL(&url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
