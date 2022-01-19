package cmd

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// GetRequestMD5 request to the url and returns the url and response md5 sha
func GetRequestMD5(url string) (string, error) {

	hasHttpPrefix := strings.Contains(url, "http")
	if !hasHttpPrefix {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %x", url, md5.Sum(bodyResponse)), nil
}

// InitProcessor Init concurrent processor
func InitProcessor(maxProcess int, urls []string) {
	var wg sync.WaitGroup
	var batches, urlIndex int

	batches = len(urls) / maxProcess
	if len(urls)%maxProcess > 0 && batches == 0 {
		batches = 1
	}

	for i := 1; i <= batches; i++ {
		for j := 0; j < maxProcess; j++ {
			if urlIndex >= len(urls) {
				break
			}

			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				message, err := GetRequestMD5(urls[index])
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(message)
			}(urlIndex)

			urlIndex++
		}
		wg.Wait()
		time.Sleep(time.Second * 1)
	}
}
