package main

import (
	"fmt"
	"github.com/Rezky08/microservices-go/infrastructure/grpc/limit/routes"
	"github.com/Rezky08/microservices-go/pb/limit"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 9301

func main() {
	netListen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen %s\n%v", netListen.Addr(), err)
	}

	grpcServer := grpc.NewServer()

	limitService := routes.Routes{}
	limit.RegisterLimitServiceServer(grpcServer, &limitService)

	log.Printf("Server started at %s", netListen.Addr())
	if err = grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
