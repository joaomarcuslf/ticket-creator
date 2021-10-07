package grpc_client

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/joaomarcuslf/ticket-creator/proto"
	"github.com/joaomarcuslf/ticket-creator/usecases"
)

func (s *server) CreateTicket(ctx context.Context, in *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	title := in.Title
	description := in.Description

	form := usecases.ValidateForm(
		title,
		description,
	)

	if form.Ok {
		encoded := "/ticket/" + usecases.EncodeTicketData(title, description)

		return &pb.CreateTicketResponse{
			RedirectUrl: &encoded,
		}, nil
	} else {
		return &pb.CreateTicketResponse{
			Errors: &pb.Errors{
				Title:       form.Title,
				Description: form.Description,
			},
			Values: &pb.Values{
				Title:       title,
				Description: description,
			},
		}, nil
	}
}

func (a *GrpcClient) CreateTicket(c *gin.Context) {
	c.Request.ParseForm()

	title := c.PostForm("title")
	description := c.PostForm("description")

	form := usecases.ValidateForm(
		title,
		description,
	)

	if form.Ok {
		encoded := usecases.EncodeTicketData(title, description)

		c.Redirect(
			http.StatusMovedPermanently,
			"/ticket/"+encoded,
		)
	} else {
		c.HTML(
			http.StatusUnprocessableEntity,
			"index.tmpl.html",
			map[string]interface{}{
				"Errors": map[string]string{
					"title":       form.Title,
					"description": form.Description,
				},
				"Values": map[string]string{
					"title":       title,
					"description": description,
				},
			},
		)
	}
}
