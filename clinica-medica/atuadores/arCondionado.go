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

type ArCondicionado struct {
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Status              string  `json:"Status"`
	IntensidadeLuminosa float32 `json:"IntensidadeLuminosa"`
}

type Server struct {
	pb.UnimplementedArCondicionadoServer
	AC ArCondicionado
}

func (s *Server) LigarArCondicionado(ctc context.Context, in *pb.ArCondicionadoRequest) (*pb.Response, error) {
	s.AC.Power = true
	s.AC.Status = "Lampada Ligada"
	//fmt.Println(s.AC.Power)
	return &pb.Response{Nome: s.AC.Nome, Tipo: s.AC.Tipo, Power: s.AC.Power, Status: s.AC.Status, IntensidadeLuminosa: s.AC.IntensidadeLuminosa}, nil
}

func (s *Server) DesligarArCondicionado(ctc context.Context, in *pb.ArCondicionadoRequest) (*pb.Response, error) {
	s.AC.Power = false
	s.AC.Status = "Lampada Desligada"
	//fmt.Println(s.AC.Power)
	return &pb.Response{Nome: s.AC.Nome, Tipo: s.AC.Tipo, Power: s.AC.Power, Status: s.AC.Status, IntensidadeLuminosa: s.AC.IntensidadeLuminosa}, nil
}

func main() {
	fmt.Println("Running Atuador Ar Condicionado")

	arCondicionado := &ArCondicionado{"Ar-Condicionado-1", "Ar-Condicionado", "localhost", "8000", false, "", 0}
	server := &Server{}
	server.AC = *arCondicionado

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterArCondicionadoServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
