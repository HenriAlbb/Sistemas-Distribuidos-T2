package main

import (
	"context"
	"example/m/pb"
	"fmt"
	"log"
	"net"

	//pb "command-line-argumentsC:\\Users\\henri.DESKTOP-5CQDVBM\\Engenharia-Diciplinas\\SISTEMAS DISTRIBUÍDOS\\T2\\clinica-medica\\pb\\lampada.pb.go"
	//pb "command-line-argumentsC:\\Users\\henri.DESKTOP-5CQDVBM\\Engenharia-Diciplinas\\SISTEMAS DISTRIBUÍDOS\\T2\\clinica-medica\\pb\\lampada_grpc.pb.go"
	"google.golang.org/grpc"
)

type Lampada struct {
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Status              string  `json:"Status"`
	IntensidadeLuminosa float32 `json:"IntensidadeLuminosa"`
}

type Server struct {
	pb.UnimplementedLampadaServer
	Lamp Lampada
}

func (s *Server) LigarLampada(ctc context.Context, in *pb.LampadaRequest) (*pb.Response, error) {
	s.Lamp.Power = true
	s.Lamp.Status = "Lampada Ligada"
	//fmt.Println(s.Lamp.Power)
	return &pb.Response{Nome: s.Lamp.Nome, Tipo: s.Lamp.Tipo, Power: s.Lamp.Power, Status: s.Lamp.Status, IntensidadeLuminosa: s.Lamp.IntensidadeLuminosa}, nil
}

func (s *Server) DesligarLampada(ctc context.Context, in *pb.LampadaRequest) (*pb.Response, error) {
	s.Lamp.Power = false
	s.Lamp.Status = "Lampada Desligada"
	//fmt.Println(s.Lamp.Power)
	return &pb.Response{Nome: s.Lamp.Nome, Tipo: s.Lamp.Tipo, Power: s.Lamp.Power, Status: s.Lamp.Status, IntensidadeLuminosa: s.Lamp.IntensidadeLuminosa}, nil
}

func main() {
	fmt.Println("Running Atuador Lampada")

	lampada := &Lampada{"Lampada-1", "Lampada", "localhost", "9000", false, "", 0}
	server := &Server{}
	server.Lamp = *lampada

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterLampadaServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
