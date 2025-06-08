package grpc

import (
	"net"

	taskpb "github.com/98y7tbnb97t/GoMicro/proto/taskpb"
	userpb "github.com/98y7tbnb97t/GoMicro/proto/userpb"
	"github.com/98y7tbnb97t/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)
	return grpcSrv.Serve(lis)
}
