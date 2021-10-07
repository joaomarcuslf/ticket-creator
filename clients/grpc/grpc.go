package grpc_client

type GrpcClient struct {
	Port string
}

func NewGrpcClient(port string) *GrpcClient {
	return &GrpcClient{
		Port: port,
	}
}

func (a *GrpcClient) Initialize() {}
