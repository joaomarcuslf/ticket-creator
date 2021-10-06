package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ShortUrl struct {
	ShortUrl string
}

func GetShortenedURL(url string) (ShortUrl, error) {
	if url == "" {
		return ShortUrl{}, fmt.Errorf("URL is empty")
	}

	if !strings.Contains(url, "http") || !strings.Contains(url, "://") {
		return ShortUrl{}, fmt.Errorf("URL is invalid")
	}

	requestBody := strings.NewReader(`
	{
		"long_url": "` + url + `"
	}
	`)

	// post some data
	res, err := http.Post(
		"https://vast-thicket-55540.herokuapp.com/create-short-url",
		"application/json; charset=UTF-8",
		requestBody,
	)

	if err != nil {
		return ShortUrl{}, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return ShortUrl{}, err
	}

	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		return ShortUrl{}, err
	}

	return ShortUrl{
		ShortUrl: raw["short_url"].(string),
	}, nil
}
