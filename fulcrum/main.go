package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	protos "fulcrum/protobuffers"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

type FulcrumServer struct {
	protos.UnimplementedFulcrumServiceServer
	planets        map[string]map[string]string
	planetVersions map[string]*[3]int
	updateOn       int
}

func main() {
	// inicia el servidor
	listener, err := net.Listen("tcp", ":50010")
	if err != nil {
		panic(err)
	}
	m := make(map[string]map[string]string)
	v := make(map[string]*[3]int)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := FulcrumServer{planets: m, planetVersions: v, updateOn: 1}
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
	fmt.Println("Aadiendo ciudad...")
	// verificamos si ya existe
	_, found := s.planetVersions[informanteMessage.PlanetName]
	if !found {
		fmt.Println("Creando planeta...")
		// si no existe el planeta inicializamos todo como corresponde
		s.planets[informanteMessage.PlanetName] = make(map[string]string)
		s.planetVersions[informanteMessage.PlanetName] = &[3]int{0, 0, 0}
		f, e := os.Create(fmt.Sprintf("%v.txt", informanteMessage.PlanetName))
		if e != nil {
			panic(e)
		}
		f.Close()

		fLog, e := os.Create(fmt.Sprintf("%v_log.txt", informanteMessage.PlanetName))
		if e != nil {
			panic(e)
		}
		fLog.Close()

	}
	// guardamos los datos
	s.planetVersions[informanteMessage.PlanetName][s.updateOn]++
	s.planets[informanteMessage.PlanetName][informanteMessage.CityName] = informanteMessage.NewValue
	planetFile, err := os.OpenFile(fmt.Sprintf("%v.txt", informanteMessage.PlanetName), os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer planetFile.Close()
	if _, err := planetFile.WriteString(
		fmt.Sprintf(
			"%v %v %v",
			informanteMessage.PlanetName,
			informanteMessage.CityName,
			informanteMessage.NewValue,
		),
	); err != nil {
		log.Fatal(err)
	}
	UpdateLog(
		fmt.Sprintf("AddCity %v %v %v", informanteMessage.PlanetName, informanteMessage.CityName, informanteMessage.NewValue),
		informanteMessage.PlanetName,
	)
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateName(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {
	// updateamos archivo planeta
	fileName := fmt.Sprintf("%v.txt", informanteMessage.PlanetName)
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Aun no se ha creado una ciudad con este nombre.")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, informanteMessage.CityName) {
			lines[i] = fmt.Sprintf("%v %v %v", informanteMessage.PlanetName, informanteMessage.NewValue, s.planets[informanteMessage.PlanetName][informanteMessage.CityName])
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
	}
	// updateamos Log
	UpdateLog(
		fmt.Sprintf("UpdateName %v %v %v", informanteMessage.PlanetName, informanteMessage.CityName, informanteMessage.NewValue),
		informanteMessage.PlanetName,
	)
	s.planetVersions[informanteMessage.PlanetName][s.updateOn]++
	s.planets[informanteMessage.PlanetName][informanteMessage.NewValue] = s.planets[informanteMessage.PlanetName][informanteMessage.CityName]
	delete(s.planets[informanteMessage.PlanetName], informanteMessage.CityName)
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) UpdateNumber(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {
	// updateamos archivo planeta
	fileName := fmt.Sprintf("%v.txt", informanteMessage.PlanetName)
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Aun no se ha creado una ciudad con este nombre.")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, informanteMessage.CityName) {
			lines[i] = fmt.Sprintf("%v %v %v", informanteMessage.PlanetName, informanteMessage.CityName, informanteMessage.NewValue)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0777)
	if err != nil {
		log.Fatalln(err)
	}
	// updateamos Log
	UpdateLog(
		fmt.Sprintf("UpdateName %v %v %v", informanteMessage.PlanetName, informanteMessage.CityName, informanteMessage.NewValue),
		informanteMessage.PlanetName,
	)
	s.planetVersions[informanteMessage.PlanetName][s.updateOn]++
	s.planets[informanteMessage.PlanetName][informanteMessage.CityName] = informanteMessage.NewValue
	return &protos.FulcrumWriteMessage{}, nil
}

func (s *FulcrumServer) DeleteCity(
	ctx context.Context,
	informanteMessage *protos.InformantMessage,
) (*protos.FulcrumWriteMessage, error) {
	fpath := fmt.Sprintf("%v.txt", informanteMessage.PlanetName)
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), informanteMessage.CityName) {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fpath, buf.Bytes(), 0777)
	if err != nil {
		log.Fatal(err)
	}
	s.planetVersions[informanteMessage.PlanetName][s.updateOn]++
	UpdateLog(
		fmt.Sprintf("UpdateName %v %v", informanteMessage.PlanetName, informanteMessage.CityName),
		informanteMessage.PlanetName,
	)
	delete(s.planets[informanteMessage.PlanetName], informanteMessage.CityName)
	return &protos.FulcrumWriteMessage{}, nil
}

func UpdateLog(action string, planetName string) {
	logFile, err := os.OpenFile(fmt.Sprintf("%v_log.txt", planetName), os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer logFile.Close()
	if _, err := logFile.WriteString(action); err != nil {
		log.Fatal(err)
	}
}
