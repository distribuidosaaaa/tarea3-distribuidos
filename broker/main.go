package main

import (
	protos "broker/protobuffers"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BrokerServer struct {
	protos.UnimplementedBrokerServiceServer
	replicas []string
}

func main() {
	// inicia el servidor
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	replicas := []string{
		"localhost:50010",
	}
	server := BrokerServer{replicas: replicas}
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
	return &protos.BrokerWriteMessage{Replica: s.replicas[0]}, nil
}

func (s *BrokerServer) UpdateName(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	return &protos.BrokerWriteMessage{Replica: s.replicas[0]}, nil
}

func (s *BrokerServer) UpdateNumber(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	return &protos.BrokerWriteMessage{Replica: s.replicas[0]}, nil
}

func (s *BrokerServer) DeleteCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	return &protos.BrokerWriteMessage{Replica: s.replicas[0]}, nil
}

func (s *BrokerServer) GetRebelds(
	ctx context.Context,
	leiaMessage *protos.LeiaMessage,
) (*protos.BrokerReadMessage, error) {
	conn, err := grpc.Dial(s.replicas[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a replica: %s", err)
	}
	replica := protos.NewFulcrumServiceClient(conn)
	response, _ := replica.GetRebelds(context.Background(), leiaMessage)
	return &protos.BrokerReadMessage{Address: s.replicas[0], Spies: response.Spies, ClockValue: response.ClockValue}, nil
}
