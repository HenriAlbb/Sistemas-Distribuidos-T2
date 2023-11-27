package atuadores

import (
	"context"
	"example/m/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ArCondicionado struct {
	ClienteAtuador      pb.ArCondicionadoClient
	AC AtuadorP
}

type ArCondicionado2 struct {
	Nome string
	Tipo string
	Ip string
	Port string
	Power string
	Status string
	TemperaturaSetada float64
}




func (l *ArCondicionado) Conectar(ip string, port string) {
	fmt.Println("Cliente Atuador arCondicionado rodando...")

	l.AC.Ip = ip
	l.AC.Port = port
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewArCondicionadoClient(conn)

	l.ClienteAtuador = client
}

func (l *ArCondicionado) Ligar() *pb.Response {
	

	req := &pb.ArCondicionadoRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.LigarArCondicionado(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.AC.Status = res.Status
	l.AC.Power = res.Power
	l.AC.Nome = res.Nome
	l.AC.Tipo = res.Tipo

	return res
}

func (l *ArCondicionado) Desligar() *pb.Response {
	l.AC.Power = false

	req := &pb.ArCondicionadoRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.DesligarArCondicionado(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.AC.Status = res.Status
	l.AC.Power = res.Power
	l.AC.Nome = res.Nome
	l.AC.Tipo = res.Tipo

	fmt.Println(res)

	return res
}



func (l *ArCondicionado) GetAtuador() AtuadorP {
	return l.AC
}