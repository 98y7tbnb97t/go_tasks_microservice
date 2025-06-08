package grpc

import (
	userpb "github.com/98y7tbnb97t/GoMicro/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewUserClient создает gRPC-клиент для Users-сервиса
func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	client := userpb.NewUserServiceClient(conn)
	return client, conn, nil
}
