package rpcconnect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(url string, interceptor grpc.UnaryClientInterceptor) (*grpc.ClientConn, error) {
	backoffConfig := backoff.DefaultConfig

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoffConfig,
		}),
	}
	if interceptor != nil {
		dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(
			interceptor,
		))
	}

	return grpc.Dial(
		url,
		dialOpts...,
	)
}
