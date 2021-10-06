package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestGetShortenedURL01(t *testing.T) {
	_, err := GetShortenedURL("")

	assert.True(t, err != nil)
}

func TestGetShortenedURL02(t *testing.T) {
	_, err := GetShortenedURL("joaomarcuslf.com")

	assert.True(t, err != nil)
}

func TestGetShortenedURL03(t *testing.T) {
	shortUrl, err := GetShortenedURL("https://joaomarcuslf.com/")

	assert.True(t, err == nil)
	assert.Equal(t, shortUrl.ShortUrl, "http://go-go-url-go.com/2wHg7NF5")
}
