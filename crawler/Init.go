package crawler

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

type Job struct {
	url   string
	proxy Proxy
}

func Task(job Job, c chan string) {

	client := req.C().
		SetUserAgent("fiyat-ajani").
		SetTimeout(5 * time.Second)
	fmt.Println(job.proxy.proxyAdress)
	if job.proxy.proxyAdress != "" {
		client.SetProxyURL(job.proxy.proxyAdress)
	}
	resp, err := client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		EnableDump().
		Get(job.url)
	if err != nil {
		// Handle error.
		// ...
		fmt.Println(err.Error())
		c <- err.Error()
		return
	}
	if resp.IsSuccess() {
		// Handle result.
		// ...
		c <- "bitti task " + job.url
		return
	}
	if resp.IsError() {
		// Handle errMsg.
		// ...
		c <- "Hata task " + job.url
		return
	}
	// Handle unexpected response (corner case).
	//err = fmt.Errorf("got unexpected response, raw dump:\n%s", resp.Dump())
	//fmt.Println(body)

}

func chunkJob(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func finisher(taskchunk []string, ch chan string, wg *sync.WaitGroup) []string {
	defer wg.Done()
	result := []string{}
	for _, task := range taskchunk {
		result = append(result, <-ch)
		fmt.Println(task)
	}
	return result
}
func Init(urlList []string, maxConcurrent int, proxys []string) {
	var wg sync.WaitGroup
	chunkedList := chunkJob(urlList, maxConcurrent)
	proxylist := createProxyList(proxys)

	for _, taskChunk := range chunkedList {
		wg.Add(1)
		ch := make(chan string)
		for _, task := range taskChunk {
			log.Println("baslıyor task", task)
			go Task(Job{url: task, proxy: proxylist.random()}, ch)
		}
		go finisher(taskChunk, ch, &wg)
		wg.Wait()
	}
}
