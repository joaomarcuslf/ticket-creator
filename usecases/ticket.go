package usecases

import (
	"fmt"
	"strings"
	"time"

	"github.com/joaomarcuslf/ticket-creator/encoders"
)

var (
	TIME_FORMAT = "2006-01-02"
	DIVIDER     = "--|--"
)

type TicketValues struct {
	Title           string
	Description     []string
	SafeDescription string
	Date            string
}

func EncodeTicketData(title string, description string) string {
	now := time.Now()

	encoded := encoders.Encode(
		fmt.Sprintf(
			"%v%s%v%s%v",
			title,
			DIVIDER,
			description,
			DIVIDER,
			now.Format(TIME_FORMAT),
		),
	)

	return encoded
}

func ExtractTicketValues(encodedUrl string) (TicketValues, error) {
	decodedUrl, err := encoders.Decode(encodedUrl)

	if err != nil {
		return TicketValues{}, err
	}

	s := strings.Split(decodedUrl, DIVIDER)

	if len(s) != 3 {
		return TicketValues{}, fmt.Errorf("Invalid ticket url")
	}

	description := strings.Split(s[1], "\r\n")

	return TicketValues{
		Title:           s[0],
		Description:     description,
		SafeDescription: s[1],
		Date:            s[2],
	}, nil
}
