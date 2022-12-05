package main

import (
	"github.com/luuwa/bad-ass-crawler/crawler"
)

func main() {
	tsk := []string{}
	proxies := []string{}
	for i := 1; i <= 10; i++ {
		tsk = append(tsk, "https://www.trendyol.com/riot-games/2850-vp-valorant-points-tr-p-40275298")
	}

	crawler.Init(tsk, 2, proxies)
}
