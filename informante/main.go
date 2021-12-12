package main

import (
	"bufio"
	"context"
	"fmt"
	protos "informante/protobuffers"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

type InformantServer struct {
	replicaDirection int
	replicaVector    []int
	broker           protos.BrokerServiceClient
}

func main() {
	finish := false
	brokerPort := 50001
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", brokerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	broker := protos.NewBrokerServiceClient(conn)
	server := InformantServer{broker: broker}

	for !finish {
		fmt.Println("...")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ejecuta el comando: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		choose(text, &server)
	}
}

func choose(command string, s *InformantServer) {
	commandList := strings.Split(command, " ")
	if commandList[0] == "AddCity" {
		message := &protos.InformantMessage{
			PlanetName: commandList[1],
			CityName:   commandList[2],
			NewValue:   commandList[3][:len(commandList[3])-1],
		}
		response, _ := s.broker.AddCity(context.Background(), message)
		conn, err := grpc.Dial(response.Replica, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se logró conectar a broker: %s", err)
		}
		fulcrum := protos.NewFulcrumServiceClient(conn)
		fulcrum.AddCity(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
	} else if commandList[0] == "UpdateName" {
		message := &protos.InformantMessage{
			PlanetName: commandList[1],
			CityName:   commandList[2],
			NewValue:   commandList[3][:len(commandList[3])-1],
		}
		response, _ := s.broker.UpdateName(context.Background(), message)
		conn, err := grpc.Dial(response.Replica, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se logró conectar a broker: %s", err)
		}
		fulcrum := protos.NewFulcrumServiceClient(conn)
		fulcrum.UpdateName(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
	} else if commandList[0] == "UpdateNumber" {
		message := &protos.InformantMessage{
			PlanetName: commandList[1],
			CityName:   commandList[2],
			NewValue:   commandList[3][:len(commandList[3])-1],
		}
		response, _ := s.broker.UpdateNumber(context.Background(), message)
		conn, err := grpc.Dial(response.Replica, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se logró conectar a broker: %s", err)
		}
		fulcrum := protos.NewFulcrumServiceClient(conn)
		fulcrum.UpdateNumber(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
	} else if commandList[0] == "DeleteCity" {
		message := &protos.InformantMessage{
			PlanetName: commandList[1],
			CityName:   commandList[2][:len(commandList[2])-1],
			NewValue:   "",
		}
		response, _ := s.broker.DeleteCity(context.Background(), message)
		conn, err := grpc.Dial(response.Replica, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("No se logró conectar a broker: %s", err)
		}
		fulcrum := protos.NewFulcrumServiceClient(conn)
		fulcrum.DeleteCity(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
	} else {
		fmt.Printf("ingresar valido\n")
	}
}
