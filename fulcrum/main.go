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
	planets map[string]map[string]int
}

func main() {
	// inicia el servidor
	listener, err := net.Listen("tcp", ":50010")
	if err != nil {
		panic(err)
	}
	m := make(map[string]map[string]int)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := FulcrumServer{planets: m}
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
	s.planets[informanteMessage.PlanetName][informanteMessage.CityName] = int(informanteMessage.NewValue)
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateName(
	ctx context.Context,
	informanteMessage *protos.InformantStringValueMessage,
) (*protos.FulcrumWriteMessage, error) {
	s.planets[informanteMessage.PlanetName][informanteMessage.CityName] = s.planets[informanteMessage.PlanetName][informanteMessage.NewValue]
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateNumber(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {
	s.planets[informanteMessage.PlanetName][informanteMessage.CityName] = int(informanteMessage.NewValue)
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) DeleteCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {
	delete(s.planets[informanteMessage.PlanetName], informanteMessage.CityName)
	return &protos.FulcrumWriteMessage{}, nil
}
