package clients

import (
	"fmt"

	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type ProductClient struct {
	Client pb.ProductServiceClient
}

func InitProductClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}
