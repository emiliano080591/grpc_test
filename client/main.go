package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.Println("Client running...")

	con, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	response, err := NewDepositClient(con, time.Second).Deposit(context.Background(), 1990.01)

	log.Println(response)
	log.Println(err)
}
