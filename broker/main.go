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
	fulcrumPort := 50010
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", fulcrumPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	fmt.Println("Añadiendo ciudad...")
	fulcrumService := protos.NewFulcrumServiceClient(conn)
	_, errorr := fulcrumService.AddCity(context.Background(), informanteMessage)
	if errorr != nil {
		log.Fatalf("Error con fulcrum no se pudo añadir ciudad")
	}
	return &protos.BrokerWriteMessage{Confirm: true, Replica: 9002}, nil
}

func (s *BrokerServer) UpdateName(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	fulcrumPort := 50010
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", fulcrumPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	fmt.Println("Cambiando nombre a ciudad...")
	fulcrumService := protos.NewFulcrumServiceClient(conn)
	_, errorr := fulcrumService.UpdateName(context.Background(), informanteMessage)
	if errorr != nil {
		log.Println("Error con fulcrum no se pudo actualizar nombre")
	}
	return &protos.BrokerWriteMessage{Confirm: true, Replica: 9002}, nil
}

func (s *BrokerServer) UpdateNumber(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	fulcrumPort := 50010
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", fulcrumPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	fmt.Println("Actualizando espias en ciudad...")
	fulcrumService := protos.NewFulcrumServiceClient(conn)
	_, errorr := fulcrumService.UpdateNumber(context.Background(), informanteMessage)
	if errorr != nil {
		log.Println("Error con fulcrum no se pudo actualizar espias")
	}
	return &protos.BrokerWriteMessage{Confirm: true, Replica: 9002}, nil
}

func (s *BrokerServer) DeleteCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.BrokerWriteMessage, error) {
	fulcrumPort := 50010
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", fulcrumPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	fmt.Println("Eliminando ciudad")
	fulcrumService := protos.NewFulcrumServiceClient(conn)
	_, errorr := fulcrumService.UpdateNumber(context.Background(), informanteMessage)
	if errorr != nil {
		log.Println("Error con fulcrum no se pudo eliminar ciudad")
	}
	return &protos.BrokerWriteMessage{Confirm: true, Replica: 9002}, nil
}
