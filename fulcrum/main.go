package main

import (
	"context"
	"fmt"
	protos "fulcrum/protobuffers"
	"net"

	"google.golang.org/grpc"
)

type FulcrumServer struct {
	protos.UnimplementedFulcrumServiceServer
}

func main() {
	// inicia el servidor
	listener, err := net.Listen("tcp", ":50010")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := FulcrumServer{}
	protos.RegisterFulcrumServiceServer(grpcServer, &server)
	startServer(listener, *grpcServer)
}

func startServer(listener net.Listener, grpcServer grpc.Server) {
	fmt.Println("Corriendo fulcrum...")
	grpcServer.Serve(listener)
	defer grpcServer.Stop()
}

func (s *FulcrumServer) AddCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {

	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateName(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {

	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateNumber(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {

	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) DeleteCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {

	return &protos.FulcrumWriteMessage{}, nil
}
