package atuadores

import (
	"context"
	"example/m/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ControleDeIncendio struct{
	ClienteAtuador      pb.ControleDeIncendioClient
	CDI AtuadorP
}


type ControleDeIncendio2 struct {
	Nome string
	Tipo string
	Ip string
	Port string
	Power string
	Status string
	TemperaturaSetada float64
}




func (l *ControleDeIncendio) Conectar(ip string, port string) {
	fmt.Println("Cliente Atuador controleDeIncendio rodando...")

	l.CDI.Ip = ip
	l.CDI.Port = port
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewControleDeIncendioClient(conn)

	l.ClienteAtuador = client
}

func (l *ControleDeIncendio) Ligar() *pb.Response {
	

	req := &pb.ControleDeIncendioRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.LigarControleDeIncendio(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.CDI.Status = res.Status
	l.CDI.Power = res.Power
	l.CDI.Nome = res.Nome
	l.CDI.Tipo = res.Tipo

	return res
}

func (l *ControleDeIncendio) Desligar() *pb.Response {
	l.CDI.Power = false

	req := &pb.ControleDeIncendioRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.DesligarControleDeIncendio(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.CDI.Status = res.Status
	l.CDI.Power = res.Power
	l.CDI.Nome = res.Nome
	l.CDI.Tipo = res.Tipo

	fmt.Println(res)

	return res
}



func (l *ControleDeIncendio) GetAtuador() AtuadorP {
	return l.CDI
}