package atuadores

import (
	"context"
	"example/m/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Lampada struct {
	ClienteAtuador      pb.LampadaClient
	La AtuadorP
}
type Lamp struct {
	IntensidadeLuminosa float64 `json:"IntensidadeLuminosa"`
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Status              string  `json:"Status"`
}
type AtuadorG interface{
	//GetTipo() string
	//GetNome() string
	//Teste() string
}




func (l *Lampada) Conectar(ip string, port string) {
	fmt.Println("Cliente Atuador Lampada rodando...")

	l.La.Ip = ip
	l.La.Port = port
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewLampadaClient(conn)

	l.ClienteAtuador = client
}

func (l *Lampada) Ligar() *pb.Response {
	

	req := &pb.LampadaRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.LigarLampada(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.La.Status = res.Status
	l.La.Power = res.Power
	l.La.Nome = res.Nome
	l.La.Tipo = res.Tipo

	return res
}

func (l *Lampada) Desligar() *pb.Response {
	l.La.Power = false

	req := &pb.LampadaRequest{
		Nome: "lampada-1",
	}

	res, err := l.ClienteAtuador.DesligarLampada(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	l.La.Status = res.Status
	l.La.Power = res.Power
	l.La.Nome = res.Nome
	l.La.Tipo = res.Tipo

	fmt.Println(res)

	return res
}



func (l *Lampada) GetAtuador() AtuadorP {
	return l.La
}