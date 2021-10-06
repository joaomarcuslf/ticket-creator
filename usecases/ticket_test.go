package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestExtractTicketValues01(t *testing.T) {
	_, err := ExtractTicketValues("")

	assert.True(t, err != nil)
}

func TestExtractTicketValues02(t *testing.T) {
	_, err := ExtractTicketValues("dGVzdGUtLXwtLXRlc3Q=")

	assert.True(t, err != nil)
}

func TestExtractTicketValues03(t *testing.T) {
	ticket, err := ExtractTicketValues("dGVzdC0tfC0tdGVzdC0tfC0tMjAyMS0xMC0wNg==")

	assert.True(t, err == nil)
	assert.Equal(t, ticket.Title, "test")
	assert.Equal(t, ticket.Description, []string{"test"})
	assert.Equal(t, ticket.SafeDescription, "test")
	assert.Equal(t, ticket.Date, "2021-10-06")
}

func TestEncodeTicketData01(t *testing.T) {
	encoded := EncodeTicketData("test", "")

	assert.True(t, encoded != "")
}
