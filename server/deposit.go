package main

import (
	"context"
	pb "github.com/emiliano080591/grpc_test/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type DepositServer struct {
	pb.UnimplementedDepositServiceServer
}

func (*DepositServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	log.Println(req.GetAmount())

	if req.GetAmount() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot deposit %v", req.GetAmount())
	}
	return &pb.DepositResponse{
		Ok: true,
	}, nil
}
