package atuadores

import (
	"example/m/pb"
)

type Atuador interface {
	Ligar() *pb.Response
	Desligar() *pb.Response
	Conectar(ip string, port string)
	GetAtuador() AtuadorP 
}

type AtuadorP struct{
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Status              string  `json:"Status"`
}

func (l *AtuadorP) GetTipo() string {
	return l.Tipo
}

func (l *AtuadorP) GetNome() string {
	return l.Nome
}