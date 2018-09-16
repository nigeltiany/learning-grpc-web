package main

import (
	"context"
	"github.com/nigeltiany/learning-grpc-web/proto"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
)

type echoServer struct {

}

func (s *echoServer) Echo (ctx context.Context, request *EchoService.EchoRequest) (*EchoService.EchoResponse,error) {
	return &EchoService.EchoResponse{ Message: request.Message }, nil
}


func main()  {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := echoServer{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	EchoService.RegisterEchoServiceServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
