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

type Planet struct {
	clockVector []int32
	address     string
}
type InformantServer struct {
	planets map[string]Planet
	broker  protos.BrokerServiceClient
}

func main() {
	finish := false
	brokerPort := 50001
	conn, err := grpc.Dial(fmt.Sprintf("10.6.40.201:%d", brokerPort), grpc.WithInsecure())//dist 61 puerto broker
	if err != nil {
		log.Fatalf("No se logró conectar a broker: %s", err)
	}
	broker := protos.NewBrokerServiceClient(conn)
	m := make(map[string]Planet)
	server := InformantServer{broker: broker, planets: m}

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
		fulcrumResponse, _ := fulcrum.AddCity(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
		fmt.Printf("vector: %v\n", fulcrumResponse.ClockValue)
		_, found := s.planets[commandList[1]]
		if !found {
			s.planets[commandList[1]] = Planet{clockVector: fulcrumResponse.ClockValue, address: response.Replica}
		}
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
		fulcrumResponse, _ := fulcrum.UpdateName(context.Background(), message)
		s.planets[commandList[1]] = Planet{clockVector: fulcrumResponse.ClockValue, address: response.Replica}
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
		fulcrumResponse, _ := fulcrum.UpdateNumber(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
		s.planets[commandList[1]] = Planet{clockVector: fulcrumResponse.ClockValue, address: response.Replica}
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
		fulcrumResponse, _ := fulcrum.DeleteCity(context.Background(), message)
		fmt.Printf("replica: %v\n", response.Replica)
		fmt.Printf("vector: %v\n", fulcrumResponse.ClockValue)
		s.planets[commandList[1]] = Planet{clockVector: fulcrumResponse.ClockValue, address: response.Replica}
	} else {
		fmt.Printf("ingresar valido\n")
	}
}
