package clients

import (
	"fmt"

	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	Client pb.AuthServiceClient
}

func InitAuthClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
