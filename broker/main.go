package main

import (
	protos "broker/protobuffers"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type BrokerServer struct {
	protos.UnimplementedBrokerServiceServer
}

func main() {
	// inicia el servidor
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := BrokerServer{}
	protos.RegisterBrokerServiceServer(grpcServer, &server)
	startServer(listener, *grpcServer)
}

func startServer(listener net.Listener, grpcServer grpc.Server) {
	fmt.Println("Corriendo Broker...")
	grpcServer.Serve(listener)
	defer grpcServer.Stop()
}

func (s *BrokerServer) AddCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	return &protos.BrokerWriteMessage{Confirm: true, Replica: 9002}, nil
}
