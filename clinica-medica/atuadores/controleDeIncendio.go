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

type ControleDeIncendio struct {
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Status              string  `json:"Status"`
	IntensidadeLuminosa float32 `json:"IntensidadeLuminosa"`
}

type Server struct {
	pb.UnimplementedControleDeIncendioServer
	CDI ControleDeIncendio
}

func (s *Server) LigarControleDeIncendio(ctc context.Context, in *pb.ControleDeIncendioRequest) (*pb.Response, error) {
	s.CDI.Power = true
	s.CDI.Status = "Lampada Ligada"
	//fmt.Println(s.CDI.Power)
	return &pb.Response{Nome: s.CDI.Nome, Tipo: s.CDI.Tipo, Power: s.CDI.Power, Status: s.CDI.Status, IntensidadeLuminosa: s.CDI.IntensidadeLuminosa}, nil
}

func (s *Server) DesligarControleDeIncendio(ctc context.Context, in *pb.ControleDeIncendioRequest) (*pb.Response, error) {
	s.CDI.Power = false
	s.CDI.Status = "Lampada Desligada"
	//fmt.Println(s.CDI.Power)
	return &pb.Response{Nome: s.CDI.Nome, Tipo: s.CDI.Tipo, Power: s.CDI.Power, Status: s.CDI.Status, IntensidadeLuminosa: s.CDI.IntensidadeLuminosa}, nil
}

func main() {
	fmt.Println("Running Atuador Controle De Incendio")

	controleDeIncendio := &ControleDeIncendio{"Controle-De-Incendio-1", "Controle-De-Incendio", "localhost", "6000", false, "", 0}
	server := &Server{}
	server.CDI = *controleDeIncendio

	listener, err := net.Listen("tcp", "localhost:6000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterControleDeIncendioServer(s, server)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
