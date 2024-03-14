package auth

import (
	"fmt"

	"github.com/rahulchacko7/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/rahulchacko7/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Coult not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)

}
