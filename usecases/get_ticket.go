package usecases

import (
	"fmt"
)

type TicketData struct {
	Title           string      `json:"title"`
	Description     []string    `json:"description"`
	SafeDescription string      `json:"description_safe"`
	Date            string      `json:"date"`
	ShortUrl        interface{} `json:"short_url"`
}

func GetTicket(encodedUrl string, scheme string, host string, path string) (TicketData, error) {
	ticket, err := ExtractTicketValues(encodedUrl)

	if err != nil {
		return TicketData{}, fmt.Errorf("Error extracting ticket values: %s", err)
	}

	shortUrl, err := GetShortenedURL(scheme + "://" + host + path)

	if err != nil {
		return TicketData{}, fmt.Errorf("Error getting shortened url: %s", err)
	}

	return TicketData{
		Title:           ticket.Title,
		Description:     ticket.Description,
		SafeDescription: ticket.SafeDescription,
		Date:            ticket.Date,
		ShortUrl:        shortUrl.ShortUrl,
	}, nil
}
