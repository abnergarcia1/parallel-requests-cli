package cmd

import (
	"testing"
)

func TestInitProcessor(t *testing.T) {
	t.Run("Should print parallel requests", func(t *testing.T) {
		urls := []string{
			"http://www.google.com",
			"http://www.adjust.com",
			"http://www.adjust.com",
			"http://www.adjust.com",
			"http://www.adjust.com",
			"http://www.adjust.com",
			"http://www.adjust.com",
		}

		maxProcess := 2
		InitProcessor(maxProcess, urls)
	})
}

func TestGetRequestMD5(t *testing.T) {
	t.Run("should return string with url and md5 sha", func(t *testing.T) {
		url := "http://www.google.com"

		response, err := GetRequestMD5(url)
		if err != nil {
			t.Errorf("method should return an url and md5 sha, but returned error:%s", err)
		}

		if response == "" {
			t.Errorf("method should return an url and md5 sha, but returned empty")
		}
	})

	t.Run("should return error if url is invalid", func(t *testing.T) {
		url := "www.google"

		response, err := GetRequestMD5(url)
		if err == nil {
			t.Errorf("method should return error, but returned no errors")
			return
		}

		if response != "" {
			t.Errorf("method should return empty string, but returned:%s", response)
			return
		}
	})
}
