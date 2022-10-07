package main

import (
	pb "github.com/emiliano080591/grpc_test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Server running...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	pb.RegisterDepositServiceServer(server, &DepositServer{})

	log.Fatalln(server.Serve(listener))
}
