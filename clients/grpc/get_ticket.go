package grpc_client

import (
	"context"
	"fmt"

	"github.com/joaomarcuslf/ticket-creator/handlers"
	pb "github.com/joaomarcuslf/ticket-creator/proto"
	"google.golang.org/grpc/metadata"
)

func (s *server) GetTicket(ctx context.Context, in *pb.GetTicketRequest) (*pb.GetTicketResponse, error) {
	encodedUrl := in.EncodedUrl

	md, _ := metadata.FromIncomingContext(ctx)

	scheme := "http"

	if md[":scheme"] != nil && md[":scheme"][0] == "https" {
		scheme = "https"
	}

	host := md[":authority"][0]
	path := "/ticket/" + encodedUrl

	ticket, err := handlers.GetTicket(
		encodedUrl,
		scheme,
		host,
		path,
	)

	if err != nil {
		return &pb.GetTicketResponse{}, err
	} else {
		return &pb.GetTicketResponse{
			Title:       ticket.Title,
			Description: ticket.SafeDescription,
			Date:        ticket.Date,
			ShortUrl:    fmt.Sprintf("%v", ticket.ShortUrl),
		}, nil
	}
}
