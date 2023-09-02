package utils

import (
	"io"
	"net/http"
)

const USERAGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

//const MOVIEURL = "https://movie.douban.com"

func Fetch(url string) ([]byte, error) {

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", USERAGENT)

	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)

}
