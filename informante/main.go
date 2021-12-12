package main

import (
	"bufio"
	"context"
	"fmt"
	protos "informante/protobuffers"
	"log"
	"os"
	"strconv"
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
		value, _ := strconv.ParseInt(commandList[3], 10, 8)
		response, _ := s.broker.AddCity(
			context.Background(),
			&protos.InformantMessage{
				PlanetName: commandList[1],
				CityName:   commandList[2],
				NewValue:   int32(value),
			},
		)
		fmt.Printf("replica: %d\n", response.Replica)
	} else if commandList[0] == "UpdateName" {
		response, _ := s.broker.UpdateName(
			context.Background(),
			&protos.InformantStringValueMessage{
				PlanetName: commandList[1],
				CityName:   commandList[2],
				NewValue:   commandList[3],
			},
		)
		fmt.Printf("replica: %d\n", response.Replica)
	} else if commandList[0] == "UpdateNumber" {
		value, _ := strconv.ParseInt(commandList[3], 10, 8)
		response, _ := s.broker.UpdateNumber(
			context.Background(),
			&protos.InformantMessage{
				PlanetName: commandList[1],
				CityName:   commandList[2],
				NewValue:   int32(value),
			},
		)
		fmt.Printf("replica: %d\n", response.Replica)
	} else if commandList[0] == "DeleteCity" {
		response, _ := s.broker.DeleteCity(
			context.Background(),
			&protos.InformantMessage{
				PlanetName: commandList[1],
				CityName:   commandList[2],
				NewValue:   0,
			},
		)
		fmt.Printf("replica: %d\n", response.Replica)
	} else {
		fmt.Printf("ingresar valido\n")
	}
}
