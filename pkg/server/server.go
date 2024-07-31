package server

import (
	"log"
	"net"

	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/handler"
	pb "github.com/ratheeshkumar/microservice_grpc_userV1/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handl *handler.UserHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("error creating listener on port 8082")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handl)
	log.Printf("lisitening on gRPC Server 8082")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")
	}

}
