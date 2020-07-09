//go:generate protoc ./account.proto --go_out=plugins=grpc:./pb
package account

import (
	"context"
	"fmt"
	"github.com/ashah360/microstore/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   a.ID,
		Name: a.Name,
	}}, nil
}
