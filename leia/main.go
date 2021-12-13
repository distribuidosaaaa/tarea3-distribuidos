package main

import (
	"bufio"
	"context"
	"fmt"
	protos "leia/protobuffers"
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
	conn, err := grpc.Dial(fmt.Sprintf("10.6.40.201:%d", brokerPort), grpc.WithInsecure())//dist 61
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
	if commandList[0] == "GetNumberRebelds" {
		message := &protos.LeiaMessage{
			PlanetName: commandList[1],
			CityName:   commandList[2][:len(commandList[2])-1],
		}
		response, _ := s.broker.GetRebelds(context.Background(), message)
		// guardo vector
		s.planets[commandList[1]] = Planet{clockVector: response.ClockValue, address: response.Address}
		fmt.Printf("espias: %v \n", response.Spies)
	} else {
		fmt.Println("Ingrese comando valido")
	}
}
