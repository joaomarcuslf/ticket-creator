package handlers

import (
	"testing"

	"github.com/joaomarcuslf/ticket-creator/usecases"
	"github.com/stretchr/testify/assert"
)

func TestGetTicket01(t *testing.T) {
	_, err := GetTicket("aaa", "https", "localhost", "8080")

	assert.True(t, err != nil)
}

func TestGetTicket02(t *testing.T) {
	ticket, err := GetTicket("dGVzdC0tfC0tdGVzdC0tfC0tMjAyMS0xMC0wNg==", "https", "localhost", "8080")

	assert.True(t, err == nil)
	assert.Equal(t, ticket.Title, "test")
	assert.Equal(t, ticket.Description, []string{"test"})
	assert.Equal(t, ticket.SafeDescription, "test")
	assert.Equal(t, ticket.Date, "2021-10-06")
	assert.Equal(t, ticket.ShortUrl, usecases.ShortUrl(usecases.ShortUrl{ShortUrl: "http://go-go-url-go.com/FGQiKTSj"}))
}
