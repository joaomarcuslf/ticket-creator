package grpc_client

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/joaomarcuslf/ticket-creator/proto"
)

type GrpcClient struct {
	Port string
}

type server struct{}

func NewGrpcClient(port string) *GrpcClient {
	return &GrpcClient{
		Port: port,
	}
}

func (a *GrpcClient) Initialize() {
	lis, err := net.Listen("tcp", ":"+a.Port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTicketServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Server started on port: " + a.Port)
}
